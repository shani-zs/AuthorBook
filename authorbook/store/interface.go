package store

import "projects/GoLang-Interns-2022/authorbook/entities"

type AuthorStorer interface {
	PostAuthor(entities.Author) (int64, error)
	PutAuthor(entities.Author) (int64, error)
	DeleteAuthor(int) (int64, error)
}

type BookStorer interface {
	GetAllBook(book entities.Book) []entities.Book
	GetBookByID(book entities.Book) entities.Book
	PostBook(book entities.Book) (int, error)
	PutBook(book entities.Book) (int, error)
	DeleteBook(book entities.Book) (int, error)
}
