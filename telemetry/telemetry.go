package telemetry

import (
	"context"
	"log"
	"sync"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
)

type telemetry struct {
	counter syncint64.Counter
}

func (t *telemetry) IncrementCounter(attrs ...attribute.KeyValue) {
	t.counter.Add(context.Background(), 1, attrs...)
}

var instance *telemetry
var once sync.Once

func Metrics() *telemetry {
	once.Do(func() {
		meter := global.MeterProvider().Meter("awesome-instrumentation")

		counter, err := meter.SyncInt64().Counter("awesome-counter")
		if err != nil {
			log.Fatalf("failed to create counter: %v", err)
			return
		}

		instance = &telemetry{
			counter: counter,
		}
	})
	return instance
}
