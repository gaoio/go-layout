//go:build wireinject

package main

import (
	"github.com/google/wire"
	"go-layout/internal/conf"
	"go-layout/internal/dao"
	"go-layout/internal/server"
	"go-layout/internal/service"
	"go-layout/pkg/app"
)

func wireApp(c *conf.Conf) *app.App {
	wire.Build(
		app.ProviderSet,
		server.ProviderSet,
		service.ProviderSet,
		dao.ProviderSet,
	)
	return new(app.App)
}
