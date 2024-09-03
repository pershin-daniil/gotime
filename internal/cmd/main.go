package main

import (
	"github.com/AndreySirin/time/internal/logger"
	"github.com/AndreySirin/time/internal/server"
)

const httpPort = ":8080"

func main() {
	lg := logger.New()
	lg.Info("Starting server...")

	httpServer := server.New(lg, httpPort)

	if err := httpServer.Run(); err != nil {
		lg.Error("Server failed to start", "error", err)

		return
	}
}
