package bookservice

import (
	"context"
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

// New : factory function
func New(bs store.BookStorer, as store.AuthorStorer) BookService {
	return BookService{bs, as}
}

// GetAllBook : implements the logic of getting all book
func (b BookService) GetAllBook(ctx context.Context, title, includeAuthor string) ([]entities.Book, error) {
	var (
		books []entities.Book
		err   error
	)

	if title != "" {
		books, err = b.bookService.GetBooksByTitle(ctx, title)
		if err != nil {
			log.Print(err)
			return []entities.Book{}, err
		}
	} else {
		books, err = b.bookService.GetAllBook(ctx)
		if err != nil {
			log.Print(err)
			return []entities.Book{}, err
		}
	}

	if includeAuthor == "true" {
		for i := range books {
			author, err := b.authorService.IncludeAuthor(ctx, books[i].AuthorID)
			if err != nil {
				log.Print(err)
				return []entities.Book{}, err
			}

			books[i].Author = author
		}
	}

	return books, nil
}

// GetBookByID : implements the logic of getting a single by
func (b BookService) GetBookByID(ctx context.Context, id int) (entities.Book, error) {
	if id <= 0 {
		return entities.Book{}, errors.New("invalid id")
	}

	book, err := b.bookService.GetBookByID(ctx, id)
	if err != nil {
		log.Print(err)
		return entities.Book{}, err
	}

	return book, nil
}

// Post : checks the book before posting
func (b BookService) Post(ctx context.Context, book *entities.Book) (entities.Book, error) {
	if book.Title == "" || book.AuthorID < 0 || checkPublication(book.Publication) {
		return entities.Book{}, errors.New("invalid constraints")
	}

	existAuthor, err := b.authorService.IncludeAuthor(ctx, book.AuthorID)
	if err != nil {
		return entities.Book{}, err
	}

	id, err := b.bookService.Post(ctx, book)
	if err != nil || id == -1 {
		return entities.Book{}, errors.New("database issue")
	}

	book.Author = existAuthor
	book.BookID = id

	return *book, nil
}

// Put :  checks the book before updating
func (b BookService) Put(ctx context.Context, book *entities.Book, id int) (entities.Book, error) {
	if book.Title == "" || book.AuthorID <= 0 || checkPublication(book.Publication) {
		return entities.Book{}, errors.New("invalid constraints")
	}

	count, err := b.bookService.Put(ctx, book, id)
	if err != nil || count <= 0 {
		return entities.Book{}, errors.New("does not exist")
	}

	return *book, nil
}

// Delete : checks before deleting a book
func (b BookService) Delete(ctx context.Context, id int) (int, error) {
	if id < 0 {
		return -1, nil
	}

	count, err := b.bookService.Delete(ctx, id)
	if err != nil || count <= 0 {
		return -1, errors.New("does not exist")
	}

	return count, nil
}

// checkPublication : validates publication
func checkPublication(publication string) bool {
	_ = strings.ToLower(publication)

	return !(publication == "penguin" || publication == "scholastic" || publication == "arihant")
}
