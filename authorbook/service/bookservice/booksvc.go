package bookservice

import (
	"log"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"projects/GoLang-Interns-2022/authorbook/store"
	"strings"
)

type BookService struct {
	bookService store.BookStorer
}

func New(s store.BookStorer) BookService {
	return BookService{s}
}

func (bs BookService) GetAllBook(title, includeAuthor string) ([]entities.Book, error) {
	var (
		books []entities.Book
		err   error
	)

	if title != "" {
		books, err = bs.bookService.GetBooksByTitle(title)
		if err != nil {
			log.Print(err)
			return []entities.Book{}, err
		}
	}

	books, err = bs.bookService.GetAllBook()
	if err != nil {
		log.Print(err)
		return []entities.Book{}, err
	}

	if includeAuthor == "true" {
		for i := range books {
			author, err := bs.bookService.IncludeAuthor(books[i].AuthorID)
			if err != nil {
				log.Print(err)
				return []entities.Book{}, err
			}

			books[i].Author = author
		}
	}

	return books, nil
}

func (bs BookService) GetBookByID(id int) (entities.Book, error) {
	if id <= 0 {
		return entities.Book{}, nil
	}

	book, err := bs.bookService.GetBookByID(id)
	if err != nil {
		log.Print(err)
		return entities.Book{}, err
	}

	return book, nil
}

func (bs BookService) Post(book *entities.Book) (entities.Book, error) {
	if book.Title == "" || book.AuthorID < 0 || checkPublication(book.Publication) {
		return entities.Book{}, nil
	}

	id, err := bs.bookService.Post(book)
	if err != nil || id == -1 {
		return entities.Book{}, err
	}

	book.BookID = id

	return *book, nil
}

func (bs BookService) Put(book *entities.Book, id int) (entities.Book, error) {
	if book.Title == "" || book.AuthorID <= 0 || checkPublication(book.Publication) {
		return entities.Book{}, nil
	}

	i, err := bs.bookService.Put(book, id)
	if err != nil || id <= -1 {
		return entities.Book{}, err
	}

	book.BookID = i

	return *book, nil
}

func (bs BookService) Delete(id int) (int, error) {
	if id < 0 {
		return -1, nil
	}

	i, err := bs.bookService.Delete(id)
	if err != nil || i == -1 {
		return -1, err
	}

	return i, nil
}

func checkPublication(publication string) bool {
	_ = strings.ToLower(publication)

	return !(publication == "penguin" || publication == "scholastic" || publication == "arihant")
}
