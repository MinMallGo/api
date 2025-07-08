package utils

import (
	"api/order_api/global"
	"context"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

const (
	serviceName    = "Go-Jaeger-Demo"
	jaegerEndpoint = "127.0.0.1:4318"
)

// SetupTracer 设置Tracer
func SetupTracer(ctx context.Context) (func(context.Context) error, error) {
	tracerProvider, err := newJaegerTraceProvider(ctx)
	if err != nil {
		return nil, err
	}
	otel.SetTracerProvider(tracerProvider)
	return tracerProvider.Shutdown, nil
}

// NewJaegerTraceProvider 创建一个 Jaeger Trace Provider
func newJaegerTraceProvider(ctx context.Context) (*traceSDK.TracerProvider, error) {
	// 创建一个使用 HTTP 协议连接本机Jaeger的 Exporter
	exp, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(jaegerEndpoint),
		otlptracehttp.WithInsecure())
	if err != nil {
		return nil, err
	}
	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceName(serviceName)))
	if err != nil {
		return nil, err
	}
	traceProvider := traceSDK.NewTracerProvider(
		traceSDK.WithResource(res),
		traceSDK.WithSampler(traceSDK.AlwaysSample()), // 采样
		traceSDK.WithBatcher(exp, traceSDK.WithBatchTimeout(time.Second)),
	)
	return traceProvider, nil
}

func InjectTracer(ctx context.Context, msg *rmq_client.Message) *rmq_client.Message {
	props := map[string]string{}
	otel.GetTextMapPropagator().Inject(ctx, propagation.MapCarrier(props))
	for k, v := range props {
		msg.AddProperty(k, v)
	}
	log.Println(msg)
	return msg
}

func ExtractTracer(ctx context.Context, msg map[string]string) context.Context {
	carrier := propagation.MapCarrier{}
	for k, v := range msg {
		carrier[k] = v
	}
	return otel.GetTextMapPropagator().Extract(ctx, carrier)
}

// Inject2Gin 把trace的上下级关系注入到ctx里面，以便后续使用
func Inject2Gin(ctx context.Context, c *gin.Context) {
	props := map[string]string{}
	otel.GetTextMapPropagator().Inject(ctx, propagation.MapCarrier(props))
	c.Set(global.TracerCtxName, props)
}

// ExtFromGin  从gin里面解析出数据来，然后转成map[string]string的格式,并解析数据到ctx里面并返回
func ExtFromGin(ctx context.Context, c *gin.Context) context.Context {
	// 解析出来你是个map[string]string，不对直接返回context.Background()
	traceParent, exists := c.Get(global.TracerCtxName)
	if !exists {
		return ctx
	}

	m, ok := traceParent.(map[string]string)
	if !ok {
		return ctx
	}

	// 解析trace到ctx里面
	carrier := propagation.MapCarrier{}
	for k, v := range m {
		carrier[k] = v
	}
	return otel.GetTextMapPropagator().Extract(ctx, carrier)
}

// InjectOTEL 需要手动注入metadata
func InjectOTEL(ctx context.Context) context.Context {
	md := metadata.New(nil)
	propagator := otel.GetTextMapPropagator()
	propagator.Inject(ctx, propagation.HeaderCarrier(md))
	return metadata.NewOutgoingContext(ctx, md)
}
