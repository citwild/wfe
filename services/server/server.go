package server

import (
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/services"
	"github.com/citwild/wfe/services/internal/middleware"
	"google.golang.org/grpc"
)

func New() *grpc.Server {
	s := grpc.NewServer(grpc.UnaryInterceptor(middleware.InitContext))
	api.RegisterAccountsServer(s, services.Accounts)
	return s
}
