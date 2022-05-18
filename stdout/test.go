package stdout

import (
	"context"
	"log"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
)

const (
	instrumentationName    = "github.com/instrumentron"
	instrumentationVersion = "v0.1.0"
)

var (
	loopCounter syncint64.Counter
	paramValue  syncint64.Histogram

	nameKey = attribute.Key("function.name")
)

func add(ctx context.Context, x, y int64) int64 {
	nameKV := nameKey.String("add")

	loopCounter.Add(ctx, 1, nameKV)
	paramValue.Record(ctx, x, nameKV)
	paramValue.Record(ctx, y, nameKV)

	return x + y
}

func multiply(ctx context.Context, x, y int64) int64 {
	nameKV := nameKey.String("multiply")

	loopCounter.Add(ctx, 1, nameKV)
	paramValue.Record(ctx, x, nameKV)
	paramValue.Record(ctx, y, nameKV)

	return x * y
}

func InstallExportPipeline(ctx context.Context) func() {
	exporter, err := stdoutmetric.New(stdoutmetric.WithPrettyPrint())
	if err != nil {
		log.Fatalf("creating stdoutmetric exporter: %v", err)
	}

	pusher := controller.New(
		processor.NewFactory(
			simple.NewWithInexpensiveDistribution(),
			exporter,
		),
		controller.WithExporter(exporter),
	)
	if err = pusher.Start(ctx); err != nil {
		log.Fatalf("starting push controller: %v", err)
	}

	global.SetMeterProvider(pusher)

	meter := global.MeterProvider().Meter(instrumentationName, metric.WithInstrumentationVersion(instrumentationVersion))

	loopCounter, err = meter.SyncInt64().Counter("function.loops")
	if err != nil {
		log.Fatalf("creating instrument: %v", err)
	}
	paramValue, err = meter.SyncInt64().Histogram("function.param")
	if err != nil {
		log.Fatalf("creating instrument: %v", err)
	}

	return func() {
		if err := pusher.Stop(ctx); err != nil {
			log.Fatalf("stopping push controller: %v", err)
		}
	}
}

func Example() {
	ctx := context.Background()

	// TODO: Registers a meter Provider globally.
	cleanup := InstallExportPipeline(ctx)
	defer cleanup()

	log.Println("the answer is", add(ctx, multiply(ctx, multiply(ctx, 2, 2), 10), 2))
}
