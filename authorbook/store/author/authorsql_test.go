package author

import (
	"errors"
	"log"
	"projects/GoLang-Interns-2022/authorbook/driver"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"testing"
)

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author

		err error
	}{
		{"valid author", entities.Author{
			0, "nilotpalx", "mrinal", "20/05/1990", "Dark horse"}, errors.New("success")},
		{"exiting author", entities.Author{
			0, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, errors.New("existing author")},
	}

	for _, tc := range testcases {
		DB := driver.Connection()
		authorStore := New(DB)

		id, err := authorStore.PostAuthor(tc.body)
		if err != nil {
			log.Print(err)
		}

		if id == -1 {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author

		err error
	}{
		{"valid author", entities.Author{
			4, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, errors.New("success")},
		{"exiting author", entities.Author{
			3, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}, errors.New("existing author")},
	}
	for _, tc := range testcases {
		DB := driver.Connection()
		authorStore := New(DB)

		id, err := authorStore.PostAuthor(tc.body)

		if id == -1 || tc.err != err {
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
		err error
	}{
		{"valid authorId", 4, nil},
		{"invalid authorId", 5, errors.New("invalid ID")},
	}

	for _, tc := range testcases {
		DB := driver.Connection()
		authorStore := New(DB)

		id, err := authorStore.DeleteAuthor(tc.target)

		if id == 0 && tc.err != err {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
