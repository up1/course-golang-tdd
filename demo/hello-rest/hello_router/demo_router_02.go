package main

import (
	"net/http"
	"regexp"
)


type route struct {
	pattern *regexp.Regexp
	verb    string
	handler http.Handler
}

type RegexpHandler struct {
	routes []*route
}

func (h *RegexpHandler) Handler(pattern *regexp.Regexp, verb string, handler http.Handler) {
	h.routes = append(h.routes, &route{pattern, verb, handler})
}

func (h *RegexpHandler) HandleFunc(r string, v string, handler func(http.ResponseWriter, *http.Request)) {
	re := regexp.MustCompile(r)
	h.routes = append(h.routes, &route{re, v, http.HandlerFunc(handler)})
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) && route.verb == r.Method {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {

	regHandler := new(RegexpHandler)

	regHandler.HandleFunc("/todo/$", "GET", ListOfTODO)
	regHandler.HandleFunc("/todo/$", "POST", CreateTOFO)
	regHandler.HandleFunc("/todo/[0-9]+$", "GET", GetTODOByID)
	regHandler.HandleFunc("/todo/[0-9]+$", "PUT", UpdateTODOByID)
	regHandler.HandleFunc("/todo/[0-9]+$", "DELETE", DeleteTODOByID)

	http.ListenAndServe(":8080", nil)
}
