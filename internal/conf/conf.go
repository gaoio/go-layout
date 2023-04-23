package conf

import (
	"github.com/BurntSushi/toml"
)

type Conf struct {
	Server         Server         `toml:"server"`
	DB             DB             `toml:"db"`
	Cache          Cache          `toml:"cache"`
	InternalServer InternalServer `toml:"internal_server"`
}

type Server struct {
	HttpAddr string `toml:"http_addr"`
}

type DB struct {
	StayDSN string `toml:"stay_dsn"`
	LiveDSN string `toml:"live_dsn"`
}

type Cache struct {
	MainRedisAddr string `toml:"main_redis_addr"`
}

type InternalServer struct {
	Msg string `toml:"msg"`
}

func Load() (conf *Conf, err error) {
	_, err = toml.DecodeFile("./conf.toml", &conf)
	return
}
