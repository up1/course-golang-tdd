package better_dao

import (
	"gopkg.in/mgo.v2"
	"readmail/model"
	"readmail/parser"
)

type Mongo struct {
	session *mgo.Session
	dbName  string
}

func (m Mongo) db() *mgo.Database {
	session := m.session.Copy()
	session.SetMode(mgo.Monotonic, true)
	return session.DB(m.dbName)
}

func (m Mongo) Save(mail parser.Mail) error {
	err := m.db().C("mail").Insert(&model.XMail{From: mail.From, To: mail.To, Subject: mail.Subject})
	return err
}

func (m Mongo) drop() error {
	err := m.db().DropDatabase()
	return err
}

func NewMongo(dbURI, dbName string) (Mongo, error) {
	session, err := mgo.Dial(dbURI)
	return Mongo{
		session: session,
		dbName:  dbName,
	}, err
}
