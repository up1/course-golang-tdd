package parser

import (
	"path/filepath"
	"testing"
)

func TestReadGoodFile(t *testing.T) {
	expected := Mail{}
	expected.From = "<tonluck7@gmail.com>"
	expected.To = "proptech11_252@trustmail.jobthai.com"
	expected.Subject = "สมัครงาน"

	path := filepath.Join("testdata", "success")
	actual := ReadFromFile(path)
	if actual != expected {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}
