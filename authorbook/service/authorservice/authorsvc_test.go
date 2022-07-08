package authorservice

import (
	"errors"
	"github.com/golang/mock/gomock"
	"log"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/store"
	"testing"
)

func TestPostAuthor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockAuthorStorer(ctrl)
	mock := New(mockStore) //defining the type of interface

	testcases := []struct {
		desc string
		body entities.Author

		expectedAuthor entities.Author
		expectedID     int
		expectedErr    error
	}{
		{desc: "valid author", body: entities.Author{
			AuthorID: 4, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"},
			expectedAuthor: entities.Author{AuthorID: 4, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990",
				PenName: "Dark horse"}, expectedID: 4, expectedErr: nil},

		{desc: "existing author", body: entities.Author{
			AuthorID: 4, FirstName: "nilotpal", LastName: "mrinal", DOB: "01/05/1990", PenName: "Dark horse"},
			expectedAuthor: entities.Author{}, expectedID: -1, expectedErr: errors.New("already exists")},

		{desc: "invalid firstname", body: entities.Author{
			AuthorID: 5, FirstName: "", LastName: "mrinal", DOB: "20/01/1990", PenName: "Dark horse"},
			expectedAuthor: entities.Author{}, expectedID: -1, expectedErr: errors.New("invalid constraints")},

		{desc: "invalid DOB", body: entities.Author{
			AuthorID: 5, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/01/0", PenName: "Dark horse"},
			expectedAuthor: entities.Author{}, expectedID: -1, expectedErr: errors.New("invalid constraints")},
	}

	for _, tc := range testcases {
		if tc.body.AuthorID != 5 {
			mockStore.EXPECT().PostAuthor(tc.body).Return(tc.expectedID, tc.expectedErr)
		}

		a, err := mock.PostAuthor(tc.body)
		if err != nil {
			log.Print(err)
		}

		if a != tc.expectedAuthor {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPutAuthor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockAuthorStorer(ctrl)
	mock := New(mockStore)

	testcases := []struct {
		desc     string
		body     entities.Author
		targetID int

		expected    entities.Author
		expectedErr error
	}{
		{desc: "existing author", body: entities.Author{
			AuthorID: 4, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"},
			targetID: 5, expected: entities.Author{AuthorID: 4, FirstName: "nilotpal", LastName: "mrinal",
				DOB: "20/05/1990", PenName: "Dark horse"}, expectedErr: nil,
		},
		{desc: "not existing author", body: entities.Author{
			AuthorID: 4, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"},
			targetID: 5, expected: entities.Author{}, expectedErr: errors.New("already exist"),
		},
		{desc: "invalid firstname", body: entities.Author{
			AuthorID: 3, FirstName: "", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"},
			targetID: 5, expected: entities.Author{}, expectedErr: errors.New("invalid constraints"),
		},
		{desc: "invalid DOB", body: entities.Author{
			AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/00/1990", PenName: "Dark horse"},
			targetID: 5, expected: entities.Author{}, expectedErr: errors.New("invalid constraints"),
		},
	}

	for _, tc := range testcases {
		if tc.body.AuthorID == 4 {
			mockStore.EXPECT().PutAuthor(tc.body, tc.targetID).Return(tc.body.AuthorID, tc.expectedErr)
		}

		a, err := mock.PutAuthor(tc.body, tc.targetID)
		if err != nil {
			log.Print(err)
		}

		if a != tc.expected {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestDeleteAuthor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := store.NewMockAuthorStorer(ctrl)
	mock := New(mockStore)

	testcases := []struct {
		desc     string
		targetID int

		expectedRowsAffected int
		expectedErr          error
	}{
		{"valid authorId", 4, 1, nil},
		{"invalid authorId", -1, 0, errors.New("invalid id")},
	}

	for _, tc := range testcases {
		if tc.targetID == 4 {
			mockStore.EXPECT().DeleteAuthor(tc.targetID).Return(tc.expectedRowsAffected, tc.expectedErr)
		}
		id, err := mock.DeleteAuthor(tc.targetID)

		if err != nil {
			log.Print(err)
		}

		if id != tc.expectedRowsAffected {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

type mockStore struct{}

func (m mockStore) PostAuthor(author2 entities.Author) (int, error) {
	if author2.AuthorID == 4 {
		return author2.AuthorID, nil
	}

	return -1, errors.New("invalid")
}

func (m mockStore) PutAuthor(author2 entities.Author, id int) (int, error) {
	if author2.AuthorID == 4 {
		return author2.AuthorID, nil
	}

	return -1, errors.New("invalid")
}

func (m mockStore) DeleteAuthor(id int) (int, error) {
	if id <= 0 {
		return -1, nil
	}

	return id, nil
}
