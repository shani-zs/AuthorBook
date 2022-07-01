package book

import (
	"database/sql"
	"projects/GoLang-Interns-2022/authorbook/entities"
)

type BookStore struct {
	DB *sql.DB
}

func New(DB *sql.DB) BookStore {
	return BookStore{DB}
}

func (bs BookStore) GetAllBook(title string, includeAuthor string) []entities.Book {
	var b []entities.Book

	return b
}

func (bs BookStore) GetBookByID(id int) entities.Book {
	var b entities.Book

	return b
}

func (bs BookStore) PostBook(book entities.Book) (int, error) {

}

func (bs BookStore) PutBook(book entities.Book) (int, error) {

}

func (bs BookStore) DeleteBook(id int) (int, error) {

}
