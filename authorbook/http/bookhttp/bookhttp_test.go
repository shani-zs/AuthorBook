package bookhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/service"
	"reflect"
	"strconv"
	"testing"
)

func TestGetAllBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockBookService(ctrl)
	mock := New(mockService)

	Testcases := []struct {
		desc          string
		title         string
		includeAuthor string

		expectedBooks  []entities.Book
		expectedErr    error
		expectedStatus int
	}{
		{desc: "success case", title: "", includeAuthor: "", expectedBooks: []entities.Book{{BookID: 1,
			AuthorID: 1, Title: "book one", Publication: "scholastic", PublishedDate: "20/06/2018",
			Author: entities.Author{}}, {BookID: 2, AuthorID: 1, Title: "book two", Publication: "penguin",
			PublishedDate: "20/08/2018", Author: entities.Author{}}}, expectedErr: nil, expectedStatus: http.StatusOK,
		},
		{desc: "invalid case", title: "book+two", includeAuthor: "true", expectedBooks: []entities.Book{},
			expectedErr: errors.New("does not exist"), expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range Testcases {
		req := httptest.NewRequest("GET", "localhost:8000/book?"+"title="+tc.title+"&"+"includeAuthor="+tc.includeAuthor, nil)
		w := httptest.NewRecorder()

		if tc.title == "" {
			mockService.EXPECT().GetAllBook(tc.title, tc.includeAuthor).Return(tc.expectedBooks, tc.expectedErr)
		}

		if tc.title == "book+two" {
			mockService.EXPECT().GetAllBook("book two", tc.includeAuthor).Return(tc.expectedBooks, tc.expectedErr)
		}

		mock.GetAllBook(w, req)

		res := w.Result()
		if !assert.Equal(t, tc.expectedStatus, res.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

func TestGetBookByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockBookService(ctrl)
	mock := New(mockService)

	Testcases := []struct {
		desc     string
		targetID string

		expected           entities.Book
		expectedStatusCode int
		expectedErr        error
	}{
		{desc: "fetching book by id", targetID: "1", expected: entities.Book{BookID: 1, AuthorID: 1, Title: "book two",
			Publication: "penguin", PublishedDate: "20/08/2018", Author: entities.Author{AuthorID: 1, FirstName: "shani",
				LastName: "kumar", DOB: "30/04/2001", PenName: "sk"}}, expectedStatusCode: http.StatusOK, expectedErr: nil,
		},
		{"invalid id", "-1", entities.Book{}, http.StatusBadRequest,
			errors.New("invalid id"),
		},
	}

	for _, tc := range Testcases {
		req := httptest.NewRequest("GET", "localhost:8000/book/{id}"+tc.targetID, nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": tc.targetID})

		id, _ := strconv.Atoi(tc.targetID)
		if tc.targetID != "-1" {
			mockService.EXPECT().GetBookByID(id).Return(tc.expected, tc.expectedErr)
		}
		mock.GetBookByID(w, req)

		result := w.Result()

		if !reflect.DeepEqual(tc.expectedStatusCode, result.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

func TestPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockBookService(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc string
		body entities.Book

		expected           entities.Book
		expectedErr        error
		expectedStatusCode int
	}{
		{desc: "invalid case", body: entities.Book{BookID: 0, AuthorID: 1, Title: "deciding decade", Publication: "penguin",
			PublishedDate: "20/03/2010", Author: entities.Author{}},
			expected: entities.Book{}, expectedErr: errors.New("something"),
			expectedStatusCode: http.StatusBadRequest,
		},

		{desc: "valid case", body: entities.Book{BookID: 0, AuthorID: 1, Title: "deciding decade",
			Publication: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{}},
			expected: entities.Book{BookID: 15, AuthorID: 1, Title: "deciding decade", Publication: "penguin",
				PublishedDate: "20/03/2010", Author: entities.Author{}},
			expectedErr: nil, expectedStatusCode: http.StatusCreated,
		},
		{desc: "unmarshalling error", body: entities.Book{}, expected: entities.Book{}, expectedErr: errors.New("something"),
			expectedStatusCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testcases {
		data, err := json.Marshal(tc.body)
		if err != nil {
			log.Printf("failed : %v", err)
		}

		if tc.desc == "unmarshalling error" {
			data = []byte("shani")
		}

		req := httptest.NewRequest("POST", "localhost:8000/book", bytes.NewBuffer(data))
		w := httptest.NewRecorder()
		if tc.desc == "invalid case" || tc.desc == "valid case" {
			mockService.EXPECT().Post(&tc.body).Return(tc.expected, tc.expectedErr)
		}
		mock.Post(w, req)

		result := w.Result()

		if !reflect.DeepEqual(tc.expectedStatusCode, result.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

func TestPut(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockBookService(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc    string
		input   entities.Book
		inputID string

		expected           entities.Book
		expectedErr        error
		expectedStatusCode int
	}{
		{desc: "invalid case", input: entities.Book{BookID: 0, AuthorID: 1, Title: "deciding decade", Publication: "penguin",
			PublishedDate: "20/03/2010", Author: entities.Author{}}, inputID: "2",
			expected: entities.Book{}, expectedErr: errors.New("something"),
			expectedStatusCode: http.StatusNotFound,
		},

		{desc: "valid case", input: entities.Book{BookID: 15, AuthorID: 1, Title: "deciding decade",
			Publication: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{}}, inputID: "4",
			expected: entities.Book{BookID: 15, AuthorID: 1, Title: "deciding decade", Publication: "penguin",
				PublishedDate: "20/03/2010", Author: entities.Author{}},
			expectedErr: nil, expectedStatusCode: http.StatusCreated,
		},
		{desc: "unmarshalling error", input: entities.Book{}, expected: entities.Book{}, expectedErr: errors.New("something"),
			expectedStatusCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testcases {
		data, err := json.Marshal(tc.input)
		if err != nil {
			log.Printf("failed : %v", err)
		}

		if tc.desc == "unmarshalling error" {
			data = []byte("shani")
		}

		req := httptest.NewRequest("PUT", "localhost:8000/book/{id}"+tc.inputID, bytes.NewBuffer(data))
		req = mux.SetURLVars(req, map[string]string{"id": tc.inputID})
		w := httptest.NewRecorder()

		id, _ := strconv.Atoi(tc.inputID)
		if tc.desc == "invalid case" || tc.desc == "valid case" {
			mockService.EXPECT().Put(&tc.input, id).Return(tc.expected, tc.expectedErr)
		}
		mock.Put(w, req)

		result := w.Result()

		if !reflect.DeepEqual(tc.expectedStatusCode, result.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}

func TestDeleteBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockBookService(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc    string
		inputID string

		expectedStatus int
		expectedErr    error
	}{
		{"valid id", "1", http.StatusNoContent, nil},
		{"invalid id", "-1", http.StatusBadRequest, errors.New("something wrong")},
		{"invalid case", "2", http.StatusNotFound, errors.New("something wrong")},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest("PUT", "localhost:8000/book/{id}"+tc.inputID, nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.inputID})
		w := httptest.NewRecorder()

		id, _ := strconv.Atoi(tc.inputID)
		if tc.desc == "valid id" || tc.desc == "invalid case" {
			mockService.EXPECT().Delete(id).Return(1, tc.expectedErr)
		}

		mock.Delete(w, req)

		result := w.Result()
		if !reflect.DeepEqual(tc.expectedStatus, result.StatusCode) {
			t.Errorf("failed for %s\n", tc.desc)
		}
	}
}
