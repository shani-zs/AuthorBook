package bookhttp

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"reflect"
	"testing"
)

func TestGetAllBook(t *testing.T) {
	Testcases := []struct {
		desc          string
		title         string
		includeAuthor string

		expectedBooks []entities.Book
	}{
		{"getting all books", "", "", []entities.Book{{1,
			1, "book one", "scholastic", "20/06/2018", entities.Author{}},
			{2, 1, "book two", "penguin", "20/08/2018", entities.Author{}}},
		},
		//{"getting book with author and particular title", "book+two", "true", []entities.Book{
		//	{2, 1, "book two", "penguin", "20/08/2018", entities.Author{1, "shani",
		//		"kumar", "30/04/2001", "sk"}}}},
	}

	for _, tc := range Testcases {
		req := httptest.NewRequest("GET", "localhost:8000/book?"+"title="+tc.title+"&"+"includeAuthor="+tc.includeAuthor, nil)
		w := httptest.NewRecorder()
		h := New(mockService{})

		h.GetAllBook(w, req)

		result := w.Result()

		body, err := io.ReadAll(result.Body)
		if err != nil {
			log.Print(err)
		}

		var books []entities.Book

		_ = json.Unmarshal(body, &books)

		if reflect.DeepEqual(books, tc.expectedBooks) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

func TestGetBookByID(t *testing.T) {
	Testcases := []struct {
		desc     string
		targetID string

		expectedBook       entities.Book
		expectedStatusCode int
	}{
		{"fetching book by id",
			"1", entities.Book{1, 1, "book two", "penguin", "20/08/2018",
				entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}}, http.StatusOK},

		{"invalid id", "-1", entities.Book{}, http.StatusBadRequest},
	}

	for _, tc := range Testcases {
		req := httptest.NewRequest("GET", "localhost:8000/book/{id}"+tc.targetID, nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": tc.targetID})
		h := New(mockService{})

		h.GetBookByID(w, req)

		result := w.Result()
		body, err := io.ReadAll(result.Body)

		if err != nil {
			log.Print(err)
		}

		var books entities.Book

		_ = json.Unmarshal(body, &books)

		if reflect.DeepEqual(books, tc.expectedBook) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

func TestPostBook(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Book

		expectedStatusCode int
	}{
		{"valid case", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}}, http.StatusCreated},
		{"already existing book", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			http.StatusBadRequest},
		{"invalid bookID", entities.Book{-4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			http.StatusBadRequest},
		{"invalid author's DOB", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/00/2001", "sk"}},
			http.StatusBadRequest},
		{"invalid title", entities.Book{4, 1, "", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}}, http.StatusBadRequest},
		{"invalid publication", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			http.StatusBadRequest},
		{"invalid published date", entities.Book{4, 1, "deciding decade", "penguin",
			"00/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}}, http.StatusBadRequest},
	}
	for _, tc := range testcases {
		b, err := json.Marshal(tc.body)
		if err != nil {
			log.Printf("failed : %v", err)
		}

		req := httptest.NewRequest("POST", "localhost:8000/book", bytes.NewBuffer(b))
		w := httptest.NewRecorder()
		h := New(mockService{})

		h.PostBook(w, req)

		result := w.Result()
		if reflect.DeepEqual(tc.expectedStatusCode, result.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

func TestPutBook(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Book

		expectedStatusCode int
	}{
		{"creating a book", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}}, http.StatusCreated},
		{"already existing book", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			http.StatusCreated},
		{"invalid bookID", entities.Book{-4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			http.StatusBadRequest},
		{"invalid author's DOB", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/00/2001", "sk"}},
			http.StatusBadRequest},
		{"invalid title", entities.Book{4, 1, "", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}}, http.StatusBadRequest},
		{"invalid publication", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			http.StatusBadRequest},
		{"invalid published date", entities.Book{4, 1, "deciding decade", "penguin",
			"00/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}}, http.StatusBadRequest},
	}
	for _, tc := range testcases {
		b, err := json.Marshal(tc.body)
		if err != nil {
			log.Printf("failed : %v", err)
		}

		req := httptest.NewRequest("PUT", "localhost:8000/book", bytes.NewBuffer(b))
		w := httptest.NewRecorder()
		h := New(mockService{})

		h.PutBook(w, req)

		result := w.Result()
		if reflect.DeepEqual(tc.expectedStatusCode, result.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

func TestDeleteBook(t *testing.T) {
	testcases := []struct {
		desc     string
		targetID string

		expectedStatus int
	}{
		{"valid id", "1", http.StatusNoContent},
		{"invalid id", "-1", http.StatusBadRequest},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest("PUT", "localhost:8000/book/{id}"+tc.targetID, nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.targetID})
		w := httptest.NewRecorder()
		h := New(mockService{})

		h.DeleteBook(w, req)

		result := w.Result()
		if reflect.DeepEqual(tc.expectedStatus, result.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}
