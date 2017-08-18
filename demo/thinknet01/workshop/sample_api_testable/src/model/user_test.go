package model

import (
	"testing"
)

func TestUser(t *testing.T) {
	var user = User{1, "somkiat", "pui", 30}
	if user == (User{}) {
		t.Fatal("Can not create")
	}

	if user.Id != 1 {
		t.Fatalf("Wrong of creation")
	}
}
