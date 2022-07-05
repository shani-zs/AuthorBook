package book

import (
	"errors"
	"projects/GoLang-Interns-2022/authorbook/driver"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"reflect"
	"testing"
)

func TestGetAllBook(t *testing.T) {
	Testcases := []struct {
		desc          string
		title         string
		includeAuthor string

		expected []entities.Book
	}{
		{"getting all books", "", "", []entities.Book{{1,
			1, "book one", "scholastic", "20/06/2018", entities.Author{}},
			{2, 1, "book two", "penguin", "20/08/2018", entities.Author{}}},
		},
		{"getting book with author and particular title", "book+two", "true", []entities.Book{
			{2, 1, "book two", "penguin", "20/08/2018", entities.Author{1, "shani",
				"kumar", "30/04/2001", "sk"}}},
		},
		{"getting book without author", "book+two", "true", []entities.Book{
			{2, 1, "book two", "penguin", "20/08/2018", entities.Author{}}},
		},
	}

	for _, tc := range Testcases {
		DB := driver.Connection()
		bookStore := New(DB)

		book := bookStore.GetAllBook(tc.title, tc.includeAuthor)

		if !reflect.DeepEqual(book, tc.expected) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
func TestGetBookByID(t *testing.T) {
	Testcases := []struct {
		desc     string
		targetID int

		expected entities.Book
	}{
		{"fetching book by id",
			1, entities.Book{1, 1, "book two", "penguin",
				"20/08/2018", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}}},

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
		{"valid case", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}}, nil},

		{"already existing book", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}}, errors.New("already existing")},
	}
	for _, tc := range testcases {
		DB := driver.Connection()
		bookStore := New(DB)

		id, err := bookStore.PostBook(&tc.body)

		if id == 0 && tc.err != err {
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
		{"creating a book", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}}, 1, nil},

		{"updating a book", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}}, 2, nil},
	}

	for _, tc := range testcases {
		DB := driver.Connection()
		bookStore := New(DB)

		id, err := bookStore.PutBook(&tc.body, tc.targetID)

		if id == 0 && tc.expectedErr != err {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
func TestDeleteBook(t *testing.T) {
	testcases := []struct {
		desc     string
		targetID int

		err error
	}{
		{"valid id", 1, nil},
		{"invalid id", -1, errors.New("invalid id")},
	}

	for _, tc := range testcases {
		DB := driver.Connection()
		bookStore := New(DB)

		id, err := bookStore.DeleteBook(tc.targetID)

		if id == 0 && tc.err != err {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
