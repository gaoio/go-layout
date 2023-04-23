package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go-layout/internal/conf"
	"go-layout/internal/server/http/handler"
	"go-layout/internal/server/http/middleware"
	"go-layout/internal/server/http/router"
	"net/http"
)

var ProviderSet = wire.NewSet(router.Set, handler.Set, middleware.Set, NewServer)

func NewServer(router *router.Router, c *conf.Conf) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	router.RegisterAPI(engine)

	return &http.Server{
		Addr:    c.Server.HttpAddr,
		Handler: engine,
	}
}
