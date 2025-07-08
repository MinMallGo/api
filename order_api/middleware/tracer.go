package middleware

import (
	"api/order_api/utils"
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// Trace 先测试一下是否可行
func Trace() gin.HandlerFunc {
	// 目的是为了每一次请求都生成tracer
	return func(c *gin.Context) {
		// 初始化和连接otel以及配置jaeger的tracer

		// 还有这里不能全局设置trace？
		spanName := c.Request.URL.String()
		ctx := context.Background()
		shutdown, err := utils.SetupTracer(ctx)
		if err != nil {
			panic(err)
		}
		defer func() {
			_ = shutdown(ctx)
		}()

		tracer := otel.Tracer("test-tracer") // todo 这里的server-name还有ip应该是nacos里面的配置来的
		baseAttrs := []attribute.KeyValue{
			attribute.String("method", c.Request.Method),
			attribute.String("ip", c.ClientIP()),
			attribute.String("agent", c.Request.UserAgent()),
		}

		// 开启span
		ctx, span := tracer.Start(ctx, spanName, trace.WithAttributes(baseAttrs...))
		otel.SetTextMapPropagator(propagation.TraceContext{})
		// 结束span
		span.End()
		// 将链路的上下文注入到gin里面
		utils.Inject2Gin(ctx, c)
		// 在 ctx 里面设置props
		c.Next()
	}
}
