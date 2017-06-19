package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/todo", ListOfTODO)
	http.HandleFunc("/todo", CreateTOFO)
	http.HandleFunc("/todo/1", GetTODOByID)
	http.HandleFunc("/todo/1", UpdateTODOByID)
	http.HandleFunc("/todo/1", DeleteTODOByID)
	http.ListenAndServe(":8080", nil)
}
