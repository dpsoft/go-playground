package main

import (
	"github.com/go-playground/server"
	"github.com/go-playground/settings"
)

func main() {
	cfg := settings.NewSettings()
	s := server.New(cfg)

	s.Run()
}
