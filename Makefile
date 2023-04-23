.PHONY: build

build:
	go build -ldflags "-X main.buildTime=$(shell date '+%Y-%m-%dT%H:%M:%S')" -gcflags "-B" -o main cmd/main.go cmd/wire_gen.go

wire:
	wire gen cmd/wire.go

run:
	go run cmd/main.go cmd/wire_gen.go