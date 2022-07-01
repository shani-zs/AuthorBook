package authorhttp

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/service"
	"projects/GoLang-Interns-2022/authorbook/service/authorservice"
	"strconv"
)

type authorHandler struct {
	authorService service.AuthorService
}

func New(a authorservice.AuthorService) authorHandler {
	return authorHandler{a}
}

func (h authorHandler) PostAuthor(w http.ResponseWriter, req *http.Request) {
	var author entities.Author

	body, _ := io.ReadAll(req.Body)
	err := json.Unmarshal(body, &author)
	if err != nil {
		log.Printf("falied %v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

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
	err := json.Unmarshal(body, &author)
	if err != nil {
		log.Printf("falied %v\n", err)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	id, err := h.authorService.PutAuthor(author)
	if err != nil || id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "sucessful posted!")
}

func (h authorHandler) DeleteAuthor(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	i, err := strconv.Atoi(params["id"])
	if err != nil {
		_, _ = w.Write([]byte("invalid parameter id"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	res, err := h.authorService.DeleteAuthor(i)
	if err != nil || res == 0 {
		_, _ = w.Write([]byte("could not delete"))
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.Write([]byte("successfully deleted!"))
	w.WriteHeader(http.StatusNoContent)
}
