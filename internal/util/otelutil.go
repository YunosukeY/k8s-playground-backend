package util

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

func NewTracer(service string) (trace.Tracer, func()) {
	url := getJaegerURL()

	ctx := context.Background()
	opt := otlptracehttp.WithEndpoint(url)
	exporter, err := otlptracehttp.New(ctx, opt)
	if err != nil {
		log.Panic().Err(err).Msg("")
		panic(err)
	}

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
		)),
	)
	otel.SetTracerProvider(provider)

	tracer := provider.Tracer("")
	shutdownProvider := func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err := provider.Shutdown(ctx); err != nil {
			log.Panic().Err(err).Msg("")
			panic(err)
		}
	}
	return tracer, shutdownProvider
}
