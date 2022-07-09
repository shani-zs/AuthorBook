package book

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"projects/GoLang-Interns-2022/authorbook/entities"
)

type Store struct {
	DB *sql.DB
}

func New(db *sql.DB) Store {
	return Store{db}
}

func (bs Store) GetAllBook(ctx context.Context) ([]entities.Book, error) {
	var (
		books []entities.Book
		Rows  *sql.Rows
		err   error
	)

	Rows, err = bs.DB.Query("SELECT * FROM book")
	if err != nil {
		log.Print(err)
		return []entities.Book{}, err
	}
	defer Rows.Close()

	for Rows.Next() {
		var book entities.Book

		err = Rows.Scan(&book.BookID, &book.AuthorID, &book.Title, &book.Publication, &book.PublishedDate)
		if err != nil {
			log.Print(err)
			return []entities.Book{}, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (bs Store) GetBooksByTitle(ctx context.Context, title string) ([]entities.Book, error) {
	var (
		Rows *sql.Rows
		err  error
	)

	Rows, err = bs.DB.Query("SELECT * FROM book WHERE title=?", title)
	if err != nil {
		log.Print(err)
		return []entities.Book{}, err
	}

	var books []entities.Book

	for Rows.Next() {
		var book entities.Book

		err = Rows.Scan(&book.BookID, &book.AuthorID, &book.Title, &book.Publication, &book.PublishedDate)
		if err != nil {
			log.Print(err)
			return []entities.Book{}, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (bs Store) IncludeAuthor(ctx context.Context, id int) (entities.Author, error) {

	Row := bs.DB.QueryRow("SELECT * FROM authorhttp where author_id=?", id)

	var author entities.Author

	if err := Row.Scan(&author.AuthorID, &author.FirstName, &author.LastName, &author.DOB, &author.PenName); err != nil {
		log.Printf("failed: %v\n", err)
		return entities.Author{}, err
	}

	return author, nil
}

func (bs Store) GetBookByID(ctx context.Context, id int) (entities.Book, error) {
	var b entities.Book

	row := bs.DB.QueryRow("select * from book where id=?", id)

	err := row.Scan(&b.BookID, &b.AuthorID, &b.Title, &b.Publication, &b.PublishedDate)
	if err != nil {
		log.Print(err)
		return entities.Book{}, err
	}

	return b, nil
}

func (bs Store) Post(ctx context.Context, book *entities.Book) (int, error) {
	// checking the authorhttp existing in the table table or not
	_, err := bs.DB.Exec("select count(author_id) from authorhttp where author_id=?", book.AuthorID)
	if err != nil {
		return 0, err
	}

	result, err := bs.DB.Exec("insert into book(author_id,title,publication,published_date)values(?,?,?,?)",
		book.AuthorID, book.Title, book.Publication, book.PublishedDate)
	if err != nil {
		log.Print(err)
		return -1, errors.New("already existing")
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Print(err)
		return -1, err
	}

	return int(id), nil
}

func (bs Store) Put(ctx context.Context, book *entities.Book, id int) (int, error) {
	var b entities.Book

	row := bs.DB.QueryRow("select * from book where id=?", id)

	err := row.Scan(&b.BookID, &b.AuthorID, &b.Title, &b.Publication, &b.PublishedDate)
	if err == nil {
		// update
		_, _ = bs.DB.Exec("update book set id=?,author_id=?,title=?,publication=?,published_date=? where id=?",
			book.BookID, book.AuthorID, book.Title, book.Publication, book.PublishedDate, id)
		return book.BookID, nil
	}

	result, err := bs.DB.Exec("insert into book(author_id,title,publication,published_date)values(?,?,?,?)",
		book.AuthorID, book.Title, book.Publication, book.PublishedDate)
	if err != nil {
		return -1, err
	}

	i, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return int(i), nil
}

func (bs Store) Delete(ctx context.Context, id int) (int, error) {
	result, err := bs.DB.Exec("delete from book where id=?", id)
	if err != nil {
		return -1, errors.New("invalid id")
	}

	count, err := result.RowsAffected()
	if err != nil {
		return -1, nil
	}

	return int(count), nil
}
