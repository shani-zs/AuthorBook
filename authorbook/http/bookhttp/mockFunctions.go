package bookhttp

import (
	"projects/GoLang-Interns-2022/authorbook/entities"
)

type mockService struct{}

func (m mockService) GetAllBook(s string, s2 string) []entities.Book {
	//TODO implement me
	panic("implement me")
}

func (m mockService) GetBookByID(i int) (entities.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockService) PostBook(book entities.Book) (entities.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockService) PutBook(book entities.Book) (entities.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockService) DeleteBook(i int) (int, error) {
	//TODO implement me
	panic("implement me")
}
