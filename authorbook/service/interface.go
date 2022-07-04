package service

import "projects/GoLang-Interns-2022/authorbook/entities"

type AuthorService interface {
	PostAuthor(entities.Author) (entities.Author, error)
	PutAuthor(entities.Author) (entities.Author, error)
	DeleteAuthor(int) (int, error)
}

type BookService interface {
	GetAllBook(string, string) []entities.Book
	GetBookByID(int) (entities.Book, error)
	PostBook(book entities.Book) (entities.Book, error)
	PutBook(book entities.Book, id int) (entities.Book, error)
	DeleteBook(int) (int, error)
}
