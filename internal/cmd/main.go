package main

import (
	"context"
	"github.com/AndreySirin/time/internal/logger"
	"github.com/AndreySirin/time/internal/server"
	"github.com/AndreySirin/time/internal/storage"
)

const (
	address  = "127.0.0.1:5432"
	username = "postgres"
	password = "postgres"
	database = "postgres"

	httpPort = ":8080"
)

func main() {
	lg := logger.New()
	lg.Info("Starting server...")

	psql, err := storage.New(lg, username, password, address, database)
	if err != nil {
		lg.Error("Failed to connect to database",
			"error", err)
		return
	}

	defer func() {
		if err = psql.Close(); err != nil {
			lg.Error("Failed to close",
				"error", err)
		}
	}()

	if err = psql.DummyMigration(context.Background()); err != nil {
		lg.Error("Failed to migrate",
			"error", err)

		return
	}

	httpServer := server.New(lg, httpPort)

	if err = httpServer.Run(); err != nil {
		lg.Error("Server failed to start", "error", err)

		return
	}

	lg.Info("Shutting down...")
}
