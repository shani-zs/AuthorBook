package main

import (
	"github.com/gorilla/mux"
	"http/AuthorBook"
	"log"
	"net/http"
	"projects/GoLang-Interns-2022/authorbook"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/book", authorbook.GetAllBook).Methods("GET")
	r.HandleFunc("/book/{id}", authorbook.GetBookByID).Methods("GET")
	r.HandleFunc("/book", AuthorBook.PostBook).Methods(http.MethodPost)
	r.HandleFunc("/author", AuthorBook.PostAuthor).Methods(http.MethodPost)
	r.HandleFunc("/book/{id}", AuthorBook.PutBook).Methods("PUT")
	r.HandleFunc("/author/{id}", AuthorBook.PutAuthor).Methods("PUT")
	r.HandleFunc("/book/{id}", AuthorBook.DeleteBook).Methods("DELETE")
	r.HandleFunc("/author/{id}", AuthorBook.DeleteAuthor).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
