package server

import (
	"net/http"
)

type Fast struct {
	Router
	server *http.Server
}

func (f *Fast) Run() {
	f.server.ListenAndServe()
}

func NewFastServer(server *http.Server) *Fast {
	mux := http.NewServeMux()
	if server == nil {
		server = &http.Server{
			Addr:    ":8080",
			Handler: mux,
		}
	}
	if server.Handler == nil {
		server.Handler = mux
	}
	return &Fast{
		server: server,
	}
}
