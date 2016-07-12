package stores

import (
	"github.com/citwild/wfe/api"
	"golang.org/x/net/context"
)

type Accounts interface {
	Create(ctx context.Context, newUser *api.User, email *api.EmailAddr) (*api.User, error)
}
