package bookservice

import "projects/GoLang-Interns-2022/authorbook/entities"

type mockStore struct{}

func (m mockStore) GetAllBook(string2 string, string3 string) []entities.Book {
	//TODO implement me
	panic("implement me")
}

func (m mockStore) GetBookByID(i int) entities.Book {
	//TODO implement me
	panic("implement me")
}

func (m mockStore) PostBook(book entities.Book) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockStore) PutBook(book entities.Book, id int) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockStore) DeleteBook(i int) (int, error) {
	//TODO implement me
	panic("implement me")
}
