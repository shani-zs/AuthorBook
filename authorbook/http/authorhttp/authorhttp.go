package authorhttp

import (
	"encoding/json"
	"fmt"
	"io"
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

func (h AuthorHandler) PostAuthor(w http.ResponseWriter, req *http.Request) {
	var author entities.Author

	body, _ := io.ReadAll(req.Body)
	_ = json.Unmarshal(body, &author)

	_, err := h.authorService.PostAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("successful posted!"))
}

func (h AuthorHandler) PutAuthor(w http.ResponseWriter, req *http.Request) {
	var author entities.Author

	body, _ := io.ReadAll(req.Body)
	_ = json.Unmarshal(body, &author)

	_, err := h.authorService.PutAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "successful posted!")
}

func (h AuthorHandler) DeleteAuthor(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	_, err := h.authorService.DeleteAuthor(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("could not delete"))

		return
	}

	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte("successfully deleted!"))
}
