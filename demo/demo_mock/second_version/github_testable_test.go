package main

import "testing"

type FakeRelease struct {
	Tag string
	Err error
}

func (f FakeRelease) GetLatestTag(repo string) (string, error) {
	if f.Err != nil {
		return "", f.Err
	}

	return f.Tag, nil

}

func TestGetLatestTag(t *testing.T) {
	f := FakeRelease{
		Tag: "1.0.0",
		Err: nil,
	}

	expected := "1.0.0"
	tag, _ := getResult(f, "xxx")
	if tag != expected {
		t.Fatalf("Expectd %s but got %s", expected, tag)
	}
}