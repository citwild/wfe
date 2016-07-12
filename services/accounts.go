package services

import (
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/stores"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var Accounts api.AccountsServer = &accounts{}

type accounts struct{}

func (s *accounts) Create(ctx context.Context, newAcct *api.NewAccount) (*api.CreatedAccount, error) {
	if newAcct.Login == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "empty login")
	}

	newUser := &api.User{Login: newAcct.Login, UID: newAcct.UID}

	var email *api.EmailAddress
	if newAcct.Email != "" {
		email = &api.EmailAddress{Email: newAcct.Email}
	}

	created, err := stores.AccountsFromContext(ctx).Create(ctx, newUser, email)
	if err != nil {
		return nil, err
	}

	err = stores.PasswordFromContext(ctx).SetPassword(ctx, created.UID, newAcct.Password)
	if err != nil {
		return nil, err
	}

	return &api.CreatedAccount{UID: created.UID}, nil
}
