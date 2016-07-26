package store

import "golang.org/x/net/context"

type key int

const (
	accountsKey key = iota
	passwordKey
)

func WithStores(ctx context.Context, s Stores) context.Context {
	if s.Accounts != nil {
		ctx = WithAccounts(ctx, s.Accounts)
	}
	if s.Password != nil {
		ctx = WithPassword(ctx, s.Password)
	}
	return ctx
}

func WithAccounts(ctx context.Context, s AccountsStore) context.Context {
	return context.WithValue(ctx, accountsKey, s)
}

func Accounts(ctx context.Context) AccountsStore {
	s, ok := ctx.Value(accountsKey).(AccountsStore)
	if !ok || s == nil {
		panic("no Accounts store set in context")
	}
	return s
}

func WithPassword(ctx context.Context, s PasswordStore) context.Context {
	return context.WithValue(ctx, passwordKey, s)
}

func Password(ctx context.Context) PasswordStore {
	s, ok := ctx.Value(passwordKey).(PasswordStore)
	if !ok || s == nil {
		panic("no Password store set in context")
	}
	return s
}
