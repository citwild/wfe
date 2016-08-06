package mongostore

import (
	"github.com/citwild/wfe/api"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCreate_lg(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := testContext(ctrl)

	user := &api.User{Login: "me", Name: "my name"}
	email := &api.EmailAddress{Email: "e@mail.com"}

	actual, err := NewAccountsStore(nil).Create(ctx, user, email)
	if err != nil {
		t.Fatal(err)
	}

	if actual.Login != user.Login {
		t.Errorf("Login: expected %v, actual %v", user.Login, actual.Login)
	}
	if actual.Name != user.Name {
		t.Errorf("Name: expected %v, actual %v", user.Name, actual.Name)
	}
}
