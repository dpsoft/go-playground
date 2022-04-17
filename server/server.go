package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/logger"
	"github.com/go-playground/router"
	"github.com/go-playground/settings"
)

type Server struct {
	settings settings.Settings
	logger   logger.Logger
}

func (s *Server) Run() {
	var host = s.settings.Server.Host
	var port = s.settings.Server.Port

	r := gin.Default()

	adminRoute := router.NewAdminRoute()
	metricRoute := router.NewMetricRoute()

	adminRoute.Register(r.Group("/admin"))
	metricRoute.Register(r.Group("/metrics"))

	_ = r.Run(host + ":" + port)
}

func New(settings settings.Settings, logger logger.Logger) *Server {
	return &Server{
		settings: settings,
		logger:   logger,
	}
}
