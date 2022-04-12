package router

import (
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"regexp"
)

type AdminRoute interface {
	Register(router *gin.RouterGroup)
}
type adminRoute struct {
	regex *regexp.Regexp
	bloom *bloom.BloomFilter
}

func (a adminRoute) Register(router *gin.RouterGroup) {
	router.POST("/override-metrics-tags", a.overrideMetricsTags)
}

func (a adminRoute) overrideMetricsTags(c *gin.Context) {
	fmt.Println(a.bloom.TestString("peach"))
	log.Print(a.bloom.TestString("peach"))
	matchString := a.regex.MatchString("peach")

	fmt.Println(matchString)
	c.JSON(200, gin.H{"message": "override-metrics-tags"})
}

func NewAdminRoute() AdminRoute {
	r, _ := regexp.Compile("p([a-z]+)ch")
	var b = bloom.NewWithEstimates(1000000, 0.001)
	b.AddString("peach")

	return adminRoute{
		regex: r,
		bloom: b,
	}
}
