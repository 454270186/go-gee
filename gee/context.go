package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

// NewContext initializes a new context
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

// get post data from form
// parse URL parameter
// set response header code
// set response header
// response a string
// response a JSON
// response data
// render html

func (ctx *Context) PostForm(key string) string {
	return ctx.Req.FormValue(key)
}

// Query parsers the raw query
func (ctx *Context) Query(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

// Status sends an HTTP response header with the provided status code
func (ctx *Context) Status(code int) {
	ctx.StatusCode = code
	ctx.Writer.WriteHeader(code)
}

// SetHeader inserts a key-value into the header map that will be sent by WriteHeader()
func (ctx *Context) SetHeader(key string, value string) {
	ctx.Writer.Header().Set(key, value)
}

func (ctx *Context) String(code int, format string, values ...interface{}) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Status(code)
	ctx.Writer.Write([]byte(fmt.Sprintf(format, values)))
}

func (ctx *Context) JSON(code int, obj interface{}) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)

	encoder := json.NewEncoder(ctx.Writer)
	err := encoder.Encode(obj)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (ctx *Context) Data(code int, data string) {
	ctx.Status(code)
	ctx.Writer.Write([]byte(data))
}

func (ctx *Context) HTML(code int, html string) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Status(code)
	ctx.Writer.Write([]byte(html))
}
