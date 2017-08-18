package step1

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

func getRepoInfoFromAPI(repo string) (string, error) {
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

func getRepoFullName(repo string) (string, error) {
	fullName, err := getRepoInfoFromAPI(repo)
	if err != nil {
		return "", fmt.Errorf("Error querying GitHub API: %s", err)
	}
	return fmt.Sprintf("Full name is %s", fullName), nil
}

func main() {
	fullName, err := getRepoFullName("golang/go")
	if err != nil {
		fmt.Printf("Got error %v\n", err)
	}

	fmt.Println(fullName)
}
