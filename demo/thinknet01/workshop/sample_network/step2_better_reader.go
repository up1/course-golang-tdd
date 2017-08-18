package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RepoInfo struct {
	Id       uint   `json:"id"`
	FullName string `json:"full_name"`
}

//Interface
type RepoInfoer interface {
	GetRepoInfoFromAPI(repo string) (string, error)
}

//Struct
type GithubRepoInfoer struct{}

func (gh GithubRepoInfoer) GetRepoInfoFromAPI(repo string) (string, error) {
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s", repo)
	response, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	info := RepoInfo{}

	if err := json.Unmarshal(body, &info); err != nil {
		return "", err
	}

	fullName := info.FullName

	return fullName, nil
}

func getRepoFullName(ri RepoInfoer, repo string) (string, error) {
	fullName, err := ri.GetRepoInfoFromAPI(repo)
	if err != nil {
		return "", fmt.Errorf("Error querying GitHub API: %s", err)
	}
	return fmt.Sprintf("Full name is %s", fullName), nil
}

func main() {
	gh := GithubRepoInfoer{}
	fullName, err := getRepoFullName(gh, "golang/go")
	if err != nil {
		fmt.Printf("Got error %v\n", err)
	}

	fmt.Println(fullName)
}
