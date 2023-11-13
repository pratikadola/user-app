package migration

import "database/sql"

func CreateTables(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT NOT NULL)")
	if err != nil {
		panic(err)
	}
}
