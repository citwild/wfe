package api

import (
	"google.golang.org/grpc"
)

func NewServer(srvs Servers, opt ...grpc.ServerOption) *grpc.Server {
	s := grpc.NewServer(opt...)

	if srvs.Accounts != nil {
		RegisterAccountsServer(s, srvs.Accounts)
	}

	return s
}
