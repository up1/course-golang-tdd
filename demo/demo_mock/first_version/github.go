package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
)

type ReleaseInfomation struct {
	Id int `json:"id"`
	TagName string `json:"tag_name"`
}

func getLatestTag(repo string) (string, error) {
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

	releases := []ReleaseInfomation{}

	if err := json.Unmarshal(body, &releases); err != nil {
		return "", err
	}

	return releases[0].TagName, nil
}

func main() {
	latestTag, err := getLatestTag("docker/machine")
	if err != nil {
		fmt.Println(os.Stderr, err)
	}

	fmt.Println(latestTag)
}
