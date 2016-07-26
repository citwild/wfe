package store

import "golang.org/x/net/context"

type PasswordStore interface {
	SetPassword(ctx context.Context, UID int32, password string) error
}
