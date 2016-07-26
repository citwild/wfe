package middleware

import (
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/service"
	"github.com/citwild/wfe/store"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Injector struct {
	servers api.Servers
	stores  store.Stores
}

func NewInjector(srvs api.Servers, strs store.Stores) *Injector {
	return &Injector{servers: srvs, stores: strs}
}

func (i *Injector) Inject(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, _ error) {
	ctx = service.WithServers(ctx, i.servers)
	ctx = store.WithStores(ctx, i.stores)
	return handler(ctx, req)
}
