package book

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"projects/GoLang-Interns-2022/authorbook/driver"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"testing"
)

func TestGetAllBook(t *testing.T) {
	Testcases := []struct {
		desc string

		expected    []entities.Book
		expectedErr error
	}{
		{desc: "getting all books", expected: []entities.Book{{BookID: 1,
			AuthorID: 1, Title: "book one", Publication: "penguin", PublishedDate: "20/06/2000",
			Author: entities.Author{}}}, expectedErr: nil,
		},
		{desc: "getting book with author and particular title",
			expected: []entities.Book{{BookID: 1, AuthorID: 1, Title: "book one", Publication: "penguin",
				PublishedDate: "20/06/2000", Author: entities.Author{AuthorID: 1, FirstName: "shani",
					LastName: "kumar", DOB: "30/04/1980", PenName: "jack"}}}, expectedErr: nil,
		},
	}

	var (
		book1=
		books = sqlmock.NewRows([]string{"id", "author_id", "title", "publication", "published_date"}).
			AddRow()
	)

	for _, tc := range Testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Print(err)
		}
		mock.ExpectExec("")
		bs := New(db)
		books, err := bs.GetAllBook(context.TODO())
		if err != nil {
			log.Print(err)
		}

	}
}
func TestGetBookByID(t *testing.T) {
	Testcases := []struct {
		desc     string
		targetID int

		expected entities.Book
	}{
		{desc: "fetching book by id",
			targetID: 1, expected: entities.Book{BookID: 1,
			AuthorID: 1, Title: "book one", Publication: "penguin", PublishedDate: "20/06/2000"}},

		{"invalid id", -1, entities.Book{}},
	}

	for _, tc := range Testcases {
		DB := driver.Connection()
		bookStore := New(DB)

		book := bookStore.GetBookByID(tc.targetID)

		if book != tc.expected {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
func TestPostBook(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Book

		err error
	}{
		{desc: "valid case", body: entities.Book{BookID: 4, AuthorID: 1, Title: "deciding decade", Publication: "penguin",
			PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1, FirstName: "shani", LastName: "kumar",
				DOB: "30/04/2001", PenName: "sk"}}},

		{desc: "already existing book", body: entities.Book{BookID: 4, AuthorID: 1, Title: "deciding decade",
			Publication: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1, FirstName: "shani",
				LastName: "kumar", DOB: "30/04/2001", PenName: "sk"}}},
	}
	for _, tc := range testcases {
		DB := driver.Connection()
		bookStore := New(DB)

		_, err := bookStore.Post(&tc.body)

		if tc.err != err {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
func TestPutBook(t *testing.T) {
	testcases := []struct {
		desc     string
		body     entities.Book
		targetID int

		expectedErr error
	}{
		{desc: "creating a book", body: entities.Book{BookID: 4, AuthorID: 1, Title: "deciding", Publication: "penguin",
			PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1, FirstName: "shani", LastName: "kumar",
				DOB: "30/04/2001", PenName: "sk"}}, targetID: 13},

		{desc: "updating a book", body: entities.Book{BookID: 4, AuthorID: 1, Title: "deciding decade",
			Publication: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1,
				FirstName: "shani", LastName: "kumar", DOB: "30/04/2001", PenName: "sk"}}, targetID: 14},
	}

	for _, tc := range testcases {
		DB := driver.Connection()
		bookStore := New(DB)

		_, err := bookStore.Put(&tc.body, tc.targetID)

		if tc.expectedErr != err {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
func TestDeleteBook(t *testing.T) {
	testcases := []struct {
		desc     string
		targetID int

		expected int
	}{
		{"valid id", 14, 1},
		{"invalid id", -1, 0},
	}

	for _, tc := range testcases {
		DB := driver.Connection()
		bookStore := New(DB)

		id, _ := bookStore.Delete(tc.targetID)

		if id != tc.expected {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
