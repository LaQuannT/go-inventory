package database

import (
	"context"
	"fmt"
	"os"
	"time"

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

	return conn, nil
}
