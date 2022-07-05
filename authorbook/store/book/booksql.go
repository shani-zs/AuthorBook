package book

import (
	"database/sql"
	"fmt"
	"log"
	"projects/GoLang-Interns-2022/authorbook/entities"
)

type Store struct {
	DB *sql.DB
}

func New(db *sql.DB) Store {
	return Store{db}
}

func (bs Store) GetAllBook(title, includeAuthor string) []entities.Book {
	var books []entities.Book

	if title != "" {
		books = FetchingAllBooks(title, bs.DB)
		if includeAuthor == "true" {
			books = BooksWithAuthor(books, bs.DB)
		}
	} else {
		books = FetchingAllBooks("", bs.DB)
		if includeAuthor == "true" {
			books = BooksWithAuthor(books, bs.DB)
		}
	}

	return books
}

func (bs Store) GetBookByID(id int) entities.Book {
	var b entities.Book

	row := bs.DB.QueryRow("select * from book where id=?", id)

	err := row.Scan(&b.BookID, &b.AuthorID, &b.Title, &b.Publication, &b.PublishedDate)
	if err != nil {
		log.Print(err)
		return entities.Book{}
	}

	return b
}

func (bs Store) PostBook(book *entities.Book) (int, error) {
	result, err := bs.DB.Exec("insert into book(author_id,title,publication,published_date)values(?,?,?,?)",
		book.AuthorID, book.Title, book.Publication, book.PublishedDate)
	if err != nil {
		fmt.Println("hello")
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("shani")
		return -1, err
	}

	return int(id), nil
}

func (bs Store) PutBook(book *entities.Book, id int) (int, error) {
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

func (bs Store) DeleteBook(id int) (int, error) {
	result, err := bs.DB.Exec("delete from book where id=?", id)
	if err != nil {
		return -1, nil
	}

	count, err := result.RowsAffected()
	if err != nil {
		return -1, nil
	}

	return int(count), nil
}

func FetchingAllBooks(title string, db *sql.DB) []entities.Book {
	var (
		Rows *sql.Rows
		err  error
	)

	if title == "" {
		Rows, err = db.Query("SELECT * FROM book")
		if err != nil {
			log.Print(err)
		}
	} else {
		Rows, err = db.Query("SELECT * FROM book where title=?", title)
		if err != nil {
			log.Print(err)
		}
	}

	var bk []entities.Book

	for Rows.Next() {
		var b entities.Book

		err = Rows.Scan(&b.BookID, &b.AuthorID, &b.Title, &b.Publication, &b.PublishedDate)
		if err != nil {
			log.Print(err)
		}
		// b.Author = entities.Author{}
		bk = append(bk, b)
	}

	return bk
}

func FetchingAuthor(id int, db *sql.DB) (int, entities.Author) {
	Row := db.QueryRow("SELECT * FROM author where author_id=?", id)

	var author entities.Author

	if err := Row.Scan(&author.AuthorID, &author.FirstName, &author.LastName, &author.DOB, &author.PenName); err != nil {
		log.Printf("failed: %v\n", err)
		return 0, entities.Author{}
	}

	return author.AuthorID, author
}

func BooksWithAuthor(books []entities.Book, db *sql.DB) []entities.Book {
	for i := range books {
		_, a := FetchingAuthor(books[i].AuthorID, db)
		books[i].Author = a
	}

	return books
}
