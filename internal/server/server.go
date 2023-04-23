package server

import (
	"github.com/google/wire"
	"go-layout/internal/server/http"
)

var ProviderSet = wire.NewSet(http.ProviderSet)
