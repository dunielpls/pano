package server

import (
	"os"
	"os/signal"
	"syscall"
)

func (srv *Server) initInterrupt() {
	// Catches SIGQUIT/SIGTERM and shuts down the server when one is received.
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		sig := <-signals

		srv.log.Printf("Caught signal %s, shutting down...", sig)
		srv.Stop()

		os.Exit(0)
	}()
}
