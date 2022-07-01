package service

import "projects/GoLang-Interns-2022/authorbook/entities"

type AuthorService interface {
	PostAuthor(entities.Author) (int64, error)
	PutAuthor(entities.Author) (int64, error)
	DeleteAuthor(int) (int64, error)
}

type BookService interface {
	GetAllBook(book entities.Book) []entities.Book
	GetBookByID(book entities.Book) (entities.Book, error)
	PostBook(book entities.Book) (int, error)
	PutBook(book entities.Book) (int, error)
	DeleteBook(book entities.Book) (int, error)
}
