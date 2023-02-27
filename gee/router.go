package gee

import (
	"fmt"
	"log"
	"net/http"
)

// Router controls the router-map
type Router struct {
	handlers map[string]HandlerFunc // handlers is the router-map table
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
	}
}

// addRouter inserts URL-path into router-map table
func (router *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	URL := method + "-" + pattern
	router.handlers[URL] = handler
}

func (router *Router) handle(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path

	if handler, ok := router.handlers[key]; ok {
		handler(w, r)
	} else {
		_, err := fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL.Path)
		if err != nil {
			log.Println(err)
		}
	}
}
