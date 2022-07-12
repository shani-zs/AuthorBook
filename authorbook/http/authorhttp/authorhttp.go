package authorhttp

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

type AuthorHandler struct {
	authorService service.AuthorService
}

// New : factory function used for injection
func New(a service.AuthorService) AuthorHandler {
	return AuthorHandler{a}
}

// Post : handles the request of posting an author
func (h AuthorHandler) Post(w http.ResponseWriter, req *http.Request) {
	var author entities.Author

	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := req.Context()

	a, err := h.authorService.Post(ctx, author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	au, err := json.Marshal(a)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(au)
}

// Put : handles the request of updating an author
func (h AuthorHandler) Put(w http.ResponseWriter, req *http.Request) {
	var author entities.Author

	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &author)
	if err != nil {
		log.Print("3")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := req.Context()

	author1, err := h.authorService.Put(ctx, author, id)
	if err != nil {
		log.Print("2")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("dose not exist"))

		return
	}

	data, err := json.Marshal(author1)
	if err != nil {
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(data)
}

// Delete : handles the request of deleting an author
func (h AuthorHandler) Delete(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := req.Context()

	_, err = h.authorService.Delete(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("could not delete"))

		return
	}

	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte("successfully deleted!"))
}
