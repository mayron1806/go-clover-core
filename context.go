package clover

import "net/http"

type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
	Params  map[string]string
}

func newContext(w http.ResponseWriter, r *http.Request, p map[string]string) *Context {
	return &Context{
		Request: r,
		Writer:  w,
		Params:  p,
	}
}

type HandlerFunc func(*Context)

func (f HandlerFunc) ServeHTTP(ctx *Context) {
	f(ctx)
}
