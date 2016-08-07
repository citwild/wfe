package store

import "golang.org/x/net/context"

type PasswordStore interface {
	SetPassword(ctx context.Context, ID string, password string) error
}
