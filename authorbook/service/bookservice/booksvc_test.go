package bookservice

import (
	"errors"
	"projects/GoLang-Interns-2022/authorbook/driver"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/store/book"
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
		bookStore := book.New(DB)
		bookService := New(bookStore)

		book := bookService.GetAllBook(tc.title,tc.includeAuthor)


		if !reflect.DeepEqual(book,tc.expected) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestGetBookByID(t *testing.T) {
	Testcases := []struct {
		desc     string
		targetID string

		expectedBody entities.Book
		expectedErr  error
	}{
		{"fetching book by id",
			"1", entities.Book{1, 1, "book two", "penguin",
			"20/08/2018", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}}, nil},

		{"invalid id", "-1", entities.Book{}, errors.New("invalid id")},
	}

	for _, tc := range Testcases {

			DB := driver.Connection()
		bookStore := book.New(DB)
		bookService := New(bookStore)

		book := bookService.GetBookBYID()


		if !reflect.DeepEqual(book,tc.expected)   {
			t.Errorf("failed for %v\n", tc.desc)
		}
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
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			nil},
		{"already existing book", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			errors.New("already existing")},
		{"invalid bookID", entities.Book{-4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			errors.New("invalid bookID")},
		{"invalid author's DOB", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/00/2001", "sk"}},
			errors.New("include DOB")},
		{"invalid title", entities.Book{4, 1, "", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			errors.New("title is empty")},
		{"invalid publication", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			errors.New("invalid publication")},
		{"invalid published date", entities.Book{4, 1, "deciding decade", "penguin",
			"00/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			errors.New("invalid published date")},
	}
	for _, tc := range testcases {

	}
}

func TestPutBook(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Book

		err error
	}{
		{"inserting a book", entities.Book{5, 2, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "james", "clear", "04/04/1990", "sk"}},
			nil},
		{"updating a book", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			errors.New("already existing")},
		{"invalid bookID", entities.Book{-4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			errors.New("invalid bookID")},
		{"invalid author's DOB", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/00/2001", "sk"}},
			errors.New("include DOB")},
		{"invalid title", entities.Book{4, 1, "", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			errors.New("title is empty")},
		{"invalid publication", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			errors.New("invalid publication")},
		{"invalid published date", entities.Book{4, 1, "deciding decade", "penguin",
			"00/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			errors.New("invalid published date")},
	}
	for _, tc := range testcases {

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

	}

}
