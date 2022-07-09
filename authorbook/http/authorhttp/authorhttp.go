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

func New(a service.AuthorService) AuthorHandler {
	return AuthorHandler{a}
}

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

	a, err := h.authorService.Post(author)
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

func (h AuthorHandler) Put(w http.ResponseWriter, req *http.Request) {
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

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	author1, err := h.authorService.Put(author, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

func (h AuthorHandler) Delete(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.authorService.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("could not delete"))

		return
	}

	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte("successfully deleted!"))
}
