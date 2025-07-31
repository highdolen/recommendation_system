package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connStr := "host=localhost port=5433 user=fedya password=secret dbname=recdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных", err)
	}

	return db
}
