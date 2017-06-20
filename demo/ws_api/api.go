package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/hello", hello2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error with ", err)
	}
}

func hello2(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		fmt.Fprintf(res, "Call with %v", req.Method)
		// TODO
		return
	}
	res.WriteHeader(http.StatusBadRequest)

}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello World"))
	fmt.Fprintf(res, "Hello World")
}


