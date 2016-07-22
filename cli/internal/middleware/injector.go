package middleware

import (
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/servers"
	"github.com/citwild/wfe/stores"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Injector struct {
	servers api.Servers
	stores  stores.Stores
}

func NewInjector(srvs api.Servers, strs stores.Stores) *Injector {
	return &Injector{servers: srvs, stores: strs}
}

func (i *Injector) Inject(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, _ error) {
	ctx = servers.WithServers(ctx, i.servers)
	ctx = stores.WithStores(ctx, i.stores)
	return handler(ctx, req)
}
