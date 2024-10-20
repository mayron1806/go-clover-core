package clover

import (
	"github.com/julienschmidt/httprouter"
)

type Router struct {
	router      *httprouter.Router
	prefix      string
	middlewares []Middleware
}

// ApplyMiddlewares applies middlewares to all routes in the group
func (r *Router) ApplyMiddlewares(middlewares ...Middleware) {
	r.middlewares = append(r.middlewares, middlewares...)
}

// Use applies middlewares to a specific route in the group based on method and path
func (r *Router) use(method, pattern string, handler httprouter.Handle, middlewares ...Middleware) {
	finalHandler := handler

	// Apply group-level middlewares first
	for _, mw := range r.middlewares {
		finalHandler = mw(finalHandler)
	}

	// Apply route-specific middlewares
	for _, mw := range middlewares {
		finalHandler = mw(finalHandler)
	}

	// Register the final handler with the method and pattern
	r.router.Handle(method, r.prefix+pattern, finalHandler)
}

// GET maps a route to the GET method
func (f *Router) GET(pattern string, handler httprouter.Handle, middlewares ...Middleware) {
	f.use("GET", f.prefix+pattern, handler, middlewares...)
}

// POST maps a route to the POST method
func (f *Router) POST(pattern string, handler httprouter.Handle, middlewares ...Middleware) {
	f.use("POST", f.prefix+pattern, handler, middlewares...)
}

// PUT maps a route to the PUT method
func (f *Router) PUT(pattern string, handler httprouter.Handle, middlewares ...Middleware) {
	f.use("PUT", f.prefix+pattern, handler, middlewares...)
}

// PATCH maps a route to the PATCH method
func (f *Router) PATCH(pattern string, handler httprouter.Handle, middlewares ...Middleware) {
	f.use("PATCH", f.prefix+pattern, handler, middlewares...)
}

// DELETE maps a route to the DELETE method
func (f *Router) DELETE(pattern string, handler httprouter.Handle, middlewares ...Middleware) {
	f.use("DELETE", f.prefix+pattern, handler, middlewares...)
}

// HEAD maps a route to the HEAD method
func (f *Router) HEAD(pattern string, handler httprouter.Handle, middlewares ...Middleware) {
	f.use("HEAD", f.prefix+pattern, handler, middlewares...)
}

// CONNECT maps a route to the CONNECT method
func (f *Router) CONNECT(pattern string, handler httprouter.Handle, middlewares ...Middleware) {
	f.use("CONNECT", f.prefix+pattern, handler, middlewares...)
}

// OPTIONS maps a route to the OPTIONS method
func (f *Router) OPTIONS(pattern string, handler httprouter.Handle, middlewares ...Middleware) {
	f.use("OPTIONS", f.prefix+pattern, handler, middlewares...)
}

// TRACE maps a route to the TRACE method
func (f *Router) TRACE(pattern string, handler httprouter.Handle, middlewares ...Middleware) {
	f.use("TRACE", f.prefix+pattern, handler, middlewares...)
}

func (f *Router) Group(pattern string) *RouteGroup {
	return NewRouter(f.prefix + pattern).Group(pattern)
}

func NewRouter(prefix string) *Router {
	return &Router{
		router: httprouter.New(),
		prefix: prefix,
	}
}
