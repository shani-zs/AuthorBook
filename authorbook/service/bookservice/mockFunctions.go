package bookservice

import "projects/GoLang-Interns-2022/authorbook/entities"

type mockStore struct{}

func (m mockStore) GetAllBook(string2, string3 string) []entities.Book {
	i := 1

	if string2 == "" && string3 == "" {
		return []entities.Book{{BookID: i,
			AuthorID: 1, Title: "book one", Publication: "scholastic", PublishedDate: "20/06/2018",
			Author: entities.Author{}}, {BookID: i + 1, AuthorID: 1, Title: "book two", Publication: "penguin",
			PublishedDate: "20/08/2018", Author: entities.Author{}}}
	}

	if string2 == "book+two" && string3 == "true" {
		i := 2

		return []entities.Book{
			{BookID: i, AuthorID: 1, Title: "book two", Publication: "penguin", PublishedDate: "20/08/2018",
				Author: entities.Author{AuthorID: 1, FirstName: "shani", LastName: "kumar", DOB: "30/04/2001",
					PenName: "sk"}}}
	}

	return []entities.Book{}
}

func (m mockStore) GetBookByID(i int) entities.Book {
	if i == 1 {
		return entities.Book{BookID: 1, AuthorID: 1, Title: "book two", Publication: "penguin",
			PublishedDate: "20/08/2018", Author: entities.Author{AuthorID: 1, FirstName: "shani", LastName: "kumar",
				DOB: "30/04/2001", PenName: "sk"}}
	}

	return entities.Book{}
}

func (m mockStore) PostBook(book *entities.Book) (int, error) {
	i := 4
	if book.BookID == i {
		return book.BookID, nil
	}

	return -1, nil
}

func (m mockStore) PutBook(book *entities.Book, id int) (int, error) {
	i := 4
	if book.BookID == i {
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
