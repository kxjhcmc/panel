package bootstrap

import (
	"log/slog"

	"gorm.io/gorm"

	"github.com/acepanel/panel/v3/internal/taskqueue"
	"github.com/acepanel/panel/v3/pkg/types"
)

// NewRunner 创建任务运行器
func NewRunner(db *gorm.DB, log *slog.Logger) types.TaskRunner {
	return taskqueue.NewRunner(db, log)
}
