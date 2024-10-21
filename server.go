package clover

import (
	"context"
	"net/http"

	"github.com/mayron1806/go-clover-core/cors"
)

type Server struct {
	httpServer *http.Server
	router     *Router
}

func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) AddCors(opts cors.CORSOptions) *Server {
	s.router.route.GlobalOPTIONS = cors.Cors(opts)
	return s
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
