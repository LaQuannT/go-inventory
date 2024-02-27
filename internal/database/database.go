package database

import (
	"context"
	"fmt"
	"os"
	"time"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

func Connect() (*pgx.Conn, error) {
	ctx := context.Background()

	username := os.Getenv("PG_USERNAME")
	pwd := os.Getenv("PG_PASSWORD")
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	dbName := os.Getenv("PG_DATABASE")
	mode := os.Getenv("PG_SSLMODE")

	path := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		username, pwd, host, port, dbName, mode,
	)

	conn, err := pgx.Connect(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err = conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to verify connection status: %w", err)
	}

	if err := migrateUp(path); err != nil {
		return nil, fmt.Errorf("unable to complete database migrations: %w", err)
	}

	return conn, nil
}

func migrateUp(connStr string) error {
	m, err := migrate.New("file://migrations", connStr)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		return err
	}

	return nil
}
