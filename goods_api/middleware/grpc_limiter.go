package middleware

import (
	"context"
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryServerInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		callOpts ...grpc.CallOption,
	) error {
		resource := fmt.Sprintf("grpc:%s", method)
		entry, blockErr := sentinel.Entry(
			resource,
			sentinel.WithTrafficType(base.Inbound),
		)
		if blockErr != nil {
			return status.Errorf(codes.ResourceExhausted, "请求过于频繁，请稍后再试")
		}
		defer entry.Exit()
		return invoker(ctx, method, req, reply, cc, callOpts...)
	}
}
