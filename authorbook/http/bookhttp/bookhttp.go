package bookhttp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/service"
	"strconv"

	"github.com/gorilla/mux"
)

type BookHandler struct {
	bookH service.BookService
}

// New : factory function
func New(bookS service.BookService) BookHandler {
	return BookHandler{bookS}
}

// GetAllBook : handles the request of getting all books
func (h BookHandler) GetAllBook(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Query().Get("title")
	includeAuthor := req.URL.Query().Get("includeAuthor")

	books, err := h.bookH.GetAllBook(title, includeAuthor)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	data, err := json.Marshal(books)
	if err != nil {
		_, _ = w.Write([]byte("could not encode"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// GetBookByID : handles the request of getting a book
func (h BookHandler) GetBookByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])
	if err != nil || id < 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid id"))

		return
	}

	book, err := h.bookH.GetBookByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("book does not exist"))

		return
	}

	data, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

// Post : handles the request of posting a book
func (h BookHandler) Post(w http.ResponseWriter, req *http.Request) {
	var book entities.Book

	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("could not read!"))

		return
	}

	err = json.Unmarshal(body, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("could not decode "))

		return
	}

	book1, err := h.bookH.Post(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("author does exist!"))

		return
	}

	data, err := json.Marshal(book1)
	if err != nil {
		log.Print(err)
		_, _ = w.Write([]byte("could not read!"))

		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(data)
}

// Put : handle the request of updating a book
func (h BookHandler) Put(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("could not read!"))

		return
	}

	var book entities.Book

	err = json.Unmarshal(body, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("could not decode "))

		return
	}

	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	book, err = h.bookH.Put(&book, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Print(err)
		return
	}

	data, err := json.Marshal(book)
	if err != nil {
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(data)
}

// Delete : handles the request of removing a book
func (h BookHandler) Delete(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])
	if err != nil || id < 0 {
		w.WriteHeader(http.StatusBadRequest)
		log.Print(err)
		return
	}

	_, err = h.bookH.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Print(err)

		return
	}

	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte("successfully deleted"))
}
