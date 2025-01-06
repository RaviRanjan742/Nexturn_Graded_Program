package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func ConnectDatabase() *sql.DB {
	db, err := sql.Open("sqlite", "./inventory.db") 
	if err != nil {
		log.Fatal(err)
	}

	
	createTable := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		price REAL NOT NULL,
		stock INTEGER NOT NULL,
		category_id INTEGER
	);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
