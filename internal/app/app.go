// Package app configures and runs application.
package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/ironsail/whydah-go-clean-template/config"
	v1 "github.com/ironsail/whydah-go-clean-template/internal/controller/http/v1"
	"github.com/ironsail/whydah-go-clean-template/pkg/httpserver"
	"github.com/ironsail/whydah-go-clean-template/pkg/logger"
	"github.com/ironsail/whydah-go-clean-template/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	// Database
	pg, err := postgres.New(cfg)
	if err != nil {
		logger.Error("app - Run - postgres.New", logger.ErrWrap(err))
	}
	defer pg.Close()

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, cfg, pg)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		logger.Error("app - Run - httpServer.Notify", logger.ErrWrap(err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		logger.Error("app - Run - httpServer.Shutdown", logger.ErrWrap(err))
	}
}
