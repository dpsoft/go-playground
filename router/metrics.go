package router

import (
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/gin-gonic/gin"
	"log"
	"regexp"
	"sync/atomic"
	"time"
)

type Metric struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

type MetricRoute interface {
	Register(router *gin.RouterGroup)
}

type metricRoute struct {
	regex  *[]regexp.Regexp
	bloom  *bloom.BloomFilter
	bloom2 atomic.Value
}

func (m metricRoute) Register(router *gin.RouterGroup) {
	router.POST("/ingest", m.ingest)

	go m.run()
}

func (m metricRoute) ingest(ctx *gin.Context) {
	var metric Metric
	if err := ctx.ShouldBindJSON(&metric); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid input!"})
		return
	}

	//telemetry.Instance().IncrementCounter()
	log.Println("creating new metric...")

	ctx.JSON(200, metric)
}

func (m metricRoute) run() {
	ticker := time.NewTicker(5 * time.Second)

	for _ = range ticker.C {
		_, _ = m.reload()
	}
}

func (m metricRoute) reload() (int, error) {
	var bloooom = bloom.NewWithEstimates(100, 0.001)
	bloooom.AddString("peacha")

	m.bloom2.Swap(bloooom)

	return fmt.Println("Tock")
}

func NewMetricRoute() MetricRoute {
	var r = []string{"p([a-z]+)ch", "p([a-z]+)ch"}
	var regexps = make([]regexp.Regexp, len(r))

	for i, v := range r {
		regex, _ := regexp.Compile(v)
		regexps[i] = *regex
	}

	var b = bloom.NewWithEstimates(1000000, 0.001)
	b.AddString("peach")

	return metricRoute{
		regex:  &regexps,
		bloom:  b,
		bloom2: atomic.Value{},
	}
}
