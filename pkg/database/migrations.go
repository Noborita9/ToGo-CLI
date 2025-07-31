package database

import (
	"database/sql"
	"fmt"
)

func CreateTables(db * sql.DB) error {
	createTablesTask := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        description TEXT NOT NULL,
        priority TEXT DEFAULT 'normal',
        completed BOOLEAN DEFAULT FALSE,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(createTablesTask)
	if err != nil {
		return fmt.Errorf("Failed to create tasks table: %w", err)
	}
	return nil
}
