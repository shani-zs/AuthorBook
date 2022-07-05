package author

import (
	"projects/GoLang-Interns-2022/authorbook/driver"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"testing"
)

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author

		expectedID int
	}{
		{desc: "valid author", body: entities.Author{
			AuthorID: 5, FirstName: "nilotpalx", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"},
			expectedID: 5},
		{desc: "exiting author", body: entities.Author{
			AuthorID: 1, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"},
			expectedID: -1},
	}

	for _, tc := range testcases {
		DB := driver.Connection()
		authorStore := New(DB)

		id, _ := authorStore.PostAuthor(tc.body)

		if id != tc.expectedID {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author

		expected int
	}{
		{desc: "valid author", body: entities.Author{
			AuthorID: 4, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"},
			expected: -1},
		{desc: "exiting author", body: entities.Author{
			AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"},
			expected: -1},
	}
	for _, tc := range testcases {
		DB := driver.Connection()
		authorStore := New(DB)

		id, _ := authorStore.PostAuthor(tc.body)

		if id != tc.expected {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		// input
		desc   string
		target int
		// output
		expected int
	}{
		{"valid authorId", 4, 1},
		{"invalid authorId", -1, 0},
	}

	for _, tc := range testcases {
		DB := driver.Connection()
		authorStore := New(DB)

		count, _ := authorStore.DeleteAuthor(tc.target)

		if count != tc.expected {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
