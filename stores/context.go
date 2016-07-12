package stores

import "golang.org/x/net/context"

type key int

const (
	accountKey key = iota
	passwordKey
)

func WithAccounts(ctx context.Context, s Accounts) context.Context {
	return context.WithValue(ctx, accountKey, s)
}

func AccountsFromContext(ctx context.Context) Accounts {
	s, ok := ctx.Value(accountKey).(Accounts)
	if !ok || s == nil {
		panic("no Accounts store set in context")
	}
	return s
}

func WithPassword(ctx context.Context, s Password) context.Context {
	return context.WithValue(ctx, passwordKey, s)
}

func PasswordFromContext(ctx context.Context) Password {
	s, ok := ctx.Value(passwordKey).(Password)
	if !ok || s == nil {
		panic("no Password store set in context")
	}
	return s
}
