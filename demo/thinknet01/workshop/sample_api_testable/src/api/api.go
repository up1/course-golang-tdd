package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"model"
)

func Handlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", listUsersHandler).Methods("GET")
	return router
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := [2]model.User{}
	users[0] = model.User{1, "user 01", "last 01", 20}
	users[1] = model.User{2, "user 02", "last 02", 30}

	res, _ := json.Marshal(users)
	_, err := w.Write(res)
	if err != nil {
		panic(err)
	}
}
