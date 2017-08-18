package better_dao

import (
	"readmail/parser"
	"testing"
)

var mongo Mongo

func init() {
	mongo, _ = NewMongo("127.0.0.1", "testing")
}

func TestBetterSaveDataToMongoDBSuccess(t *testing.T) {
	mongo.drop()

	m := parser.Mail{}
	result := mongo.Save(m)
	if result != nil {
		t.Fatalf("Can not save data to mongodb")
	}
}
