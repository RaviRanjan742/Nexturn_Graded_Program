
package db

import (
    "database/sql"
    _ "modernc.org/sqlite"
)

func InitDB(filepath string) (*sql.DB, error) {
    db, err := sql.Open("sqlite", filepath)
    if err != nil {
        return nil, err
    }

    
    sqlStmt := `
    CREATE TABLE IF NOT EXISTS blogs (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        author TEXT NOT NULL,
        timestamp DATETIME
    );`

    _, err = db.Exec(sqlStmt)
    if err != nil {
        return nil, err
    }

    return db, nil
}