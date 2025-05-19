package storage

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func newSqliteDB(options map[string]any) (*sql.DB, error) {
	path, ok := options["path"].(string)
	if !ok {
		return nil, fmt.Errorf("file path not specified in storage options")
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	// Open the database
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	// Configure connection pool
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
