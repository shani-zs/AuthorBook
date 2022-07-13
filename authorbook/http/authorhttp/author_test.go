package authorhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"testing"

	"net/http"
	"net/http/httptest"

	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/service"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

// TestPost : to test Post handler
func TestPost(t *testing.T) {
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
		{desc: "returning error from svc", input: entities.Author{AuthorID: 4, FirstName: "nilotpal", LastName: "mrinal",
			DOB: "20/01/1990", PenName: "Dark horse"}, expected: entities.Author{},
			expectedStatus: http.StatusBadRequest, expectedErr: errors.New("not valid constraints"),
		},
		{desc: "unmarshalling error ", input: entities.Author{}, expected: entities.Author{},
			expectedStatus: http.StatusBadRequest, expectedErr: nil,
		},
	}

	for _, tc := range testcases {
		data, _ := json.Marshal(tc.input)

		if tc.desc == "unmarshalling error " {
			data = []byte("hello")
		}

		req := httptest.NewRequest("POST", "localhost:8000/author", bytes.NewReader(data))
		w := httptest.NewRecorder()
		ctx := req.Context()

		if tc.input.AuthorID == 4 {
			mockService.EXPECT().Post(ctx, tc.input).Return(tc.expected, tc.expectedErr)
		} else {
			mockService.EXPECT().Post(ctx, tc.input).Return(tc.expected, tc.expectedErr).AnyTimes()
		}

		mock.Post(w, req)

		res := w.Result()
		if tc.expectedStatus != res.StatusCode {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

// TestPut : to test the put handler
func TestPut(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockAuthorService(ctrl)
	mock := New(mockService)

	testcases := []struct {
		desc           string
		input          entities.Author
		TargetID       string
		expected       entities.Author
		expectedStatus int
		expectedErr    error
	}{
		{desc: "valid case:", input: entities.Author{
			AuthorID: 3, FirstName: "amit", LastName: "kumar", DOB: "20/01/1990", PenName: "Dark horse"},
			TargetID: "4", expected: entities.Author{AuthorID: 3, FirstName: "amit",
				LastName: "kumar", DOB: "20/01/1990", PenName: "Dark horse"}, expectedStatus: http.StatusCreated,
			expectedErr: nil,
		},
		{desc: "strconv error", input: entities.Author{AuthorID: 3, FirstName: "kumar", LastName: "vis",
			DOB: "20/01/1990", PenName: "Dark horse"}, expected: entities.Author{},
			expectedStatus: http.StatusBadRequest, expectedErr: nil,
		},
		{desc: "unmarshalling error ", input: entities.Author{}, expected: entities.Author{},
			expectedStatus: http.StatusBadRequest, expectedErr: nil,
		},
		{desc: "error from svc layer", input: entities.Author{}, TargetID: "5", expected: entities.Author{},
			expectedStatus: http.StatusNotFound, expectedErr: errors.New("invalid error"),
		},
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
		id, _ := strconv.Atoi(tc.TargetID)
		ctx := req.Context()

		mockService.EXPECT().Put(ctx, tc.input, id).Return(tc.expected, tc.expectedErr).AnyTimes()

		mock.Put(w, req)

		res := w.Result()
		if tc.expectedStatus != res.StatusCode {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

// TestDelete : to test the delete handler
func TestDelete(t *testing.T) {
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
		{desc: "invalid authorId", expectedStatus: http.StatusBadRequest, expectedErr: errors.New("invalid")},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest("DELETE", "localhost:8000/author/{id}"+tc.target, nil)
		req = mux.SetURLVars(req, map[string]string{"id": tc.target})
		w := httptest.NewRecorder()

		id, err := strconv.Atoi(tc.target)
		if err != nil {
			log.Print(err)
		}

		ctx := req.Context()
		mockService.EXPECT().Delete(ctx, id).Return(tc.expectedStatus, tc.expectedErr).AnyTimes()

		mock.Delete(w, req)

		res := w.Result()
		if tc.expectedStatus != res.StatusCode {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
