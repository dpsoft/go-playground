package dispatcher

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

var inventoryKey = attribute.Key("inventory-service-attribute")

func getInventory(ctx *gin.Context, h syncfloat64.Counter) {
	h.Add(ctx, 1)
	ctx.JSON(200, gin.H{"inventory": inventory})
}

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
