package clover

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	router     *Router
}

func (s *Server) ListenAndServe() error {
	s.httpServer.Handler = s.router.route
	return http.ListenAndServe(s.httpServer.Addr, s.router.route)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) Router() *Router {
	return s.router
}

func NewServer(server *http.Server) *Server {
	if server == nil {
		server = &http.Server{
			Addr: ":8080",
		}
	}
	return &Server{
		httpServer: server,
		router:     NewRouter(NewServeClover(), ""),
	}
}
