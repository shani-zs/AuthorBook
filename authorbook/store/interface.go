package store

import "projects/GoLang-Interns-2022/authorbook/entities"

type AuthorStorer interface {
	PostAuthor(entities.Author) (int, error)
	PutAuthor(entities.Author) (int, error)
	DeleteAuthor(int) (int, error)
}

type BookStorer interface {
	GetAllBook(string2 string, string3 string) []entities.Book
	GetBookByID(int) entities.Book
	PostBook(book entities.Book) (int, error)
	PutBook(book entities.Book, id int) (int, error)
	DeleteBook(int) (int, error)
}
