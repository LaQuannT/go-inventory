package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func buildConnStr() string {
	env := struct {
		Username string
		Password string
		Server   string
		Port     string
		Database string
		Ssl      string
	}{
		os.Getenv("PG_USERNAME"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_SERVER"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_DATABASE"),
		os.Getenv("PG_SSLMODE"),
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", env.Username, env.Password, env.Server, env.Port, env.Database, env.Ssl)
}

func Connect() (*sql.DB, error) {
	path := buildConnStr()

	db, err := sql.Open("postgres", path)
	if err != nil {
		return nil, fmt.Errorf("database/Connect: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("database/Connect: %w", err)
	}

	if err = migrateUp(db); err != nil {
		return nil, fmt.Errorf("database/Connect: %w", err)
	}

	return db, nil
}

func migrateUp(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "database/mirgrations",
	}
	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("migrateUp: %w", err)
	}
	return nil
}
