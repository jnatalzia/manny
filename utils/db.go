package utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DBConn *sql.DB

func ConnectToDB() {
	var err error
	connStr := "host=localhost user=local dbname=manipulation password=postgres sslmode=disable" // sslmode=verify-full"
	DBConn, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
