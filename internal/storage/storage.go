package storage

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log/slog"
	"net/url"
)

const moduleName = "storage"

type Storage struct {
	lg *slog.Logger
	db *sql.DB
}

func New(
	lg *slog.Logger,
	username string,
	password string,
	address string,
	database string,
) (*Storage, error) {
	dsn := (&url.URL{
		Scheme: "postgresql",
		User:   url.UserPassword(username, password),
		Host:   address,
		Path:   database,
	}).String()

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("init db: %v", err)
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %v", err)
	}

	return &Storage{
		lg: lg.With("module", moduleName),
		db: sqlDB,
	}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) DummyMigration(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS users
(
    id          uuid PRIMARY KEY   NOT NULL DEFAULT gen_random_uuid(),
    name        VARCHAR            NOT NULL,
    description VARCHAR,
    created_at   timestamp          NOT NULL default now(),
    updated_at   timestamp          NOT NULL default now(),
    deleted  bool DEFAULT FALSE NOT NULL
);`

	if _, err := s.db.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("create table: %v", err)
	}

	s.lg.Info("Migration is succeed...")

	return nil
}
