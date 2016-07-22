package middleware

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func InitContext(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, _ error) {
	return handler(ctx, req)
}
