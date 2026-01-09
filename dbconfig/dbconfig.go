package dbconfig

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB(databaseURL string) *sql.DB {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Failed to connect to DB", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Database connection failed", err)
	}
	fmt.Println("Database connected successfully")
	return db
}
