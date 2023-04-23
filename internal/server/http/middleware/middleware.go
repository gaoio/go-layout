package middleware

import (
	"github.com/google/wire"
	"go-layout/internal/service"
)

var Set = wire.Struct(new(Middleware), "*")

type Middleware struct {
	Srv service.Service
}
