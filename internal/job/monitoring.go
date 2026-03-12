package job

import (
	"log/slog"
	"time"

	"github.com/spf13/cast"
	"gorm.io/gorm"

	"github.com/acepanel/panel/v3/internal/app"
	"github.com/acepanel/panel/v3/internal/biz"
	"github.com/acepanel/panel/v3/pkg/tools"
)

// Monitoring 系统监控
type Monitoring struct {
	db          *gorm.DB
	log         *slog.Logger
	settingRepo biz.SettingRepo
	lastRun     time.Time
}

func NewMonitoring(db *gorm.DB, log *slog.Logger, setting biz.SettingRepo) *Monitoring {
	return &Monitoring{
		db:          db,
		log:         log,
		settingRepo: setting,
	}
}

func (r *Monitoring) Run() {
	if app.Status != app.StatusNormal {
		return
	}

	monitor, err := r.settingRepo.Get(biz.SettingKeyMonitor)
	if err != nil || !cast.ToBool(monitor) {
		return
	}

	// 根据采集间隔判断是否该采集
	interval, _ := r.settingRepo.GetInt(biz.SettingKeyMonitorInterval, 1)
	if interval < 1 {
		interval = 1
	}
	if !r.lastRun.IsZero() && time.Since(r.lastRun) < time.Duration(interval)*time.Minute-30*time.Second {
		return
	}
	r.lastRun = time.Now()

	info := tools.CurrentInfo(nil, nil)
	info.TopProcesses = tools.CollectTopProcesses()

	// 去除部分数据以减少数据库存储
	info.Disk = nil
	info.Cpus = nil

	if app.Status != app.StatusNormal {
		return
	}

	if err = r.db.Create(&biz.Monitor{Info: info}).Error; err != nil {
		r.log.Warn("failed to create monitor record", slog.String("type", biz.OperationTypeMonitor), slog.Uint64("operator_id", 0), slog.Any("err", err))
		return
	}

	// 删除过期数据
	dayStr, err := r.settingRepo.Get(biz.SettingKeyMonitorDays)
	if err != nil {
		return
	}
	day := cast.ToInt(dayStr)
	if day <= 0 || app.Status != app.StatusNormal {
		return
	}
	if err = r.db.Where("created_at < ?", time.Now().AddDate(0, 0, -day).Format(time.DateTime)).Delete(&biz.Monitor{}).Error; err != nil {
		r.log.Warn("failed to delete monitor record", slog.String("type", biz.OperationTypeMonitor), slog.Uint64("operator_id", 0), slog.Any("err", err))
		return
	}
}
