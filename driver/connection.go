package driver

import (
	"database/sql"
	"log"
)

func Connection() *sql.DB {
	DB, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/authorbook")
	if err != nil {
		log.Fatal("failed to connect with database:\n", err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal("failed to ping", pingErr)
	}

	return DB
}
