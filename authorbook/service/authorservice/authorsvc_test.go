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
		{"valid author", entities.Author{
			4, "nilotpal", "mrinal", "20/05/1990", "Dark horse"},
			entities.Author{4, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}},
		{"existing author", entities.Author{
			3, "nilotpal", "mrinal", "00/05/1990", "Dark horse"}, entities.Author{}},
		{"invalid firstname", entities.Author{
			3, " ", "mrinal", "20/00/1990", "Dark horse"}, entities.Author{}},
		{"invalid DOB", entities.Author{
			3, "nilotpal", "mrinal", "20/01/0", "Dark horse"}, entities.Author{}},
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
		{"valid author", entities.Author{
			4, "nilotpal", "mrinal", "20/05/1990", "Dark horse"},
			entities.Author{4, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}},
		{"existing author", entities.Author{
			3, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, entities.Author{}},
		{"invalid firstname", entities.Author{
			3, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, entities.Author{}},
		{"invalid DOB", entities.Author{
			3, "nilotpal", "mrinal", "20/00/1990", "Dark horse"}, entities.Author{}},
		{"valid author", entities.Author{
			5, "nilotpal", "mrinal", "20/05/1990", "Dark horse"},
			entities.Author{}},
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
