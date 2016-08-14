package mongostore

import (
	"testing"
	"time"

	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/test/testdb"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes"
)

func TestCreate_lg(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	db := testdb.New()
	err := db.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	s, err := db.NewSession()
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := testContext(ctrl)

	now, _ := ptypes.TimestampProto(time.Now())
	user := &api.User{Login: "me", Name: "my name", RegisteredAt: now}
	email := &api.EmailAddress{Email: "e@mail.com"}

	created, err := NewAccountsStore(s).Create(ctx, user, email)
	if err != nil {
		t.Fatal(err)
	}

	if created.Login != user.Login {
		t.Errorf("Login: expected %v, actual %v", user.Login, created.Login)
	}
	if created.Name != user.Name {
		t.Errorf("Name: expected %v, actual %v", user.Name, created.Name)
	}

	got, err := NewUsersStore(s).Get(ctx, api.UserSpec{Login: "me"})
	if err != nil {
		t.Fatal(err)
	}

	if got.Login != created.Login {
		t.Errorf("Login: expected %v, actual %v", created.Login, got.Login)
	}
	if got.Name != created.Name {
		t.Errorf("Name: expected %v, actual %v", created.Name, got.Name)
	}
	// MongoDB stores timestamps with millisecond precision
	if got.RegisteredAt.Seconds != created.RegisteredAt.Seconds {
		t.Errorf("RegisteredAt: expected %v, actual %v", created.RegisteredAt.Seconds, got.RegisteredAt.Seconds)
	}
}
