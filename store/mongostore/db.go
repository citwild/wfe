package mongostore

import "gopkg.in/mgo.v2"

type db struct {
	session *mgo.Session
}

func newDB(s *mgo.Session) *db {
	return &db{session: s.Copy()}
}

func (d *db) Close() {
	d.session.Close()
}

func (d *db) C(name string) *mgo.Collection {
	return d.session.DB("wfe").C(name)
}
