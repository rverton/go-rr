package database

import (
	"database/sql"
	_ "embed"
)

//go:embed schema.sql
var Schema string

// ConfigureSqlite sets some sensible defaults for sqlite
func ConfigureSqlite(db *sql.DB) error {
	pragmas := []string{
		"busy_timeout = 5000",
		"journal_mode = WAL",
		"synchronous = NORMAL",
		"cache_size = 1000000000", // 1GB
		"foreign_keys = true",
		"temp_store = memory",
		"mmap_size = 3000000000",
	}

	for _, pragma := range pragmas {
		_, err := db.Exec("PRAGMA " + pragma)
		if err != nil {
			return err
		}
	}
	return nil
}
