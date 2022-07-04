package main

import (
	"github.com/gorilla/mux"
	"projects/GoLang-Interns-2022/authorbook/http/bookhttp"
	"projects/GoLang-Interns-2022/authorbook/service/bookservice"
	"projects/GoLang-Interns-2022/authorbook/store/book"

	//"github.com/gorilla/mux"
	"log"
	"net/http"
	"projects/GoLang-Interns-2022/authorbook/driver"
	"projects/GoLang-Interns-2022/authorbook/http/authorhttp"
	"projects/GoLang-Interns-2022/authorbook/service/authorservice"
	"projects/GoLang-Interns-2022/authorbook/store/author"
)

func main() {
	r := mux.NewRouter()

	DB := driver.Connection()
	defer DB.Close()

	authorStore := author.New(DB)
	authorService := authorservice.New(authorStore)
	authorHandler := authorhttp.New(authorService)
	//author endpoints
	r.HandleFunc("/author", authorHandler.PostAuthor).Methods("POST")
	r.HandleFunc("/author/{id}", authorHandler.PutAuthor).Methods("PUT")
	r.HandleFunc("/author/{id}", authorHandler.DeleteAuthor).Methods("DELETE")

	bookStore := book.New(DB)
	bookService := bookservice.New(bookStore)
	bookHandler := bookhttp.New(bookService)
	//book  endpoints
	r.HandleFunc("/book", bookHandler.GetAllBook).Methods("GET")
	r.HandleFunc("/book/{id}", bookHandler.GetBookByID).Methods("GET")
	r.HandleFunc("/book", bookHandler.PostBook).Methods("POST")
	r.HandleFunc("/book/{id}", bookHandler.PutBook).Methods("PUT")
	r.HandleFunc("/book/{id}", bookHandler.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
