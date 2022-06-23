package AuthorBook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http/httptest"
	"testing"
)

func TestGetAllBook(t *testing.T) {

	Testcases := []struct {
		desc        string
		methodInput string
		target      string
		//body        io.Reader
		expected []book
	}{
		{"test for fetching books", "GET", "http://localhost:8060/book", []book{{1,
			"Atomic habits", &author{"james ", "clear", "09-08-1990", "ddhjfh"},
			"penguin", string(2018)}, {2, "the defining decade",
			&author{"james ", "clear", "09-08-1990", "ddhjfh"},
			"scholastic", "04/05/2018"}}},
	}

	for _, tc := range Testcases {
		req := httptest.NewRequest(tc.methodInput, tc.target, nil)
		w := httptest.NewRecorder()
		GetAllBook(w, req)

		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)
		defer resp.Body.Close()

		var book []book

		err := json.Unmarshal(body, &book) //Unmarshal([]byte,*struct) -->opposite of unmarshal marshal(struct,[]byte)
		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, book, tc.expected)
		//if book != tc.expected {
		//	t.Errorf("failed for %s\n", tc.desc)
		//}

	}

}

func TestGetBookById(t *testing.T) {
	Testcases := []struct {
		desc        string
		methodInput string
		target      string
		expected    book
	}{
		{"test for fetching book by id", "GET", "http://localhost:8060/book/{23}", book{1, "Atomic habits", &author{"james", "clear", "09-08-1990", "ddhjfh"}, "penguin", "06-08-2018"}},
		//
	}

	for _, tc := range Testcases {

		req := httptest.NewRequest(tc.methodInput, tc.target, nil)
		w := httptest.NewRecorder()
		GetBookById(w, req)

		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)

		defer resp.Body.Close()

		var book book
		err := json.Unmarshal(body, &book) //Unmarshal([]byte,*struct)
		if err != nil {
			log.Fatal(err)
		}

		if book != tc.expected {
			t.Errorf("failed for %s\n", tc.desc)
		}

	}
}

func TestPostBook(t *testing.T) {

	testcases := []struct {
		desc        string
		inputMethod string
		target      string
		body        book
		expected    int
	}{
		{"test for invalid id", "POST", "http://localhost:8060/book/", book{-3, "deciding decade",
			&author{"meg", "jag", "08/08/1970", "jj"}, "scholastic", "07/08/2006"}, 404},
		{"test for invalid authors DOB", "POST", "http://localhost:8060/book/", book{4,
			"deciding decade", &author{"meg", "jag", "08/00/1970", "jj"}, "scholastic", "07/08/2006"}, 404},
		{"test for invalid authors first name", "POST", "http://localhost:8060/book/", book{3, "deciding decade",
			&author{" ", "jag", "08/08/1970", "jj"}, "penguin", "07/08/2006"}, 404},
		{"test for invalid last name", "POST", "http://localhost:8060/book/", book{5, "deciding decade",
			&author{"meg", "jag", "08/08/1970", "jj"}, "scholastic", "07/08/2006"}, 404},
		{"test for not existing author", "POST", "http://localhost:8060/book/", book{6, "deciding",
			&author{"shani", "kumar", "08/08/1970", "jj"}, "arihant", "07/08/2006"}, 404},
		{"test for invalid publication", "POST", "http://localhost:8060/book/", book{7,
			"deciding decade", &author{"meg", "jag", "08-08-1970", "jj"}, "McGrowHill", "07-08-2006"}, 404},
		{"test for invalid publication date", "POST", "http://localhost:8060/book/", book{8,
			"deciding decade", &author{"meg", "jag", "08/08/1970", "jj"}, "scholastic", "00/08/006"}, 404},
		{"test for valid case", "POST", "http://localhost:8060/book/", book{2, "decding decade",
			&author{"meg", "jag", "08-08-1970", "jj"}, "scholastic", "07/08/2006"}, 201},
		{"test for invalid title", "POST", "http://localhost:8060/book/", book{7, " ",
			&author{"meg", "jag", "08/08/1970", "jj"}, "scholastic", "07/08/2006"}, 404},
		{"test for invalid range of publication date", "POST", "http://localhost:8060/book/", book{7, " ",
			&author{"meg", "jag", "08/08/1970", "jj"}, "scholastic", "07/08/1870"}, 404},
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
		//body,_:=io.ReadAll(res.Body)
		if res.StatusCode != tc.expected {
			t.Errorf("failed for %s", tc.desc)
		}
	}
}

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc        string
		inputMethod string
		target      string
		body        *author
		expected    int
	}{
		{"test for valid author", "POST", "https://localhost:8060/author", &author{
			"nilotpal", "mrinal", "20/05/1990", "Dark horse"}, 201},
		{"test for exiting author", "POST", "https://localhost:8060/author", &author{"nilotpal",
			"mrinal", "20/05/1990", "Dark horse"}, 404},
		{"invalid first name", "POST", "https://localhost:8060/author", &author{" ",
			"mrinal", "20/05/1990", "Dark horse"}, 404},
		{"test for invalid lastname", "POST", "https://localhost:8060/author", &author{
			"manoj", " ", "20/05/1990", "Dark horse"}, 404},
		{"test for invalid DOB", "POST", "https://localhost:8060/author", &author{"manoj",
			"kumar", "-00/05/1990", "Dark horse"}, 404},
	}

	for _, tc := range testcases {

		b, err := json.Marshal(tc.body)
		if err != nil {
			fmt.Println("error:", err)
		}

		req := httptest.NewRequest(tc.inputMethod, tc.target, bytes.NewBuffer(b))
		w := httptest.NewRecorder()
		PostAuthor(w, req)

		res := w.Result()
		//body,_:=io.ReadAll(res.Body)
		if res.StatusCode != tc.expected {
			t.Errorf("failed for %s", tc.desc)
		}
	}
}
