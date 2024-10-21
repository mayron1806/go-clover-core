package clover

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mayron1806/go-clover-core/cloverdb"
	"github.com/mayron1806/go-clover-core/cloverlog"
)

type Clover struct {
	router *Router
	server *http.Server
	logger *cloverlog.Logger
	db     *cloverdb.Database
}

func (c *Clover) Run() {
	if c.db != nil {
		defer c.db.Close()
	}
	c.logger.Info(fmt.Sprintf("Clover server running on %s", c.server.Addr))
	c.server.ListenAndServe()
}

func (c *Clover) ConfigureServer(server *http.Server) {
	r := httprouter.New()
	if server == nil {
		server = &http.Server{
			Addr:    ":8080",
			Handler: r,
		}
	}
	server.Handler = r
	c.server = server

	c.router = NewRouter(r, "/")
}
func (c *Clover) Router() *Router {
	return c.router
}
func NewClover() *Clover {
	logger := cloverlog.NewLogger(cloverlog.LoggerOptions{
		HideTime:   true,
		HidePrefix: true,
	})
	logger.Info(logo)
	return &Clover{
		logger: logger,
	}
}
