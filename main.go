package main

import (
	log "github.com/go-playground/logger"
	"github.com/go-playground/server"
	"github.com/go-playground/settings"
)

func main() {
	cfg := settings.NewSettings()
	logger := log.NewBuiltinLogger()
	s := server.New(cfg, logger)

	s.Run()
}
