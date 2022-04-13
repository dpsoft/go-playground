package router

import (
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/gin-gonic/gin"
	"log"
	"regexp"
)

type Metric struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

type MetricRoute interface {
	Register(router *gin.RouterGroup)
}

type metricRoute struct {
	regex *[]regexp.Regexp
	bloom *bloom.BloomFilter
}

func (m metricRoute) Register(router *gin.RouterGroup) {
	router.POST("/ingest", m.ingest)
}

func (m metricRoute) ingest(ctx *gin.Context) {
	var metric Metric
	if err := ctx.ShouldBindJSON(&metric); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input!"})
		return
	}

	log.Println("creating new metric...")

	ctx.JSON(200, metric)
}

func NewMetricRoute() MetricRoute {
	var x []regexp.Regexp

	r, _ := regexp.Compile("p([a-z]+)ch")

	x = append(x, *r)

	var b = bloom.NewWithEstimates(1000000, 0.001)
	b.AddString("peach")

	return metricRoute{
		regex: &x,
		bloom: b,
	}
}
