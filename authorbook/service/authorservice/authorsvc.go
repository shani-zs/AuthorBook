package authorservice

import (
	"errors"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/store"
	"strconv"
	"strings"
)

type AuthorService struct {
	datastore store.AuthorStorer
}

// New : factory function
func New(s store.AuthorStorer) AuthorService {
	return AuthorService{s}
}

// Post : checks the author before posting
func (s AuthorService) Post(a entities.Author) (entities.Author, error) {
	if a.FirstName == "" || !checkDob(a.DOB) {
		return entities.Author{}, errors.New("invalid constraints")
	}

	id, err := s.datastore.Post(a)
	if err != nil || id <= 0 {
		return entities.Author{}, err
	}

	a.AuthorID = id

	return a, nil
}

// Put : checks the author before updating
func (s AuthorService) Put(a entities.Author, id int) (entities.Author, error) {
	if a.FirstName == "" || !checkDob(a.DOB) {
		return entities.Author{}, nil
	}

	existAuthor, err := s.datastore.IncludeAuthor(id)
	if err != nil || existAuthor.AuthorID != id {
		return entities.Author{}, errors.New("dose not exist")
	}

	i, err := s.datastore.Put(a, id)
	if err != nil || i < 0 {
		return entities.Author{}, err
	}

	a.AuthorID = i

	return a, nil
}

// Delete : Deletes the author at particular id
func (s AuthorService) Delete(id int) (int, error) {
	if id < 0 {
		return 0, errors.New("invalid id")
	}

	countRows, err := s.datastore.Delete(id)
	if err != nil || countRows <= 0 {
		return 0, errors.New("does not exist")
	}

	return countRows, nil
}

// checkDob : validates the DOB
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
