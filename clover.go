package clover

import (
	"fmt"
	"net/http"

	"github.com/mayron1806/go-clover-core/cloverlog"
)

var logger = cloverlog.NewLogger(cloverlog.LoggerOptions{
	HideTime:   true,
	HidePrefix: true,
})

type Clover struct {
	Router
	server *http.Server
}

func (f *Clover) Run() {
	logger.Info(fmt.Sprintf("Clover server running on %s", f.server.Addr))
	f.server.ListenAndServe()
}

func NewCloverServer(server *http.Server) *Clover {
	logger.Info(logo)
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
	return &Clover{
		server: server,
	}
}
