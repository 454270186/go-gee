package gee

import "net/http"

type H map[string]interface{}

// Context encapsulates:
// two objects: Request and ResponseWriter
// carried info: request info and response info
type Context struct {
	// http Objects
	Writer http.ResponseWriter
	Req    *http.Request

	// request info
	Path   string
	Method string

	// response info
	StatusCode int
}
