package authorservice

import (
	"errors"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/store"
	"projects/GoLang-Interns-2022/authorbook/store/author"
	"strconv"
	"strings"
)

type AuthorService struct {
	datastore store.AuthorStorer
}

// New : factory function
func New(s author.AuthorStore) AuthorService {
	return AuthorService{s}
}

func (s AuthorService) PostAuthor(a entities.Author) (int64, error) {

	if a.FirstName == "" || !checkDob(a.DOB) {
		return 0, errors.New("not valid constraints")
	}

	id, err := s.datastore.PostAuthor(a)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func checkDob(dob string) bool {
	Dob := strings.Split(dob, "/")
	day, _ := strconv.Atoi(Dob[0])
	month, _ := strconv.Atoi(Dob[1])
	year, _ := strconv.Atoi(Dob[2])

	switch {
	case day <= 0 || day > 31:
		return false
	case month <= 0 || month > 12:
		return false
	case year <= 0:
		return false
	}

	return true
}

// PutAuthor : business logic of putathor
func (s AuthorService) PutAuthor(a entities.Author) (int64, error) {
	if a.FirstName == "" || !checkDob(a.DOB) {
		return 0, errors.New("not valid constraints")
	}

	id, err := s.datastore.PostAuthor(a)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// DeleteAuthor : Deletes the author at particular id

func (s AuthorService) DeleteAuthor(id int) (int64, error) {
	if id < 0 {
		return 0, errors.New("not valid id")
	}

	count, err := s.datastore.DeleteAuthor(id)
	if err != nil {
		return 0, err
	}

	return count, nil
}
