package authorhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"testing"

	"github.com/gorilla/mux"
)

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author

		expected int
	}{
		{desc: "invalid DOB", body: entities.Author{
			AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/00/1990", PenName: "Dark horse"}, expected: http.StatusBadRequest},
		{"nil body", entities.Author{}, http.StatusBadRequest},
		{desc: "exiting author", body: entities.Author{
			AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"}, expected: http.StatusBadRequest},
		{desc: "invalid firstname", body: entities.Author{
			AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"}, expected: http.StatusBadRequest},
		{desc: "valid author", body: entities.Author{
			AuthorID: 4, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"}, expected: http.StatusCreated},
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
		{desc: "valid author", body: entities.Author{
			AuthorID: 4, FirstName: "Shani", LastName: "mrinal", DOB: "20/05/1970", PenName: "Dh"},
			expected: http.StatusCreated},
		{desc: "exiting author", body: entities.Author{
			AuthorID: 3, FirstName: "shani", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"},
			expected: http.StatusBadRequest},
		{desc: "invalid firstname", body: entities.Author{
			AuthorID: 3, LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"},
			expected: http.StatusBadRequest},
		{desc: "invalid DOB", body: entities.Author{
			AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/00/1990", PenName: "Dark horse"},
			expected: http.StatusBadRequest},
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
