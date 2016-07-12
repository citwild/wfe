package backend

import (
	"golang.org/x/net/context"
	"github.com/citwild/wfe/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"github.com/citwild/wfe/store"
)

var Accounts api.AccountsServer = &accounts{}

type accounts struct{}

func (s *accounts) Create(ctx context.Context, newAcct *api.NewAccount) (*api.CreatedAccount, error) {
	if newAcct.Login == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "empty login")
	}

	newUser := &api.User{Login: newAcct.Login, UID: newAcct.UID}

	var email *api.EmailAddr
	if newAcct.Email != "" {
		email = &api.EmailAddr{Email: newAcct.Email}
	}

	created, err := store.AccountsFromContext(ctx).Create(ctx, newUser, email)
	if err != nil {
		return nil, err
	}

	return &api.CreatedAccount{UID: created.UID}, nil;
}