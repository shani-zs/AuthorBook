package store

import (
	"projects/GoLang-Interns-2022/authorbook/entities"
)

type AuthorStorer interface {
	Post(author entities.Author) (int, error)
	Put(author entities.Author, id int) (int, error)
	Delete(id int) (int, error)
	IncludeAuthor(id int) (entities.Author, error)
}

type BookStorer interface {
	GetAllBook() ([]entities.Book, error)
	GetBooksByTitle(title string) ([]entities.Book, error)

	GetBookByID(id int) (entities.Book, error)
	Post(book *entities.Book) (int, error)
	Put(book *entities.Book, id int) (int, error)
	Delete(id int) (int, error)
}
