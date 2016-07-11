package backend

import (
	"golang.org/x/net/context"
	"github.com/citwild/wfe/api/wfe"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"github.com/citwild/wfe/pkg/store"
)

var Accounts wfe.AccountsServer = &accounts{}

type accounts struct{}

func (s *accounts) Create(ctx context.Context, newAcct *wfe.NewAccount) (*wfe.CreatedAccount, error) {
	if newAcct.Login == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "empty login")
	}

	newUser := &wfe.User{Login: newAcct.Login, UID: newAcct.UID}

	var email *wfe.EmailAddr
	if newAcct.Email != "" {
		email = &wfe.EmailAddr{Email: newAcct.Email}
	}

	created, err := store.AccountsFromContext(ctx).Create(ctx, newUser, email)
	if err != nil {
		return nil, err
	}

	return &wfe.CreatedAccount{UID: created.UID}, nil;
}