package mongostore

import (
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/store"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type dbUser struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Login string
	Name  string
}

func (du *dbUser) fromUser(u *api.User) {
	du.ID = bson.ObjectId(u.ID)
	du.Login = u.Login
	du.Name = u.Name
}

func (du *dbUser) toUser() *api.User {
	return &api.User{
		ID:    string(du.ID),
		Login: du.Login,
		Name:  du.Name,
	}
}

type UsersStore struct {
	session *mgo.Session
}

var _ store.UsersStore = (*UsersStore)(nil)

func NewUsersStore(s *mgo.Session) *UsersStore {
	return &UsersStore{session: s}
}
