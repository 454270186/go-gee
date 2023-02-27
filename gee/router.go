package gee

import "log"

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
	log.Printf("Route %4s - %s", method, pattern)
	URL := method + "-" + pattern
	router.handlers[URL] = handler
}

func (router *Router) handle(c *Context) {
	key := c.Method + "-" + c.Path

	if handler, ok := router.handlers[key]; ok {
		handler(c)
	} else {
		log.Printf("404 NOT FOUND: %s\n", c.Path)
		c.String(404, "404 NOT FOUND: %s\n", c.Path)
	}
}
