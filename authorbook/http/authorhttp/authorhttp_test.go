package author

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/service"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

func TestPostAuthor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockAuthorService(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc  string
		input entities.Author

		expected       entities.Author
		expectedStatus int
		expectedErr    error
	}{
		{desc: "valid case:", input: entities.Author{
			AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/01/1990", PenName: "Dark horse"},
			expected: entities.Author{
				AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/01/1990", PenName: "Dark horse"},
			expectedStatus: http.StatusCreated, expectedErr: nil,
		},
		{desc: "returning error from svc", input: entities.Author{AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal",
			DOB: "20/01/1990", PenName: "Dark horse"}, expected: entities.Author{},
			expectedStatus: http.StatusBadRequest, expectedErr: errors.New("not valid constraints"),
		},
		{desc: "unmarshalling error ", input: entities.Author{}, expected: entities.Author{},
			expectedStatus: http.StatusBadRequest, expectedErr: nil,
		},
		//{desc: "exiting author", body: entities.Author{
		//	AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"}, expectedStatus: http.StatusBadRequest},
		//{desc: "invalid firstname", body: entities.Author{
		//	AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"}, expectedStatus: http.StatusBadRequest},
		//{desc: "valid author", body: entities.Author{
		//	AuthorID: 4, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"}, expectedStatus: http.StatusCreated},
	}

	for _, tc := range testcases {
		data, err := json.Marshal(tc.input)
		if err != nil {
			log.Print(err)
		}

		if tc.desc == "unmarshalling error " {
			data = []byte("hello")
		}

		req := httptest.NewRequest("POST", "localhost:8000/author", bytes.NewReader(data))
		w := httptest.NewRecorder()

		if tc.input.AuthorID == 3 {
			mockService.EXPECT().PostAuthor(tc.input).Return(tc.expected, tc.expectedErr)
		}

		mock.Post(w, req)

		var a entities.Author

		res := w.Result()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Print(err)
		}

		err = json.Unmarshal(body, &a)
		if err != nil {
			log.Print(err)
		}

		if tc.expectedStatus != res.StatusCode || a != tc.expected {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

// TODO: see it later on
func TestPutAuthor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockAuthorService(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc     string
		input    entities.Author
		TargetID string

		expected       entities.Author
		expectedStatus int
		expectedErr    error
	}{
		{desc: "valid case:", input: entities.Author{
			AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/01/1990", PenName: "Dark horse"}, TargetID: "4",
			expected: entities.Author{
				AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/01/1990", PenName: "Dark horse"},
			expectedStatus: http.StatusCreated, expectedErr: nil,
		},
		//{desc: "returning error from svc", input: entities.Author{AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal",
		//	DOB: "20/01/1990", PenName: "Dark horse"}, expectedStatus: entities.Author{},
		//	expectedStatus: http.StatusBadRequest, expectedErr: errors.New("not valid constraints"),
		//},
		//{desc: "unmarshalling error ", input: entities.Author{}, expectedStatus: entities.Author{},
		//	expectedStatus: http.StatusBadRequest, expectedErr: nil,
		//},
	}
	for _, tc := range testcases {
		data, err := json.Marshal(tc.input)
		if err != nil {
			log.Print(err)
		}

		if tc.desc == "unmarshalling error " {
			data = []byte("hello")
		}

		req := httptest.NewRequest("PUT", "localhost:8000/author/{id}"+tc.TargetID, bytes.NewReader(data))
		req = mux.SetURLVars(req, map[string]string{"id": tc.TargetID})
		w := httptest.NewRecorder()

		for tc.input.AuthorID == 3 {
			mockService.EXPECT().PutAuthor(tc.input, tc.input.AuthorID).Return(tc.expected, tc.expectedErr)
		}

		mock.Put(w, req)

		res := w.Result()
		if tc.expectedStatus != res.StatusCode {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestDeleteAuthor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockAuthorService(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc   string
		target string

		expectedStatus int
		expectedErr    error
	}{
		{"valid authorId", "4", http.StatusNoContent, nil},
		{"invalid authorId", "-3", http.StatusBadRequest, errors.New("invalid")},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest("DELETE", "localhost:8000/author/{id}"+tc.target, nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.target})
		w := httptest.NewRecorder()

		id, err := strconv.Atoi(tc.target)
		if err != nil {
			log.Print(err)
		}

		mockService.EXPECT().DeleteAuthor(id).Return(tc.expectedStatus, tc.expectedErr)

		mock.Delete(w, req)

		res := w.Result()
		if tc.expectedStatus != res.StatusCode {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
