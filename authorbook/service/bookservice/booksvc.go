package bookservice

import (
	"errors"
	"log"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/store"
	"strings"
)

type BookService struct {
	bookService   store.BookStorer
	authorService store.AuthorStorer
}

func New(bs store.BookStorer, as store.AuthorStorer) BookService {
	return BookService{bs, as}
}

func (b BookService) GetAllBook(title, includeAuthor string) ([]entities.Book, error) {
	var (
		books []entities.Book
		err   error
	)

	if title != "" {
		books, err = b.bookService.GetBooksByTitle(title)
		if err != nil {
			log.Print(err)
			return []entities.Book{}, err
		}
	} else {
		books, err = b.bookService.GetAllBook()
		if err != nil {
			log.Print(err)
			return []entities.Book{}, err
		}
	}
	if includeAuthor == "true" {
		for i := range books {
			author, err := b.authorService.IncludeAuthor(books[i].AuthorID)
			if err != nil {
				log.Print(err)
				return []entities.Book{}, err
			}

			books[i].Author = author
		}
	}

	return books, nil
}

func (b BookService) GetBookByID(id int) (entities.Book, error) {
	if id <= 0 {
		return entities.Book{}, errors.New("invalid id")
	}

	book, err := b.bookService.GetBookByID(id)
	if err != nil {
		log.Print(err)
		return entities.Book{}, err
	}

	return book, nil
}

func (b BookService) Post(book *entities.Book) (entities.Book, error) {
	if book.Title == "" || book.AuthorID < 0 || checkPublication(book.Publication) {
		return entities.Book{}, errors.New("invalid constraints")
	}

	id, err := b.bookService.Post(book)
	if err != nil || id == -1 {
		return entities.Book{}, err
	}

	book.BookID = id

	return *book, nil
}

func (b BookService) Put(book *entities.Book, id int) (entities.Book, error) {
	if book.Title == "" || book.AuthorID <= 0 || checkPublication(book.Publication) {
		return entities.Book{}, nil
	}

	i, err := b.bookService.Put(book, id)
	if err != nil || id <= -1 {
		return entities.Book{}, err
	}

	book.BookID = i

	return *book, nil
}

func (b BookService) Delete(id int) (int, error) {
	if id < 0 {
		return -1, nil
	}

	i, err := b.bookService.Delete(id)
	if err != nil || i == -1 {
		return -1, err
	}

	return i, nil
}

func checkPublication(publication string) bool {
	_ = strings.ToLower(publication)

	return !(publication == "penguin" || publication == "scholastic" || publication == "arihant")
}
