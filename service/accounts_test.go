package service

import (
	"testing"

	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/store"
	"github.com/citwild/wfe/store/mockstore"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
	"time"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := testContext(ctrl)

	acct := store.Accounts(ctx).(*mockstore.MockAccountsStore)
	acct.EXPECT().Create(ctx, gomock.Any(), &api.EmailAddress{Email: "e@mail.com"}).
		Return(&api.User{ID: "id"}, nil)

	pass := store.Password(ctx).(*mockstore.MockPasswordStore)
	pass.EXPECT().SetPassword(ctx, "id", "pass")

	actual, err := NewAccountsServer().Create(ctx, &api.NewAccount{Login: "me", Password: "pass", Email: "e@mail.com"})
	if err != nil {
		t.Fatal(err)
	}

	expected := &api.CreatedAccount{ID: "id"}
	if *actual != *expected {
		t.Errorf("Account: expected %+v, actual %+v", expected, actual)
	}
}

func TestTimestamp(t *testing.T) {
	tm := time.Now()
	ts, _ := ptypes.TimestampProto(tm)
	tm2, _ := ptypes.Timestamp(ts)
	ts2, _ := ptypes.TimestampProto(tm2)

	if ts.Nanos != ts2.Nanos {
		t.Errorf("expected %v, actual %v", tm.Nanosecond(), tm2.Nanosecond())
	}
}
