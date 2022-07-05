package bookhttp

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBook(t *testing.T) {
	Testcases := []struct {
		desc          string
		title         string
		includeAuthor string

		expectedBooks []entities.Book
	}{
		{desc: "getting all books", title: "", includeAuthor: "", expectedBooks: []entities.Book{{BookID: 1,
			AuthorID: 1, Title: "book one", Publication: "scholastic", PublishedDate: "20/06/2018",
			Author: entities.Author{}},
			{BookID: 2, AuthorID: 1, Title: "book two", Publication: "penguin", PublishedDate: "20/08/2018", Author: entities.Author{}}},
		},
		{desc: "getting book with author and particular title", title: "book+two", includeAuthor: "true",
			expectedBooks: []entities.Book{{BookID: 2, AuthorID: 1, Title: "book two", Publication: "penguin",
				PublishedDate: "20/08/2018", Author: entities.Author{AuthorID: 1, FirstName: "shani", LastName: "kumar",
					DOB: "30/04/2001", PenName: "sk"}}}},
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

		if !assert.Equal(t, tc.expectedBooks, books) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

func TestGetBookByID(t *testing.T) {
	Testcases := []struct {
		desc               string
		targetID           string
		expectedBook       entities.Book
		expectedStatusCode int
	}{
		{desc: "fetching book by id", targetID: "1", expectedBook: entities.Book{BookID: 1, AuthorID: 1, Title: "book two",
			Publication: "penguin", PublishedDate: "20/08/2018", Author: entities.Author{AuthorID: 1, FirstName: "shani",
				LastName: "kumar", DOB: "30/04/2001", PenName: "sk"}}, expectedStatusCode: http.StatusOK},
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

		if !reflect.DeepEqual(books, tc.expectedBook) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

func TestPostBook(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Book

		expectedBook entities.Book
	}{
		{desc: "valid case", body: entities.Book{BookID: 4, AuthorID: 1, Title: "deciding decade", Publication: "penguin",
			PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1, FirstName: "shani", LastName: "kumar",
				DOB: "30/04/2001", PenName: "sk"}},
			expectedBook: entities.Book{BookID: 4, AuthorID: 1, Title: "deciding decade", Publication: "penguin",
				PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1, FirstName: "shani", LastName: "kumar",
					DOB: "30/04/2001", PenName: "sk"}}},

		{desc: "already existing book", body: entities.Book{BookID: 1, AuthorID: 1, Title: "deciding decade",
			Publication: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1, FirstName: "shani",
				LastName: "kumar", DOB: "30/04/2001", PenName: "sk"}}},
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

		var books entities.Book

		body, err := io.ReadAll(result.Body)
		if err != nil {
			log.Print(err)
		}

		_ = json.Unmarshal(body, &books)

		if !reflect.DeepEqual(tc.expectedBook, books) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

func TestPutBook(t *testing.T) {
	testcases := []struct {
		desc         string
		body         entities.Book
		targetID     string
		expectedBook entities.Book
	}{
		{desc: "creating a book", body: entities.Book{BookID: 4, AuthorID: 1, Title: "deciding decade",
			Publication: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1, FirstName: "shani",
				LastName: "kumar", DOB: "30/04/2001", PenName: "sk"}}, targetID: "1",
			expectedBook: entities.Book{BookID: 4, AuthorID: 1, Title: "deciding decade", Publication: "penguin",
				PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1, FirstName: "shani", LastName: "kumar",
					DOB: "30/04/2001", PenName: "sk"}}},
		{desc: "already existing book", body: entities.Book{BookID: 4, AuthorID: 1, Title: "decade",
			Publication: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1, FirstName: "shani",
				LastName: "kumar", DOB: "30/04/2001", PenName: "sk"}},
			targetID: "2", expectedBook: entities.Book{BookID: 4, AuthorID: 1, Title: "decade", Publication: "penguin",
				PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1, FirstName: "shani", LastName: "kumar",
					DOB: "30/04/2001", PenName: "sk"}}},
	}
	for _, tc := range testcases {
		b, err := json.Marshal(tc.body)
		if err != nil {
			log.Printf("failed : %v", err)
		}

		req := httptest.NewRequest("PUT", "localhost:8000/book/{id}"+tc.targetID, bytes.NewBuffer(b))
		req = mux.SetURLVars(req, map[string]string{"id": tc.targetID})
		w := httptest.NewRecorder()
		h := New(mockService{})

		h.PutBook(w, req)

		result := w.Result()

		var book entities.Book

		data, err := io.ReadAll(result.Body)
		if err != nil {
			log.Print(err)
		}

		_ = json.Unmarshal(data, &book)

		if !reflect.DeepEqual(tc.expectedBook, book) {
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
		if !reflect.DeepEqual(tc.expectedStatus, result.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}
