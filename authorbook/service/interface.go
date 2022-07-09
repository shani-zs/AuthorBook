package service

import (
	"projects/GoLang-Interns-2022/authorbook/entities"
)

type AuthorService interface {
	Post(author entities.Author) (entities.Author, error)
	Put(author entities.Author, id int) (entities.Author, error)
	Delete(id int) (int, error)
}

type BookService interface {
	GetAllBook(title string, includeAuthor string) ([]entities.Book, error)
	GetBookByID(id int) (entities.Book, error)
	Post(book *entities.Book) (entities.Book, error)
	Put(book *entities.Book, id int) (entities.Book, error)
	Delete(id int) (int, error)
}
