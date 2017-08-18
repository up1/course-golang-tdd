package main

import (
	"testing"
)

type MockRepoInfoer struct{}

func (mr MockRepoInfoer) GetRepoInfoFromAPI(repo string) (string, error) {
  return repo, nil
}

func TestGetFullnamefromGithubRepository(t *testing.T) {
	expectedFullname := "demo/test"
	mock := MockRepoInfoer{}
	f, err := getRepoFullName(mock, expectedFullname)
	if err != nil {
		t.Fatalf("Got error %v", err)
	}

	if f != "Full name is demo/test" {
		t.Fatalf("Expected %v but got %v", expectedFullname, f)
	}

}
