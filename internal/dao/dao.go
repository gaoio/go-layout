package dao

import (
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"go-layout/internal/conf"
	"go-layout/internal/pkg/gormx"
	"go-layout/internal/pkg/redisx"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(New)

type Dao struct {
	DB  *gorm.DB
	RDS *redis.Client
}

func New(c *conf.Conf) *Dao {
	return &Dao{
		DB:  gormx.NewDB(c.DB.StayDSN),
		RDS: redisx.New(c.Cache.MainRedisAddr),
	}
}

func (d *Dao) Close() {
	d.RDS.Close()
}
