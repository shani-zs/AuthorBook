package authorservice

import (
	"errors"
	"log"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"testing"
)

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author

		expected entities.Author
	}{
		{desc: "valid author", body: entities.Author{
			AuthorID: 4, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"},
			expected: entities.Author{AuthorID: 4, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990",
				PenName: "Dark horse"}},
		{desc: "existing author", body: entities.Author{
			AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "00/05/1990", PenName: "Dark horse"}},
		{desc: "invalid firstname", body: entities.Author{
			AuthorID: 3, FirstName: " ", LastName: "mrinal", DOB: "20/00/1990", PenName: "Dark horse"}},
		{desc: "invalid DOB", body: entities.Author{
			AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/01/0", PenName: "Dark horse"}},
	}

	for _, tc := range testcases {
		m := New(mockStore{})

		id, err := m.PostAuthor(tc.body)
		if err != nil {
			log.Print(err)
		}

		if id != tc.expected {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author

		expected entities.Author
	}{
		{desc: "valid author", body: entities.Author{
			AuthorID: 4, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"},
			expected: entities.Author{AuthorID: 4, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990",
				PenName: "Dark horse"}},
		{desc: "existing author", body: entities.Author{
			AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"}},
		{desc: "invalid firstname", body: entities.Author{
			AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"}},
		{desc: "invalid DOB", body: entities.Author{
			AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/00/1990", PenName: "Dark horse"}},
		{desc: "valid author", body: entities.Author{
			AuthorID: 5, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"}},
	}

	for _, tc := range testcases {
		m := New(mockStore{})

		a, err := m.PutAuthor(tc.body)
		if err != nil {
			log.Print(err)
		}

		if a != tc.expected {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc   string
		target int

		expectedID int
	}{
		{"valid authorId", 4, 4},
		{"invalid authorId", -1, -1},
	}

	for _, tc := range testcases {
		m := New(mockStore{})

		id, err := m.DeleteAuthor(tc.target)

		if err != nil {
			log.Print(err)
		}

		if id != tc.expectedID {
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

func (m mockStore) PutAuthor(author2 entities.Author) (int, error) {
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
