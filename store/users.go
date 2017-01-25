package store

import (
	"fmt"
	"github.com/citwild/wfe/api"
	"golang.org/x/net/context"
)

type UsersStore interface {
	Get(ctx context.Context, spec api.UserSpec) (*api.User, error)
}

type UserNotFoundError struct {
	Login string
	ID    string
}

func (e *UserNotFoundError) Error() string {
	if e.Login != "" {
		return fmt.Sprintf("user %s not found", e.Login)
	} else {
		return fmt.Sprintf("user %s not found", e.ID)
	}
}
