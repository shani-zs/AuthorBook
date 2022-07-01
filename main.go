package main

import (
	"github.com/gorilla/mux"
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

	authorStore := author.New(DB)
	authorService := authorservice.New(authorStore)
	authorHandler := authorhttp.New(authorService)

	r.HandleFunc("/author", authorHandler.PostAuthor).Methods("POST")
	r.HandleFunc("/author/{id}", authorHandler.PutAuthor).Methods("PUT")
	r.HandleFunc("/author/{id}", authorHandler.DeleteAuthor).Methods("DELETE")
	//r.HandleFunc()

	log.Fatal(http.ListenAndServe(":8000", r))
}
