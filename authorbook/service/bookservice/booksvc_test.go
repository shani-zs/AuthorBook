package bookservice

import (
	"errors"
	"github.com/golang/mock/gomock"
	"log"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/store"
	"reflect"
	"testing"
)

func TestGetAllBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAuthorStore := store.NewMockAuthorStorer(ctrl)
	mockBookStore := store.NewMockBookStorer(ctrl)
	mock := New(mockBookStore, mockAuthorStore)

	Testcases := []struct {
		desc          string
		title         string
		includeAuthor string

		expected   []entities.Book
		expetedErr error
	}{
		{desc: "getting all books", title: "", includeAuthor: "", expected: []entities.Book{},
			expetedErr: errors.New("empty")},
		{desc: "getting book with author and particular", title: "book two", includeAuthor: "",
			expected: []entities.Book{}, expetedErr: errors.New("empty"),
		},
		//{desc: "getting book with author and particular title", title: "book", includeAuthor: "true",
		//	expected: []entities.Book{}, expetedErr: errors.New("empty"),
		//},
	}

	for _, tc := range Testcases {
		if tc.title == "" {
			mockBookStore.EXPECT().GetAllBook().Return(tc.expected, tc.expetedErr)
		}
		if tc.title == "book two" {
			mockBookStore.EXPECT().GetBooksByTitle(tc.title).Return(tc.expected, tc.expetedErr)
		}
		//if tc.includeAuthor == "true" {
		//	mockBookStore.EXPECT().GetBooksByTitle(tc.title).Return(tc.expected, nil)
		//	mockAuthorStore.EXPECT().IncludeAuthor(1).Return(tc.expected, tc.expetedErr)
		//}

		books, err := mock.GetAllBook(tc.title, tc.includeAuthor)
		if err != nil {
			log.Print(err)
		}
		if !reflect.DeepEqual(books, tc.expected) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestGetBookByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAuthorStore := store.NewMockAuthorStorer(ctrl)
	mockBookStore := store.NewMockBookStorer(ctrl)
	mock := New(mockBookStore, mockAuthorStore)

	Testcases := []struct {
		desc     string
		targetID int

		expectedBody entities.Book
		expectedErr  error
	}{
		{desc: "fetching book by id",
			targetID: 1, expectedBody: entities.Book{}, expectedErr: errors.New("invalid id"),
		},
		{"invalid id", -1, entities.Book{}, errors.New("invalid id")},
	}

	for _, tc := range Testcases {
		if tc.targetID == 1 {
			mockBookStore.EXPECT().GetBookByID(tc.targetID).Return(tc.expectedBody, tc.expectedErr)
		}
		book, _ := mock.GetBookByID(tc.targetID)

		if !reflect.DeepEqual(book, tc.expectedBody) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAuthorStore := store.NewMockAuthorStorer(ctrl)
	mockBookStore := store.NewMockBookStorer(ctrl)
	mock := New(mockBookStore, mockAuthorStore)

	testcases := []struct {
		desc  string
		input entities.Book

		expected    entities.Book
		expectedErr error
	}{
		{desc: "success case", input: entities.Book{BookID: 0, AuthorID: 1, Title: "deciding decade",
			Publication: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{}},
			expected: entities.Book{BookID: 12, AuthorID: 1, Title: "deciding decade", Publication: "penguin",
				PublishedDate: "20/03/2010", Author: entities.Author{}}, expectedErr: nil,
		},
		{desc: "invalid publication", input: entities.Book{BookID: 1, AuthorID: 1, Title: "deciding decade",
			Publication: "pen", PublishedDate: "20/03/2010", Author: entities.Author{}},
			expected: entities.Book{}, expectedErr: nil,
		},
		{desc: "error", input: entities.Book{AuthorID: 1, Title: "deciding decade",
			Publication: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{}},
			expectedErr: errors.New("something"),
		},
	}
	for _, tc := range testcases {
		if tc.desc != "invalid publication" {
			mockBookStore.EXPECT().Post(&tc.input).Return(tc.expected.BookID, tc.expectedErr)
		}
		book, _ := mock.Post(&tc.input)
		if !reflect.DeepEqual(book, tc.expected) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPut(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAuthorStore := store.NewMockAuthorStorer(ctrl)
	mockBookStore := store.NewMockBookStorer(ctrl)
	mock := New(mockBookStore, mockAuthorStore)

	testcases := []struct {
		desc    string
		input   entities.Book
		inputID int

		expected    entities.Book
		expectedErr error
	}{
		{desc: "success case", input: entities.Book{BookID: 12, AuthorID: 1, Title: "deciding decade",
			Publication: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{}}, inputID: 1,
			expected: entities.Book{BookID: 12, AuthorID: 1, Title: "deciding decade", Publication: "penguin",
				PublishedDate: "20/03/2010", Author: entities.Author{}}, expectedErr: nil,
		},
		{desc: "invalid publication", input: entities.Book{BookID: 1, AuthorID: 1, Title: "deciding decade",
			Publication: "pen", PublishedDate: "20/03/2010", Author: entities.Author{}},
			expected: entities.Book{}, expectedErr: nil,
		},
		{desc: "error", input: entities.Book{AuthorID: 1, Title: "deciding decade",
			Publication: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{}},
			expectedErr: errors.New("something went wrong"),
		},
	}
	for _, tc := range testcases {
		if tc.desc != "invalid publication" {
			mockBookStore.EXPECT().Put(&tc.input, tc.inputID).Return(tc.expected.BookID, tc.expectedErr)
		}
		book, _ := mock.Put(&tc.input, tc.inputID)
		if !reflect.DeepEqual(book, tc.expected) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAuthorStore := store.NewMockAuthorStorer(ctrl)
	mockBookStore := store.NewMockBookStorer(ctrl)
	mock := New(mockBookStore, mockAuthorStore)

	testcases := []struct {
		desc    string
		inputID int

		expectedID  int
		expectedErr error
	}{
		{"valid id", 1, 1, nil},
		{"invalid id", -1, -1, nil},
		{"error case", 1, -1, errors.New("something went wrong")},
	}

	for _, tc := range testcases {
		if tc.desc != "invalid id" {
			mockBookStore.EXPECT().Delete(tc.inputID).Return(tc.expectedID, tc.expectedErr)
		}
		id, _ := mock.Delete(tc.inputID)

		if !reflect.DeepEqual(id, tc.expectedID) {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
