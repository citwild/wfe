package service

import (
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/store"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type AccountsServer struct{}

var _ api.AccountsServer = (*AccountsServer)(nil)

func NewAccountsServer() *AccountsServer {
	return &AccountsServer{}
}

func (s *AccountsServer) Create(ctx context.Context, newAcct *api.NewAccount) (*api.CreatedAccount, error) {
	if newAcct.Login == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "empty login")
	}

	newUser := &api.User{Login: newAcct.Login, UID: newAcct.UID}

	var email *api.EmailAddress
	if newAcct.Email != "" {
		email = &api.EmailAddress{Email: newAcct.Email}
	}

	created, err := store.Accounts(ctx).Create(ctx, newUser, email)
	if err != nil {
		return nil, err
	}

	err = store.Password(ctx).SetPassword(ctx, created.UID, newAcct.Password)
	if err != nil {
		return nil, err
	}

	return &api.CreatedAccount{UID: created.UID}, nil
}
