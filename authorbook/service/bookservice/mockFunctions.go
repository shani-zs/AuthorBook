package bookservice

import "projects/GoLang-Interns-2022/authorbook/entities"

type mockStore struct{}

func (m mockStore) GetAllBook(string2, string3 string) []entities.Book {
	if string2 == "" && string3 == "" {
		return []entities.Book{{1,
			1, "book one", "scholastic", "20/06/2018", entities.Author{}},
			{2, 1, "book two", "penguin", "20/08/2018", entities.Author{}}}
	}
	if string2 == "book+two" && string3 == "true" {
		return []entities.Book{
			{2, 1, "book two", "penguin", "20/08/2018", entities.Author{1, "shani",
				"kumar", "30/04/2001", "sk"}}}
	}

	return []entities.Book{}
}

func (m mockStore) GetBookByID(i int) entities.Book {
	if i == 1 {
		return entities.Book{1, 1, "book two", "penguin",
			"20/08/2018", entities.Author{1, "shani", "kumar", "30/04/2001", "sk"}}
	}
	return entities.Book{}
}

func (m mockStore) PostBook(book *entities.Book) (int, error) {
	if book.BookID == 4 {
		return book.BookID, nil
	}
	return -1, nil
}

func (m mockStore) PutBook(book *entities.Book, id int) (int, error) {
	if book.BookID == 4 {
		return book.BookID, nil
	}
	return -1, nil
}

func (m mockStore) DeleteBook(i int) (int, error) {
	if i < 0 {
		return -1, nil
	}

	return i, nil
}
