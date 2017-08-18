package api

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server   *httptest.Server
	reader   io.Reader
	usersUrl string
)

func init() {
	server = httptest.NewServer(Handlers())
	usersUrl = fmt.Sprintf("%s/users", server.URL)
}

func TestListUsersWithSuccess200(t *testing.T) {
	reader = strings.NewReader("")
	request, err := http.NewRequest("GET", usersUrl, reader)
	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
