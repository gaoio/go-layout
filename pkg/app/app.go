package app

import (
	"github.com/google/wire"
	"go-layout/internal/service"
	"log"
	"net/http"
)

var ProviderSet = wire.Struct(new(App), "*")

type App struct {
	HTTPServer *http.Server
	Svc        *service.Service
}

func (a *App) Run() error {
	log.Printf("http server: %s", a.HTTPServer.Addr)
	return a.HTTPServer.ListenAndServe()
}

func (a *App) Close() {
	log.Println("close...")
	a.Svc.Close()
}
