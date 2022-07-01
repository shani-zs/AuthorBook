package bookhttp

import (
	"net/http"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"testing"
)

func TestGetAllBook(t *testing.T) {
	Testcases := []struct {
		desc          string
		title         string
		includeAuthor string

		expectedBooks      []entities.Book
		expectedStatusCode int
	}{
		{"getting all books", "", "", []entities.Book{{1,
			1, "book one", "scholastic", "20/06/2018", entities.Author{}},
			{2, 1, "book two", "penguin", "20/08/2018", entities.Author{}}},
			http.StatusOK},
		{"getting book with author and particular title", "book+two", "true", []entities.Book{
			{2, 1, "book two", "penguin", "20/08/2018", entities.Author{1, "shani",
				"kumar", "30/04/2001", "sk"}}}, http.StatusOK},
		{"getting book without author", "book+two", "true", []entities.Book{
			{2, 1, "book two", "penguin", "20/08/2018", entities.Author{}}}, http.StatusOK},
	}

	for _, tc := range Testcases {

	}
}

func TestGetBookByID(t *testing.T) {
	Testcases := []struct {
		desc     string
		targetID string

		expectedBody       entities.Book
		expectedStatusCode int
	}{
		{"fetching book by id",
			"1", entities.Book{1, 1, "book two", "penguin", "20/08/2018",
				entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}}, http.StatusOK},

		{"invalid id", "-1", entities.Book{}, http.StatusBadRequest},
	}

	for _, tc := range Testcases {

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

	}
}

func TestDeleteBook(t *testing.T) {
	testcases := []struct {
		desc     string
		targetID int

		expectedStatus int
	}{
		{"valid id", 1, http.StatusNoContent},
		{"invalid id", -1, http.StatusBadRequest},
	}

	for _, tc := range testcases {

	}

}
