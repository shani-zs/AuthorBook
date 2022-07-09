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
		{desc: "getting all books", title: "", includeAuthor: "", expected: []entities.Book{{BookID: 1,
			AuthorID: 1, Title: "book one", Publication: "scholastic", PublishedDate: "20/06/2018", Author: entities.Author{}},
			{BookID: 2, AuthorID: 1, Title: "book two", Publication: "penguin", PublishedDate: "20/08/2018",
				Author: entities.Author{}}},
		},
		{desc: "getting book with authorhttp and particular title", title: "book+two", includeAuthor: "true",
			expected: []entities.Book{{BookID: 2, AuthorID: 1, Title: "book two", Publication: "penguin",
				PublishedDate: "20/08/2018", Author: entities.Author{AuthorID: 1, FirstName: "shani",
					LastName: "kumar", DOB: "30/04/2001", PenName: "sk"}}},
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
		{desc: "fetching book by id",
			targetID: 1, expectedBody: entities.Book{BookID: 1, AuthorID: 1, Title: "book two", Publication: "penguin",
				PublishedDate: "20/08/2018", Author: entities.Author{AuthorID: 1, FirstName: "shani", LastName: "kumar",
					DOB: "30/04/2001", PenName: "sk"}}},

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
		{desc: "already existing book", body: entities.Book{BookID: 1, AuthorID: 1, Title: "deciding decade",
			Publication: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1, FirstName: "shani",
				LastName: "kumar", DOB: "30/04/2001", PenName: "sk"}}},
		{desc: "invalid bookID", body: entities.Book{BookID: -4, AuthorID: 1, Title: "deciding decade",
			Publication: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1, FirstName: "shani",
				LastName: "kumar", DOB: "30/04/2001", PenName: "sk"}}},
	}
	for _, tc := range testcases {
		b := New(mockStore{})

		book, err := b.PostBook(&tc.body)
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
		{desc: "inserting a book", body: entities.Book{BookID: 4, AuthorID: 1, Title: "decade", Publication: "penguin",
			PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1, FirstName: "shani", LastName: "kumar",
				DOB: "04/04/1990", PenName: "sk"}}, expectedBook: entities.Book{BookID: 4, AuthorID: 1, Title: "decade",
			Publication: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1,
				FirstName: "shani", LastName: "kumar", DOB: "04/04/1990", PenName: "sk"}}},

		{desc: "updating a book", body: entities.Book{BookID: 3, AuthorID: 1, Title: "book three", Publication: "penguin",
			PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1, FirstName: "shani", LastName: "kumar",
				DOB: "30/04/2001", PenName: "sk"}}},
	}
	for _, tc := range testcases {
		b := New(mockStore{})

		book, err := b.PostBook(&tc.body)
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
