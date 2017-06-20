package main

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Release interface {
	GetLatestTag(string) (string, error)
}

type ReleaseInformation struct {
	Id int `json:"id"`
	TagName string `json:"tag_name"`
}

type GithubRelease struct {
}

func (github GithubRelease) GetLatestTag(repo string) (string, error) {
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)
	response, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	releases := []ReleaseInformation{}

	if err := json.Unmarshal(body, &releases); err != nil {
		return "", err
	}

	return releases[0].TagName, nil
}

func getResult(r Release, repo string) (string, error) {
	tag, err := r.GetLatestTag(repo)
	if err != nil {
		return "", fmt.Errorf("Error from github %s", err)
	}

	return tag, nil
}



func main() {
	github := GithubRelease{}
	latestTag, err := getResult( github, "docker/machine")
	if err != nil {
		fmt.Println(os.Stderr, err)
	}

	fmt.Println(latestTag)
}
