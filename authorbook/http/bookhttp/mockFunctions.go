package bookhttp

import (
	"projects/GoLang-Interns-2022/authorbook/entities"
)

type mockService struct{}

func (m mockService) GetAllBook(s, s2 string) []entities.Book {
	if s2 == "" && s == "" {
		return []entities.Book{{1,
			1, "book one", "scholastic", "20/06/2018", entities.Author{}},
			{2, 1, "book two", "penguin", "20/08/2018", entities.Author{}}}
	}
	if s == "book+two" && s2 == "true" {
		return []entities.Book{
			{2, 1, "book two", "penguin", "20/08/2018", entities.Author{1, "shani",
				"kumar", "30/04/2001", "sk"}}}
	}

	return []entities.Book{}
}

func (m mockService) GetBookByID(i int) (entities.Book, error) {
	// TODO implement me
	panic("implement me")
}

func (m mockService) PostBook(book *entities.Book) (entities.Book, error) {
	// TODO implement me
	panic("implement me")
}

func (m mockService) PutBook(book *entities.Book, id int) (entities.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m mockService) DeleteBook(i int) (int, error) {
	//TODO implement me
	panic("implement me")
}
