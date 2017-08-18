package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"readmail/model"
	"readmail/parser"
)

func Save(m parser.Mail) error {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		return err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("mail")

	err = c.Insert(&model.XMail{From: m.From, To: m.To, Subject: m.Subject})
	return err
}

func Find() ([]model.XMail, error) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		return nil, err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("mail")

	query := bson.M{}
	var result []model.XMail
	err = c.Find(query).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func drop() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	err = session.DB("test").DropDatabase()
	if err != nil {
		panic(err)
	}
}

func count() (int, error) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		return 0, err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("mail")
	return c.Count()
}
