package bookservice

import (
	"errors"
	"log"
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
		b := New(mockStore{})
		book := b.GetAllBook(tc.title, tc.includeAuthor)
		if !reflect.DeepEqual(book, tc.expected) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestGetBookByID(t *testing.T) {
	Testcases := []struct {
		desc     string
		targetID int

		expectedBody entities.Book
		expectedErr  error
	}{
		{"fetching book by id",
			1, entities.Book{1, 1, "book two", "penguin",
			"20/08/2018", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}}, nil},

		{"invalid id", -1, entities.Book{}, errors.New("invalid id")},
	}

	for _, tc := range Testcases {
		b := New(mockStore{})
		book, err := b.GetBookByID(tc.targetID)
		if err != nil {
			log.Print(err)
		}

		if !reflect.DeepEqual(book, tc.expectedBody) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPostBook(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Book

		expectedBook entities.Book
	}{
		{"valid case", entities.Book{4, 1, "deciding decade", "penguin", "20/03/2010",
			entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			entities.Book{4, 1, "deciding decade", "penguin", "20/03/2010",
				entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},},

		{"already existing book", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			entities.Book{}},

		{"invalid bookID", entities.Book{-4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			entities.Book{}},

		{"invalid author's DOB", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/00/2001", "sk"}},
			entities.Book{}},
		{"invalid title", entities.Book{4, 1, "", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			entities.Book{}},
		{"invalid publication", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			entities.Book{}},
		{"invalid published date", entities.Book{4, 1, "deciding decade", "penguin",
			"00/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			entities.Book{}},
	}
	for _, tc := range testcases {
		b := New(mockStore{})
		book, err := b.PostBook(tc.body)
		if err != nil {
			log.Print(err)
		}

		if !reflect.DeepEqual(book, tc.expectedBook) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPutBook(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Book

		expectedBook entities.Book
	}{
		{"inserting a book", entities.Book{5, 2, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "james", "clear", "04/04/1990", "sk"}},
			entities.Book{5, 2, "deciding decade", "penguin",
				"20/03/2010", entities.Author{1, "james", "clear", "04/04/1990", "sk"}},},

		{"updating a book", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			entities.Book{}},
		{"invalid bookID", entities.Book{-4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			entities.Book{}},
		{"invalid author's DOB", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/00/2001", "sk"}},
			entities.Book{}},
		{"invalid title", entities.Book{4, 1, "", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			entities.Book{}},
		{"invalid publication", entities.Book{4, 1, "deciding decade", "penguin",
			"20/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			entities.Book{}},
		{"invalid published date", entities.Book{4, 1, "deciding decade", "penguin",
			"00/03/2010", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}},
			entities.Book{}},
	}
	for _, tc := range testcases {
		b := New(mockStore{})
		book, err := b.PostBook(tc.body)
		if err != nil {
			log.Print(err)
		}

		if !reflect.DeepEqual(book, tc.expectedBook) {
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
		{"valid id", 1, 1},
		{"invalid id", -1, -1},
	}

	for _, tc := range testcases {
		b := New(mockStore{})
		id, err := b.DeleteBook(tc.targetID)
		if err != nil {
			log.Print(err)
		}

		if !reflect.DeepEqual(id, tc.expected) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
