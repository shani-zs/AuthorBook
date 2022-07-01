package bookservice

import (
	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/store"

	"projects/GoLang-Interns-2022/authorbook/store/book"
)

type BookService struct {
	bookService store.AuthorStorer
}

func New(s book.BookStore) BookService {
	return BookService{s}
}

func (bs BookService) GetAllBook(title string, includeAuthor string) []entities.Book {

}

func (bs BookService) GetBookByID(id int) (entities.Book, error) {

}
