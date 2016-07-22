package middleware

import (
	"github.com/citwild/wfe/services"
	"github.com/citwild/wfe/stores"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Injector struct {
	services services.Services
	stores   stores.Stores
}

func NewInjector(svcs services.Services, strs stores.Stores) *Injector {
	return &Injector{services: svcs, stores: strs}
}

func (i *Injector) Inject(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, _ error) {
	ctx = services.WithServices(ctx, i.services)
	ctx = stores.WithStores(ctx, i.stores)
	return handler(ctx, req)
}
