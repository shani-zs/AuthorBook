package author

import (
	"database/sql"
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
	_, err := s.DB.Exec("insert into author(author_id,first_name,last_name,dob,pen_name)values(?,?,?,?,?)",
		author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName)
	if err != nil {
		log.Print(err)
		return -1, err
	}

	return author.AuthorID, nil
}

// PutAuthor : inserts an author if that does not exist and update author if exists
func (s Store) PutAuthor(author entities.Author, id int) (int, error) {
	res, err := s.DB.Exec("select count(author_id) from author where author_id=?", id)
	if err != nil {
		log.Print(err)
		return -1, err
	}

	rA, _ := res.RowsAffected()
	if rA > 0 {
		res, err := s.DB.Exec("update author set author_id=?,first_name=?,last_name=?,dob=?,pen_name=? where author_id=?",
			author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName, id)
		if err != nil {
			log.Print(err)
			return -1, err
		}

		id, _ := res.LastInsertId()
		author.AuthorID = int(id)

		return int(id), nil
	}

	return -1, err
}

// DeleteAuthor :  deletes an author
func (s Store) DeleteAuthor(id int) (int, error) {
	res, err := s.DB.Exec("delete from author where author_id=?", id)
	if err != nil {
		return -1, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return int(count), err
	}

	return int(count), nil
}
