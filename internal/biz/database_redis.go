package biz

import (
	"github.com/acepanel/panel/v3/internal/http/request"
	"github.com/acepanel/panel/v3/pkg/db"
)

type DatabaseRedisRepo interface {
	Databases(req *request.DatabaseRedisDatabases) (int, error)
	Data(req *request.DatabaseRedisData) ([]db.RedisKV, int, error)
	KeyGet(req *request.DatabaseRedisKeyGet) (*db.RedisKV, error)
	KeySet(req *request.DatabaseRedisKeySet) error
	KeyDelete(req *request.DatabaseRedisKeyDelete) error
	KeyTTL(req *request.DatabaseRedisKeyTTL) error
	KeyRename(req *request.DatabaseRedisKeyRename) error
	Clear(req *request.DatabaseRedisClear) error
}
