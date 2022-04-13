package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/router"
	"github.com/go-playground/stdout"
)

func main() {

	stdout.Example()

	r := gin.Default()
	route := router.NewAdminRoute()
	route.Register(r.Group("/admin"))
	_ = r.Run("localhost:8080")
}
