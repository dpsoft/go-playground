package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/router"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument/syncfloat64"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/histogram"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	"go.opentelemetry.io/otel/sdk/metric/export/aggregation"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	selector "go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"log"
	"net/http"
)

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Id    int     `json:"id"`
}

type Inventory struct {
	Products []Product `json:"products"`
}

var inventory = Inventory{
	Products: []Product{
		{Name: "Laptop", Price: 1000, Id: 1},
		{Name: "Mobile", Price: 500, Id: 2},
		{Name: "Tablet", Price: 700, Id: 3},
	},
}

func initMeter() {
	config := prometheus.Config{DefaultHistogramBoundaries: []float64{1, 2, 5, 10, 20, 50}}
	c := controller.New(
		processor.NewFactory(
			selector.NewWithHistogramDistribution(histogram.WithExplicitBoundaries(config.DefaultHistogramBoundaries)),
			//aggregation.StatelessTemporalitySelector(),
			aggregation.CumulativeTemporalitySelector(),
			//aggregation.DeltaTemporalitySelector(),
			processor.WithMemory(true),
		),
	)

	exporter, err := prometheus.New(config, c)
	if err != nil {
		log.Panicf("failed to initialize prometheus exporter %v", err)
	}

	global.SetMeterProvider(exporter.MeterProvider())

	http.HandleFunc("/", exporter.ServeHTTP)

	go func() {
		_ = http.ListenAndServe(":2222", nil)
	}()

	fmt.Println("Prometheus server running on :2222")
}

func getInventory(ctx *gin.Context, h syncfloat64.Counter) {
	h.Add(ctx, 1)
	ctx.JSON(200, gin.H{"inventory": inventory})
}

var inventoryKey = attribute.Key("inventory-service-attribute")

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denomiator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	//type T struct{ a, b, c int }
	//var x = T{1, 2, 3}
	//var y = T{4, 5, 6}
	//var z = T{7, 8, 9}
	//
	//var v atomic.Value
	//v.Store(x)
	//fmt.Println(v)
	//
	//old := v.Swap(y)
	//fmt.Println(v)
	//fmt.Println(old.(T))
	//
	//swapped := v.CompareAndSwap(x, z)
	//fmt.Println(swapped, v) // false {{4 5 6}}
	//swapped = v.CompareAndSwap(y, z)
	//fmt.Println(swapped, v) // true {{7 8 9}}
	//
	//filter := bloom.NewWithEstimates(1000000, 0.001)
	//
	//for i := 0; i < 1000000; i++ {
	//	filter.AddString(fmt.Sprint("Love-", i))
	//}
	//
	//fmt.Print(filter.K())
	//if filter.TestString("Love-99") {
	//	fmt.Println("Diego is in the filter")
	//}
	//remainder, mod, errr := calcRemainderAndMod(1, 0)
	//
	//if errr != nil {
	//	fmt.Println(errr)
	//	os.Exit(1)
	//}
	//fmt.Println(remainder, mod)
	//
	//initMeter()
	//meter := global.MeterProvider().Meter("inventory-service")
	//
	//histogram, err := meter.SyncFloat64().Histogram("histogram")
	//if err != nil {
	//	log.Panicf("failed to initialize instrument: %v", err)
	//}
	//
	//counter, err := meter.SyncFloat64().Counter("counter")
	//if err != nil {
	//	log.Panicf("failed to initialize instrument: %v", err)
	//}
	//
	//ctx := context.Background()
	//
	//commonLabels := []attribute.KeyValue{inventoryKey.Int(10), attribute.String("A", "1"), attribute.String("B", "2"), attribute.String("C", "3")}
	//notSoCommonLabels := []attribute.KeyValue{inventoryKey.Int(13)}
	//
	//histogram.Record(ctx, 12.0, commonLabels...)
	//
	//counter.Add(ctx, 13.0, commonLabels...)
	//counter.Add(ctx, 1.0, notSoCommonLabels...)

	r := gin.Default()
	route := router.NewAdminRoute()
	route.Register(r.Group("/admin"))
	_ = r.Run("localhost:8080")
}
