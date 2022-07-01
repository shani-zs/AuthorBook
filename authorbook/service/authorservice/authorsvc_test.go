package authorservice

import (
	"log"
	"net/http"

	"projects/GoLang-Interns-2022/authorbook/driver"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/store/author"
	"testing"
)

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author
	}{
		{"valid author", entities.Author{
			4, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}},
		{"exiting author", entities.Author{
			3, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}},
		{"invalid firstname", entities.Author{
			3, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}},
		{"invalid DOB", entities.Author{
			3, "nilotpal", "mrinal", "20/00/1990", "Dark horse"}},
	}

	for _, tc := range testcases {

		DB := driver.Connection()
		authorStore := author.New(DB)
		authorService := New(authorStore)

		id, err := authorService.PostAuthor(tc.body)

		if err != nil {
			log.Print(err)
		}

		if id == 0 {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author
	}{
		{"valid author", entities.Author{
			4, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}},
		{"exiting author", entities.Author{
			3, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}},
		{"invalid firstname", entities.Author{
			3, "nilotpal", "mrinal", "20/05/1990", "Dark horse"}},
		{"invalid DOB", entities.Author{
			3, "nilotpal", "mrinal", "20/00/1990", "Dark horse"}},
	}

	for _, tc := range testcases {

		DB := driver.Connection()
		authorStore := author.New(DB)
		authorService := New(authorStore)

		id, err := authorService.PostAuthor(tc.body)

		if err != nil {
			log.Print(err)
		}

		if id == 0 {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		//input
		desc   string
		target int
		//output
		expected int
	}{
		{"valid authorId", 4, http.StatusNoContent},
		{"invalid authorId", 5, http.StatusBadRequest},
	}

	for _, tc := range testcases {
		DB := driver.Connection()
		authorStore := author.New(DB)
		authorService := New(authorStore)

		id, err := authorService.DeleteAuthor(tc.target)

		if err != nil {
			log.Print(err)
		}

		if id == 0 {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
