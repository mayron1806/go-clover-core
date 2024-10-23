package clover

import (
	"net/http"
	"strings"

	"github.com/mayron1806/go-clover-core/logger"
)

type Router struct {
	route       *ServeClover
	subRouters  []*Router
	prefix      string
	middlewares []IMiddleware
	logger      *logger.Logger
}

// Applymiddleware.IMiddlewares applies middlewares to all routes in the group
func (r *Router) Applymiddleware(middlewares ...IMiddleware) {
	r.middlewares = append(r.middlewares, middlewares...)
}

// Use applies middlewares to a specific route in the group based on method and path
func (r *Router) use(method, pattern string, handler HandlerFunc, middlewares ...IMiddleware) {
	finalHandler := handler

	// apply middlewares in reverse order
	for i := len(middlewares) - 1; i >= 0; i-- {
		finalHandler = middlewares[i](finalHandler)
	}
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		finalHandler = r.middlewares[i](finalHandler)
	}

	// Register the route with method handling
	r.route.Handle(method, pattern, finalHandler)
}

// GET maps a route to the GET method
func (f *Router) GET(pattern string, handler HandlerFunc, middlewares ...IMiddleware) {
	f.logger.Infof("Registering route of type GET with path %s", f.mountRoute(pattern))
	f.use(http.MethodGet, f.mountRoute(pattern), handler, middlewares...)
}

// POST maps a route to the POST method
func (f *Router) POST(pattern string, handler HandlerFunc, middlewares ...IMiddleware) {
	f.logger.Infof("Registering route of type POST with path %s", f.mountRoute(pattern))
	f.use(http.MethodPost, f.mountRoute(pattern), handler, middlewares...)
}

// PUT maps a route to the PUT method
func (f *Router) PUT(pattern string, handler HandlerFunc, middlewares ...IMiddleware) {
	f.logger.Infof("Registering route of type PUT with path %s", f.mountRoute(pattern))
	f.use(http.MethodPut, f.mountRoute(pattern), handler, middlewares...)
}

// PATCH maps a route to the PATCH method
func (f *Router) PATCH(pattern string, handler HandlerFunc, middlewares ...IMiddleware) {
	f.logger.Infof("Registering route of type PATCH with path %s", f.mountRoute(pattern))
	f.use(http.MethodPatch, f.mountRoute(pattern), handler, middlewares...)
}

// DELETE maps a route to the DELETE method
func (f *Router) DELETE(pattern string, handler HandlerFunc, middlewares ...IMiddleware) {
	f.logger.Infof("Registering route of type DELETE with path %s", f.mountRoute(pattern))
	f.use(http.MethodDelete, f.mountRoute(pattern), handler, middlewares...)
}

// HEAD maps a route to the HEAD method
func (f *Router) HEAD(pattern string, handler HandlerFunc, middlewares ...IMiddleware) {
	f.logger.Infof("Registering route of type HEAD with path %s", f.mountRoute(pattern))
	f.use(http.MethodHead, f.mountRoute(pattern), handler, middlewares...)
}

// CONNECT maps a route to the CONNECT method
func (f *Router) CONNECT(pattern string, handler HandlerFunc, middlewares ...IMiddleware) {
	f.logger.Infof("Registering route of type CONNECT with path %s", f.mountRoute(pattern))
	f.use(http.MethodConnect, f.mountRoute(pattern), handler, middlewares...)
}

// OPTIONS maps a route to the OPTIONS method
func (f *Router) OPTIONS(pattern string, handler HandlerFunc, middlewares ...IMiddleware) {
	f.logger.Infof("Registering route of type OPTIONS with path %s", f.mountRoute(pattern))
	f.use(http.MethodOptions, f.mountRoute(pattern), handler, middlewares...)
}

// TRACE maps a route to the TRACE method
func (f *Router) TRACE(pattern string, handler HandlerFunc, middlewares ...IMiddleware) {
	f.logger.Infof("Registering route of type TRACE with path %s", f.mountRoute(pattern))
	f.use(http.MethodTrace, f.mountRoute(pattern), handler, middlewares...)
}

func (f *Router) AddSubRoute(pattern string) *Router {
	subRouter := NewRouter(f.route, f.mountRoute(pattern))
	f.subRouters = append(f.subRouters, subRouter)
	return subRouter
}

func (f *Router) mountRoute(pattern string) string {
	result := strings.ReplaceAll(("/" + f.prefix + pattern), "//", "/")
	result = strings.TrimSuffix(result, "/")
	return result
}

func NewRouter(c *ServeClover, prefix string) *Router {
	return &Router{
		route:  c,
		prefix: prefix,
		logger: logger.NewLogger(logger.LoggerOptions{Prefix: "ROUTER"}),
	}
}
