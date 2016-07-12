package stores

import "golang.org/x/net/context"

type Password interface {
	SetPassword(ctx context.Context, UID int32, password string) error
}
