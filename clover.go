package clover

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/mayron1806/go-clover-core/db"
	"github.com/mayron1806/go-clover-core/logger"
)

type Clover struct {
	server *Server
	logger *logger.Logger
	db     *db.Database
}

func (c *Clover) Run() {
	if c.db != nil {
		defer c.db.Close()
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		if err := c.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			c.logger.Errorf("Server error: %v", err)
		}
	}()

	<-stop

	c.logger.Infof("Shutting down server...")

	if err := c.server.Shutdown(context.Background()); err != nil {
		c.logger.Errorf("Server error: %v", err)
	}

	c.logger.Infof("Server stopped")
}

func (c *Clover) ConfigureServer(server *http.Server, force bool) *Server {
	if c.server == nil || force {
		c.server = NewServer(server)
	}
	return c.server
}
func (c *Clover) Router() *Router {
	return c.server.router
}
func NewClover() *Clover {
	logger := logger.NewLogger(logger.LoggerOptions{
		Prefix: "CLOVER",
	})
	fmt.Printf("%s\n", logo)
	return &Clover{
		logger: logger,
	}
}
