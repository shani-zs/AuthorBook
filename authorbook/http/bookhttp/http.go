package bookhttp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/service"
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
	ctx := req.Context()

	books, err := h.bookH.GetAllBook(ctx, title, includeAuthor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	data, err := json.Marshal(books)
	if err != nil {
		_, _ = w.Write([]byte("could not encode"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.Header().Set("content-type", "json/application")
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

	ctx := req.Context()

	book, err := h.bookH.GetBookByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	data, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("content-type", "json/application")
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

	ctx := req.Context()

	book1, err := h.bookH.Post(ctx, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	data, err := json.Marshal(book1)
	if err != nil {
		log.Print(err)

		_, _ = w.Write([]byte("could not read!"))

		return
	}

	w.Header().Set("content-type", "json/application")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(data)
}

// Put : handle the request of updating a book
func (h BookHandler) Put(w http.ResponseWriter, req *http.Request) {
	var (
		body []byte
		data []byte
		book entities.Book
		id   int
	)

	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	err = json.Unmarshal(body, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("could not decode "))

		return
	}

	params := mux.Vars(req)

	id, err = strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	ctx := req.Context()

	book, err = h.bookH.Put(ctx, &book, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	data, err = json.Marshal(book)
	if err != nil {
		return
	}

	w.Header().Set("content-type", "json/application")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(data)
}

// Delete : handles the request of removing a book
func (h BookHandler) Delete(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	ctx := req.Context()

	err = h.bookH.Delete(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(err.Error()))

		return
	}

	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte("successfully deleted"))

	return
}
