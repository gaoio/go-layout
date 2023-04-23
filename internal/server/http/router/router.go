package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go-layout/internal/server/http/handler"
	"go-layout/internal/server/http/middleware"
)

var Set = wire.Struct(new(Router), "*")

type Router struct {
	Handler    *handler.Handler
	Middleware *middleware.Middleware
}

func (r *Router) RegisterAPI(engine *gin.Engine) {
	engine.Use(gin.Recovery(), r.Middleware.Trace(), r.Middleware.CORS())
	engine.GET("/ping", r.Handler.Ping)
}
