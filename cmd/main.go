package main

import (
	"go-layout/internal/conf"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c, err := conf.Load()
	if err != nil {
		panic(err)
	}

	app := wireApp(c)

	go func() {
		if err = app.Run(); err != nil {
			panic(err)
		}
	}()

	cc := make(chan os.Signal, 1)
	signal.Notify(cc, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-cc
		log.Printf("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			app.Close()
			log.Println("exit!!!")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
