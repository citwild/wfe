package store

import "golang.org/x/net/context"

type contextKey int

const (
	accountKey contextKey = iota
)

func WithAccounts(ctx context.Context, s Accounts) context.Context {
	return context.WithValue(ctx, accountKey, s)
}

func AccountsFromContext(ctx context.Context) Accounts {
	s, ok := ctx.Value(accountKey).(Accounts)
	if !ok || s == nil {
		panic("no Accounts set in context")
	}
	return s
}
