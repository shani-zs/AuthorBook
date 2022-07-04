package bookhttp

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/service"
	"strconv"
)

type bookHandler struct {
	bookH service.BookService
}

func New(bookS service.BookService) bookHandler {
	return bookHandler{bookS}
}

func (h bookHandler) GetAllBook(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Query().Get("title")
	includeAuthor := req.URL.Query().Get("includeAuthor")
	books := h.bookH.GetAllBook(title, includeAuthor)

	data, err := json.Marshal(books)
	if err != nil {
		w.Write([]byte("could not encode"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h bookHandler) GetBookByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.Write([]byte("invalid id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, err := h.bookH.GetBookByID(id)
	if err != nil {
		w.Write([]byte("book does not exist"))
		return
	}

	data, err := json.Marshal(book)
	if err != nil {
		w.Write([]byte("could not encode"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h bookHandler) PostBook(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.Write([]byte("could not read!"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	var book entities.Book

	err = json.Unmarshal(body, &book)
	if err != nil {
		w.Write([]byte("could not decode "))
		return
	}

	book1, err := h.bookH.PostBook(book)
	if err != nil {
		log.Print(err)
		w.Write([]byte("could not read!"))
		return
	}

	data, err := json.Marshal(book1)
	if err != nil {
		log.Print(err)
		w.Write([]byte("could not read!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h bookHandler) PutBook(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.Write([]byte("could not read!"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	var book entities.Book

	err = json.Unmarshal(body, &book)
	if err != nil {
		w.Write([]byte("could not decode "))
		return
	}

	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	book, err = h.bookH.PutBook(book, id)
	if err != nil {
		log.Print(err)
		return
	}

	data, err := json.Marshal(book)
	if err != nil {
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (h bookHandler) DeleteBook(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Print(err)
		return
	}

	_, err = h.bookH.DeleteBook(id)
	if err != nil {
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte("successfully deleted"))
}
