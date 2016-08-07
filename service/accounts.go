package service

import (
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/store"
	"github.com/golang/protobuf/ptypes/timestamp"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
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

	t := time.Now()
	now := timestamp.Timestamp{Seconds: int64(t.Second()), Nanos: int32(t.Nanosecond())}
	newUser := &api.User{
		Login:        newAcct.Login,
		RegisteredAt: &now,
	}

	var email *api.EmailAddress
	if newAcct.Email != "" {
		email = &api.EmailAddress{Email: newAcct.Email}
	}

	created, err := store.Accounts(ctx).Create(ctx, newUser, email)
	if err != nil {
		return nil, err
	}

	err = store.Password(ctx).SetPassword(ctx, created.ID, newAcct.Password)
	if err != nil {
		return nil, err
	}

	return &api.CreatedAccount{ID: created.ID}, nil
}
