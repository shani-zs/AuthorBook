package author

import (
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

// PostAuthor : insert an author
func (s Store) PostAuthor(author entities.Author) (int, error) {
	var a entities.Author

	Row := s.DB.QueryRow("select * from author where author_id=?", author.AuthorID)

	err := Row.Scan(&a.AuthorID)
	if err != nil {
		log.Print(err)
	}

	if a.AuthorID == author.AuthorID {
		return -1, errors.New("already exits")
	}

	_, err = s.DB.Exec("insert into author(author_id,first_name,last_name,dob,pen_name)values(?,?,?,?,?)",
		author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName)
	if err != nil {
		return -1, err
	}

	return author.AuthorID, nil
}

// PutAuthor : inserts an author if that does not exist and update author if exists
func (s Store) PutAuthor(author entities.Author) (int, error) {
	var a entities.Author

	Row := s.DB.QueryRow("select * from author where author_id=?", author.AuthorID)

	err := Row.Scan(&a.AuthorID, &a.FirstName, &a.LastName, &a.DOB, &a.PenName)
	if err != nil {
		res, _ := s.DB.Exec("insert into author(author_id,first_name,last_name,DOB,pen_name)values(?,?,?,?,?)",
			author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName)
		id, _ := res.LastInsertId()

		return int(id), nil
	}

	res, _ := s.DB.Exec("update author set author_id=?,first_name=?,last_name=?,dob=?,pen_name=? where author_id=?",
		author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName, a.AuthorID)
	id, _ := res.LastInsertId()
	author.AuthorID = int(id)

	return int(id), nil
}

// DeleteAuthor :  deletes an author
func (s Store) DeleteAuthor(id int) (int, error) {
	res, _ := s.DB.Exec("delete from author where author_id=?", id)

	count, err := res.RowsAffected()
	if err != nil {
		return int(count), err
	}

	return int(count), nil
}
