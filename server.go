package clover

import (
	"context"
	"net/http"
	"strings"
)

type Server struct {
	httpServer *http.Server
	router     *Router
}

func (s *Server) ListenAndServe() {
	s.httpServer.ListenAndServe()
}
func (s *Server) Shutdown(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}

type CORSOptions struct {
	AllowedOrigins []string
	AllowedHeaders []string
	AllowedMethods []string
}

func (s *Server) AddCors(opts CORSOptions) {
	s.router.route.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", strings.Join(opts.AllowedOrigins, ","))
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(opts.AllowedMethods, ","))
		w.Header().Set("Access-Control-Allow-Headers", strings.Join(opts.AllowedHeaders, ","))

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
	})
}

func (s *Server) Router() *Router {
	return s.router
}

func NewServer(server *http.Server, router *Router) *Server {
	return &Server{
		httpServer: server,
		router:     router,
	}
}
