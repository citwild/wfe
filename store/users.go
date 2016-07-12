package store

import (
	"golang.org/x/net/context"
	"github.com/citwild/wfe/api"
)

type Accounts interface {
	Create(ctx context.Context, newUser *api.User, email *api.EmailAddr) (*api.User, error)
}
