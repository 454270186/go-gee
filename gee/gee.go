package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(c *Context)

// Engine implement the interface of http.Handler
type Engine struct {
	router *Router
}

// New initializes a engine
func New() *Engine {
	return &Engine{
		router: NewRouter(),
	}
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.router.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.router.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) {
	err := http.ListenAndServe(addr, engine)
	if err != nil {
		log.Println(err)
	}
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := NewContext(w, r)
	engine.router.handle(context)
}
