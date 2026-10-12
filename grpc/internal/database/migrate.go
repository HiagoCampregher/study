package database

import "database/sql"

// Migrate creates the tables required by the examples. It is intentionally
// small so the study can run immediately after a fresh clone.
func Migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS courses (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			category_id TEXT NOT NULL,
			FOREIGN KEY (category_id) REFERENCES categories(id)
		);
	`)
	return err
}
