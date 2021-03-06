package service

import (
	"github.com/citwild/wfe/api"
	"golang.org/x/net/context"
)

type key int

const (
	accountsKey key = iota
)

func WithServers(ctx context.Context, s api.Servers) context.Context {
	if s.Accounts != nil {
		ctx = WithAccounts(ctx, s.Accounts)
	}
	return ctx
}

func WithAccounts(ctx context.Context, s api.AccountsServer) context.Context {
	return context.WithValue(ctx, accountsKey, s)
}

func Accounts(ctx context.Context) api.AccountsServer {
	s, ok := ctx.Value(accountsKey).(api.AccountsServer)
	if !ok || s == nil {
		panic("no Accounts service set in context")
	}
	return s
}
