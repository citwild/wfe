package stores

import (
	"github.com/citwild/wfe/api"
	"golang.org/x/net/context"
)

type AccountsStore interface {
	Create(ctx context.Context, newUser *api.User, email *api.EmailAddress) (*api.User, error)
}
