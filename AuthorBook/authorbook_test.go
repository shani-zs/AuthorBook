package AuthorBook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllBook(t *testing.T) {

	Testcases := []struct {
		desc        string
		methodInput string
		target      string
		//body        io.Reader
		expected []Book
	}{
		{"test for fetching books", "GET", "http://localhost:8000/book", []Book{{"1",
			1, "book one", "penguin", "20/06/2018", &Author{1, "shani",
				"kumar", "30/04/2001", "sk"}},
			{"2", 1, "book two", "penguin", "20/08/2018", &Author{1, "shani",
				"kumar", "30/04/2001", "sk"}}},
		},
	}

	for _, tc := range Testcases {
		req := httptest.NewRequest(tc.methodInput, tc.target, nil)
		w := httptest.NewRecorder()
		GetAllBook(w, req)

		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)

		var book []Book

		err := json.Unmarshal(body, &book) //Unmarshal([]byte,v any)
		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, tc.expected, book)

	}

}

func TestGetBookById(t *testing.T) {
	Testcases := []struct {
		desc               string
		methodInput        string
		bookId             string
		expected           Book
		expectedStatusCode int
	}{
		{"fetching book by id", "GET", "2",
			Book{"2", 1, "book two", "penguin", "20/08/2018",
				&Author{1, "shani", "kumar", "30/04/2001", "sk"}}, http.StatusOK},
		{"invalid id", "GET", "-2",
			Book{"-2", 1, "book two", "penguin", "20/08/2018",
				&Author{1, "shani", "kumar", "30/04/2001", "sk"}}, http.StatusBadRequest},
		{"invalid method", "POST", "2",
			Book{}, http.StatusBadRequest},
	}

	for _, tc := range Testcases {

		req := httptest.NewRequest(tc.methodInput, "http://localhost:8000/book/{id}"+tc.bookId, nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": tc.bookId})
		GetBookById(w, req)

		resp := w.Result()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var book Book
		_ = json.Unmarshal(body, &book) //Unmarshal([]byte,any type)

		if resp.StatusCode != tc.expectedStatusCode {
			t.Errorf("failed for %v\n", tc.desc)
		}
		//assert.Equal(t, tc.expected, book)
	}
}

func TestPostBook(t *testing.T) {

	testcases := []struct {
		desc        string
		inputMethod string
		target      string
		body        Book
		expected    int
	}{
		{"valid case", "POST", "http://localhost:8000/book", Book{"4", 1, "deciding decade",
			"penguin", "20/03/2010", &Author{1, "shani",
				"kumar", "30/04/2001", "sk"}},
			http.StatusCreated},
		{"invalid bookId", "POST", "http://localhost:8000/book", Book{"-4", 1, "deciding decade",
			"penguin", "20/03/2010", &Author{1, "shani",
				"kumar", "30/04/2001", "sk"}},
			http.StatusBadRequest},
		{"invalid author DOB", "POST", "http://localhost:8000/book", Book{"4", 1, "deciding decade",
			"penguin", "20/03/2010", &Author{1, "shani",
				"kumar", "30/00/2001", "sk"}},
			http.StatusBadRequest},
		{"invalid author's firstName", "POST", "http://localhost:8000/book", Book{"4", 1, "deciding decade",
			"penguin", "20/03/2010", &Author{1, "",
				"kumar", "30/04/2001", "sk"}},
			http.StatusBadRequest},
		{"not existing author", "POST", "http://localhost:8000/book", Book{"5", 1, "deciding decade",
			"penguin", "20/03/2010", &Author{1, "shani",
				"kumar", "30/00/2001", "sk"}},
			http.StatusBadRequest},
		{"invalid publication", "POST", "http://localhost:8000/book", Book{"6", 1, "deciding decade",
			"McGrowHill", "20/03/2010", &Author{1, "shani",
				"kumar", "30/00/2001", "sk"}},
			http.StatusBadRequest},
		{"invalid publishedDate", "POST", "http://localhost:8000/book", Book{"7", 1, "deciding decade",
			"McGrowHill", "20/03/2010", &Author{1, "shani",
				"kumar", "30/00/2001", "sk"}},
			http.StatusBadRequest},
		{"invalid title", "POST", "http://localhost:8000/book", Book{"7", 1, "",
			"McGrowHill", "20/03/2010", &Author{1, "shani",
				"kumar", "30/00/2001", "sk"}},
			http.StatusBadRequest},
	}

	for _, tc := range testcases {

		b, err := json.Marshal(tc.body)
		if err != nil {
			fmt.Println("error:", err)
		}

		req := httptest.NewRequest(tc.inputMethod, tc.target, bytes.NewBuffer(b))
		w := httptest.NewRecorder()
		PostBook(w, req)

		res := w.Result()
		if res.StatusCode != tc.expected {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc        string
		inputMethod string
		target      string
		body        Author
		expected    int
	}{
		{"valid author", "POST", "https://localhost:8000/author", Author{
			4, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, http.StatusBadRequest},
		{"exiting author", "POST", "https://localhost:8000/author", Author{
			3, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, http.StatusBadRequest},
		{"invalid firstname", "POST", "https://localhost:8000/author", Author{
			3, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, http.StatusBadRequest},
		{"invalid DOB", "POST", "https://localhost:8000/author", Author{
			3, "nilotpal", "mrinal", "20/00/1990", "Dark horse"}, http.StatusBadRequest},
	}

	for _, tc := range testcases {

		author, err := json.Marshal(tc.body)
		if err != nil {
			fmt.Println("error:", err)
		}

		req := httptest.NewRequest(tc.inputMethod, tc.target, bytes.NewBuffer(author))
		w := httptest.NewRecorder()
		PostAuthor(w, req)

		res := w.Result()
		if res.StatusCode != tc.expected {
			t.Errorf("failed for %s", tc.desc)
		}
	}
}

func TestPutBook(t *testing.T) {
	testcases := []struct {
		desc        string
		inputMethod string
		target      string
		body        Book
		expected    int
	}{
		{"updating a existing id", "PUT", "2", Book{"4", 1, "deciding decade",
			"penguin", "20/03/2008", &Author{1, "shani",
				"kumar", "30/04/2001", "sk"}}, http.StatusCreated},
		{"inserting a new id", "PUT", "1", Book{"4", 1, "life",
			"arihant", "20/03/2000", &Author{1, "shani",
				"kumar", "30/04/2001", "sk"}}, http.StatusCreated},
	}

	for _, tc := range testcases {

		ReqBody, err := json.Marshal(tc.body)
		if err != nil {
			fmt.Errorf("failed %v\n", err)
		}
		req := httptest.NewRequest(tc.inputMethod, "https://localhost:8000/book/{id}"+tc.target, bytes.NewBuffer(ReqBody))
		req = mux.SetURLVars(req, map[string]string{"id": tc.target})
		w := httptest.NewRecorder()
		PutBook(w, req)

		res := w.Result()
		if res.StatusCode != tc.expected {
			t.Errorf("failed for %s", tc.desc)
		}
		assert.Equal(t, tc.expected, res.StatusCode)
	}
}

func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc        string
		inputMethod string
		target      string
		author      Author
		expected    int
	}{
		{"updating a existing id", "PUT", "7", Author{1, "shani",
			"kumar", "30/04/2001", "sk"}, http.StatusCreated},
		{"inserting a new id", "PUT", "1", Author{10, "stephen",
			"hawkins", "30/04/2001", "sk"}, http.StatusCreated},
	}

	for _, tc := range testcases {

		ReqBody, err := json.Marshal(tc.author)
		if err != nil {
			fmt.Errorf("failed %v\n", err)
		}
		req := httptest.NewRequest(tc.inputMethod, "https://localhost:8000/author/{id}"+tc.target, bytes.NewBuffer(ReqBody))
		req = mux.SetURLVars(req, map[string]string{"id": tc.target})
		w := httptest.NewRecorder()
		PutAuthor(w, req)

		res := w.Result()
		if res.StatusCode != tc.expected {
			t.Errorf("failed for %s", tc.desc)
		}
		assert.Equal(t, tc.expected, res.StatusCode)
	}
}

func TestDeleteBook(t *testing.T) {
	testcases := []struct {
		desc        string
		inputMethod string
		target      string
		expected    int
	}{
		{"valid id", "DELETE", "4", http.StatusNoContent},
		{"invalid id", "DELETE", "-4", http.StatusBadRequest},
	}

	for _, tc := range testcases {

		req := httptest.NewRequest(tc.inputMethod, "https://localhost:8000/book/{id}"+tc.target, nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": tc.target})
		DeleteBook(w, req)

		res := w.Result()
		if res.StatusCode != tc.expected {
			t.Errorf("failed for %s", tc.desc)
		}

	}
}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc        string
		inputMethod string
		target      string
		expected    int
	}{
		{"valid authorId", "DELETE", "7", http.StatusNoContent},
		{"invalid authorId", "DELETE", "-3", http.StatusBadRequest},
	}

	for _, tc := range testcases {

		req := httptest.NewRequest(tc.inputMethod, "https://localhost:8000/author/{id}"+tc.target, nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": tc.target})
		DeleteAuthor(w, req)

		res := w.Result()
		if res.StatusCode != tc.expected {
			t.Errorf("failed for %s", tc.desc)
		}
		assert.Equal(t, tc.expected, res.StatusCode)
	}
}
