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

// go test  //  runs test for current package
// go test ./... //for current and all the subpackages
// go test ./... -v // Gives you detailed output of the TestFunctions
// go test ./... -v -coverprofile=coverage.out  // gives you a general coverage and generates a report in coverage.out
// go tool cover -func coverage.out // this will show you function wise coverage report
// go tool cover -html coverage.out // this will show you line wise coverage report for all files
