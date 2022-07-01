package driver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connection() *sql.DB {
	DB, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/AuthorBook")
	if err != nil {
		log.Fatal("failed to connect with database:\n", err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal("failed to ping", pingErr)
	}

	return DB
}
