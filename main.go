package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/functions"
	"github.com/go-playground/router"
	"math"
)

func main() {

	numbers := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice := functions.MapSlice(numbers, func(n float64) float64 {
		return math.Pow(n, 2)
	})

	println(slice)

	r := gin.Default()
	route := router.NewAdminRoute()
	route.Register(r.Group("/admin"))
	_ = r.Run("localhost:8080")
}
