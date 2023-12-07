package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func Connect(path string) (*sql.DB, error) {
	db, err := sql.Open("postgres", path)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	if err = migrateUp(db); err != nil {
		return nil, err
	}

	return db, nil
}

func migrateUp(db *sql.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "internal/database/mirgrations",
	}
	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}
	return nil
}
