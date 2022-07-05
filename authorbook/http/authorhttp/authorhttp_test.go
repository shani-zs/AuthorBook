package authorhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"testing"
)

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author

		expected int
	}{
		{"valid author", entities.Author{
			4, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, http.StatusCreated},
		{"exiting author", entities.Author{
			3, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, http.StatusBadRequest},
		{"invalid firstname", entities.Author{
			3, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, http.StatusBadRequest},
		{"invalid DOB", entities.Author{
			3, "nilotpal", "mrinal", "20/00/1990", "Dark horse"}, http.StatusBadRequest},
		{"nil body", entities.Author{}, http.StatusBadRequest},
	}

	for _, tc := range testcases {
		data, err := json.Marshal(tc.body)
		if err != nil {
			log.Print(err)
		}

		req := httptest.NewRequest("POST", "localhost:8000/author", bytes.NewReader(data))
		w := httptest.NewRecorder()
		h := New(mockService{})

		h.PostAuthor(w, req)

		res := w.Result()
		if tc.expected != res.StatusCode {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author

		expected int
	}{
		{"valid author", entities.Author{
			4, "Shani", "mrinal", "20/05/1970", "Dh"}, http.StatusCreated},
		{"exiting author", entities.Author{
			3, "shani", "mrinal", "20/05/1990", "Dark horse"}, http.StatusBadRequest},
		{"invalid firstname", entities.Author{
			3, "", "mrinal", "20/05/1990", "Dark horse"}, http.StatusBadRequest},
		{"invalid DOB", entities.Author{
			3, "nilotpal", "mrinal", "20/00/1990", "Dark horse"}, http.StatusBadRequest},
		{"nil body", entities.Author{}, http.StatusBadRequest},
	}

	for _, tc := range testcases {
		data, err := json.Marshal(tc.body)
		if err != nil {
			log.Print(err)
		}

		req := httptest.NewRequest("PUT", "localhost:8000/author", bytes.NewReader(data))
		w := httptest.NewRecorder()
		h := New(mockService{})

		h.PutAuthor(w, req)

		res := w.Result()
		if tc.expected != res.StatusCode {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc   string
		target string

		expected int
	}{
		{"valid authorId", "4", http.StatusNoContent},
		{"invalid authorId", "-3", http.StatusBadRequest},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest("DELETE", "localhost:8000/author/{id}"+tc.target, nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.target})
		w := httptest.NewRecorder()
		h := New(mockService{})

		h.DeleteAuthor(w, req)

		res := w.Result()
		if tc.expected != res.StatusCode {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

type mockService struct{}

func (h mockService) PutAuthor(author2 entities.Author) (entities.Author, error) {
	if author2.AuthorID == 4 {
		return entities.Author{}, nil
	}

	return entities.Author{}, errors.New("invalid constraints")
}

func (h mockService) DeleteAuthor(id int) (int, error) {
	if id == 4 {
		return id, nil
	}

	return -1, errors.New("invalid")
}

func (h mockService) PostAuthor(author2 entities.Author) (entities.Author, error) {
	if author2.AuthorID == 4 {
		return entities.Author{}, nil
	}

	return entities.Author{}, errors.New("invalid constraints")
}
