package mongostore

import (
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/store"
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type dbUser struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Login        string
	Name         string
	RegisteredAt time.Time
}

func (du *dbUser) fromUser(u *api.User) error {
	du.ID = bson.ObjectId(u.ID)
	du.Login = u.Login
	du.Name = u.Name
	t, err := ptypes.Timestamp(u.RegisteredAt)
	du.RegisteredAt = t
	return err
}

func (du *dbUser) toUser() (*api.User, error) {
	t, err := ptypes.TimestampProto(du.RegisteredAt)
	if err != nil {
		return nil, err
	}
	return &api.User{
		ID:           string(du.ID),
		Login:        du.Login,
		Name:         du.Name,
		RegisteredAt: t,
	}, nil
}

type UsersStore struct {
	session *mgo.Session
}

var _ store.UsersStore = (*UsersStore)(nil)

func NewUsersStore(s *mgo.Session) *UsersStore {
	return &UsersStore{session: s}
}

func (s *UsersStore) Get(ctx context.Context, spec api.UserSpec) (*api.User, error) {
	db := newDB(s.session)

	var u *api.User
	var err error
	if spec.ID != "" && spec.Login != "" {
		u, err = s.getByQuery(db, bson.M{"_id": spec.ID, "login": spec.Login})
	} else if spec.ID != "" {
		u, err = s.getByQuery(db, bson.M{"_id": spec.ID})
	} else if spec.Login != "" {
		u, err = s.getByQuery(db, bson.M{"login": spec.Login})
	} else {
		u, err = nil, &store.UserNotFoundError{}
	}
	return u, err
}

func (s *UsersStore) getByQuery(db *db, query interface{}) (*api.User, error) {
	var u dbUser
	err := db.C("users").Find(query).One(&u)
	if err != nil {
		return nil, err
	}
	return u.toUser()
}
