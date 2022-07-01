package authorhttp

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"projects/GoLang-Interns-2022/authorbook/driver"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/service/authorservice"
	"projects/GoLang-Interns-2022/authorbook/store/author"
	"testing"
)

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author

		expected int
	}{
		{"valid author", entities.Author{
			4, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, http.StatusBadRequest},
		{"exiting author", entities.Author{
			3, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, http.StatusBadRequest},
		{"invalid firstname", entities.Author{
			3, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, http.StatusBadRequest},
		{"invalid DOB", entities.Author{
			3, "nilotpal", "mrinal", "20/00/1990", "Dark horse"}, http.StatusBadRequest},
	}

	for _, tc := range testcases {
		data, err := json.Marshal(tc.body)
		if err != nil {
			log.Print(err)
		}

		req := httptest.NewRequest("POST", "localhost:8000/author", bytes.NewReader(data))
		w := httptest.NewRecorder()

		DB := driver.Connection()
		authorStore := author.New(DB)
		authorService := authorservice.New(authorStore)
		h := New(authorService)

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
			4, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, http.StatusCreated},
		{"exiting author", entities.Author{
			3, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, http.StatusBadRequest},
		{"invalid firstname", entities.Author{
			3, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, http.StatusBadRequest},
		{"invalid DOB", entities.Author{
			3, "nilotpal", "mrinal", "20/00/1990", "Dark horse"}, http.StatusBadRequest},
	}

	for _, tc := range testcases {
		data, err := json.Marshal(tc.body)
		if err != nil {
			log.Print(err)
		}

		req := httptest.NewRequest("POST", "localhost:8000/author", bytes.NewReader(data))
		w := httptest.NewRecorder()

		DB := driver.Connection()
		authorStore := author.New(DB)
		authorService := authorservice.New(authorStore)
		h := New(authorService)

		h.PostAuthor(w, req)

		res := w.Result()
		if tc.expected != res.StatusCode {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		//input
		desc   string
		target string
		//output
		expected int
	}{
		{"valid authorId", "4", http.StatusNoContent},
		{"invalid authorId", "-3", http.StatusBadRequest},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest("DELETE", "localhost:8000/author/{id}"+tc.target, nil)

		w := httptest.NewRecorder()

		DB := driver.Connection()
		authorStore := author.New(DB)
		authorService := authorservice.New(authorStore)
		h := New(authorService)

		h.DeleteAuthor(w, req)

		res := w.Result()

		assert.Equal(t, tc.expected, res.StatusCode)
	}
}
