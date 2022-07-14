package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"projects/GoLang-Interns-2022/authorbook/driver"
	"projects/GoLang-Interns-2022/authorbook/http/authorhttp"
	"projects/GoLang-Interns-2022/authorbook/http/bookhttp"
	"projects/GoLang-Interns-2022/authorbook/service/authorservice"
	"projects/GoLang-Interns-2022/authorbook/service/bookservice"
	"projects/GoLang-Interns-2022/authorbook/store/author"
	"projects/GoLang-Interns-2022/authorbook/store/book"
)

func main() {
	r := mux.NewRouter()

	DB := driver.Connection()
	defer DB.Close()

	authorStore := author.New(DB)
	authorService := authorservice.New(authorStore)
	authorHandler := authorhttp.New(authorService)
	// author endpoints
	r.HandleFunc("/author", authorHandler.Post).Methods("POST")
	r.HandleFunc("/author/{id}", authorHandler.Put).Methods("PUT")
	r.HandleFunc("/author/{id}", authorHandler.Delete).Methods("DELETE")

	bookStore := book.New(DB)
	bookService := bookservice.New(bookStore, authorStore)
	bookHandler := bookhttp.New(bookService)
	// book  endpoints
	r.HandleFunc("/book", bookHandler.GetAllBook).Methods("GET")
	r.HandleFunc("/book/{id}", bookHandler.GetBookByID).Methods("GET")
	r.HandleFunc("/book", bookHandler.Post).Methods("POST")
	r.HandleFunc("/book/{id}", bookHandler.Put).Methods("PUT")
	r.HandleFunc("/book/{id}", bookHandler.Delete).Methods("DELETE")

	fmt.Println("server started...")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Print(err)
	}
}
