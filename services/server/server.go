package server

import (
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/services"
	"github.com/citwild/wfe/services/internal/middleware"
	"google.golang.org/grpc"
)

func New() *grpc.Server {
	s := grpc.NewServer(grpc.UnaryInterceptor(middleware.InitContext))
	RegisterAll(s, services.NewServices())
	return s
}

func RegisterAll(s *grpc.Server, svcs services.Services) {
	if svcs.Accounts != nil {
		api.RegisterAccountsServer(s, svcs.Accounts)
	}
}
