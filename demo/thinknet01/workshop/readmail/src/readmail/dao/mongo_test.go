package dao

import (
	"readmail/parser"
	"testing"
)

func TestSaveDataToMongoDBSuccess(t *testing.T) {
	drop()

	m := parser.Mail{}
	result := Save(m)
	if result != nil {
		t.Fatalf("Can not save data to mongodb")
	}
}

func TestSaveAndCountData(t *testing.T) {
	drop()

	m := parser.Mail{From: "somkiat.p@gmail.com"}
	Save(m)

	count, err := count()
	if count != 1 {
		t.Fatalf("Can not find data to mongodb with %v", err)
	}
}

func TestSaveAndFindData(t *testing.T) {
	drop()

	m := parser.Mail{From: "somkiat.p@gmail.com"}
	Save(m)

	xm, err := Find()
	if err != nil {
		t.Fatalf("Can not find data to mongodb with %v", err)
	}

	if xm[0].From != m.From {
		t.Fatalf("Incorrect From %v", xm[0].From)
	}
}
