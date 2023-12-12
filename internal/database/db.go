package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func Connect(path string) (*sql.DB, error) {
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
		Dir: "internal/database/mirgrations",
	}
	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("migrateUp: %w", err)
	}
	return nil
}
