package handler

import (
	"github.com/google/wire"
	"go-layout/internal/service"
)

var Set = wire.Struct(new(Handler), "*")

type Handler struct {
	Srv *service.Service
}
