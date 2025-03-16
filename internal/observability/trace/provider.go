package trace

import (
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"

	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func NewProvider(
	cfg *entity.AppConfig,
	exporter tracesdk.SpanExporter,
) (*tracesdk.TracerProvider, error) {
	if !cfg.Trace.Enabled {
		return nil, nil
	}

	tracerProvider := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exporter),
		tracesdk.WithSampler(tracesdk.TraceIDRatioBased(cfg.Trace.SampleRatio)),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.Name),
		)),
	)

	return tracerProvider, nil
}

func RunProvider(tracerProvider *tracesdk.TracerProvider) {
	if tracerProvider == nil {
		return
	}

	otel.SetTracerProvider(tracerProvider)
}
