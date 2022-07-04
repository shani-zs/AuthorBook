package authorhttp

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/service"
	"strconv"
)

type authorHandler struct {
	authorService service.AuthorService
}

func New(a service.AuthorService) authorHandler {
	return authorHandler{a}
}

func (h authorHandler) PostAuthor(w http.ResponseWriter, req *http.Request) {
	var author entities.Author

	body, _ := io.ReadAll(req.Body)
	err := json.Unmarshal(body, &author)

	_, err = h.authorService.PostAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("successful posted!"))
}

func (h authorHandler) PutAuthor(w http.ResponseWriter, req *http.Request) {
	var author entities.Author

	body, _ := io.ReadAll(req.Body)
	_ = json.Unmarshal(body, &author)

	_, err := h.authorService.PutAuthor(author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "sucessful posted!")
}

func (h authorHandler) DeleteAuthor(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	_, err := h.authorService.DeleteAuthor(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("could not delete"))

		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("successfully deleted!"))
}
