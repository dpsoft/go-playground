package main

import (
	"github.com/go-playground/stdout"
	"github.com/go-playground/telemetry"
)

func main() {
	//cfg := settings.NewSettings()
	//logger := log.NewBuiltinLogger()
	//s := server.New(cfg, logger)

	telemetry.Metrics().IncrementCounter()
	stdout.Example()

	//s.Run()
}
