package middleware

import (
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/service"
	"github.com/citwild/wfe/store"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func NewUnaryServiceInjector(s api.Servers) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, next grpc.UnaryHandler) (resp interface{}, err error) {
		ctx = service.WithServers(ctx, s)
		return next(ctx, req)
	}
}

func NewUnaryStoreInjector(s store.Stores) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, next grpc.UnaryHandler) (resp interface{}, err error) {
		ctx = store.WithStores(ctx, s)
		return next(ctx, req)
	}
}
