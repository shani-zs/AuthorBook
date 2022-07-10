package bookservice

import (
	"errors"
	"github.com/golang/mock/gomock"
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
		{desc: "getting book with author and particular title", title: "book", includeAuthor: "true",
			expected: []entities.Book{}, expetedErr: errors.New("empty"),
		},
	}

	for _, tc := range Testcases {
		mockBookStore.EXPECT().GetAllBook().Return(tc.expected, tc.expetedErr).AnyTimes()
		mockBookStore.EXPECT().GetBooksByTitle(tc.title).Return(tc.expected, tc.expetedErr).AnyTimes()

		books, _ := mock.GetAllBook(tc.title, tc.includeAuthor)

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

		mockBookStore.EXPECT().GetBookByID(tc.targetID).Return(tc.expectedBody, tc.expectedErr).AnyTimes()
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

		expected     entities.Book
		expectedErr  error
		expectedErr1 error
	}{
		{desc: "success case", input: entities.Book{BookID: 0, AuthorID: 1, Title: "deciding decade",
			Publication: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1, FirstName: "shani",
				LastName: "kumar", DOB: "30/05/1999", PenName: "sk"}},
			expected: entities.Book{BookID: 12, AuthorID: 1, Title: "deciding decade", Publication: "penguin",
				PublishedDate: "20/03/2010", Author: entities.Author{AuthorID: 1, FirstName: "shani",
					LastName: "kumar", DOB: "30/05/1999", PenName: "sk"}}, expectedErr: nil, expectedErr1: nil,
		},

		{desc: "author does not exist", input: entities.Book{BookID: 1, AuthorID: 3, Title: "deciding decade",
			Publication: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{}},
			expected: entities.Book{}, expectedErr: errors.New("issue"), expectedErr1: errors.New("author does not exist"),
		},

		{desc: "invalid publication", input: entities.Book{BookID: 1, AuthorID: 3, Title: "deciding decade",
			Publication: "pen", PublishedDate: "20/03/2010", Author: entities.Author{}},
			expected: entities.Book{}, expectedErr: nil, expectedErr1: nil,
		},
	}
	for _, tc := range testcases {
		mockBookStore.EXPECT().Post(&tc.input).Return(tc.expected.BookID, tc.expectedErr).AnyTimes()
		mockAuthorStore.EXPECT().IncludeAuthor(tc.input.AuthorID).Return(tc.input.Author, tc.expectedErr1).AnyTimes()

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
