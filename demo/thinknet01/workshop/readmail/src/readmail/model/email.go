package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type XMail struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	From      string
	To        string
	Subject   string
	Timestamp time.Time
}
