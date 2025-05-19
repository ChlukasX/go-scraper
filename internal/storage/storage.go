package storage

import "fmt"

type DataStore interface {
	Close() error
}

func New(strgType string, options map[string]any) (DataStore, error) {
	switch strgType {
	case "sqlite", "sqlite3":
		return newSqliteDB(options)
	default:
		return nil, fmt.Errorf("Unsuported data store type: %s", strgType)
	}
}
