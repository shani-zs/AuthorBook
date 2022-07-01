package author

import (
	"database/sql"
	"errors"
	"net/http"
	"projects/GoLang-Interns-2022/authorbook/entities"
)

type AuthorStore struct {
	DB *sql.DB
}

func New(DB *sql.DB) AuthorStore {
	return AuthorStore{DB}
}

// PostAuthor : insert an author
func (s AuthorStore) PostAuthor(author entities.Author) (int64, error) {

	var a entities.Author

	Row := s.DB.QueryRow("select * from author where author_id=?", author.AuthorID)

	err := Row.Scan(&a.AuthorID)

	if a.AuthorID == author.AuthorID {
		return 0, errors.New("already exits")
	}

	res, err := s.DB.Exec("insert into author(author_id,first_name,last_name,dob,pen_name)values(?,?,?,?,?)",
		author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName)
	if err != nil {
		return http.StatusBadRequest, err
	}

	id, _ := res.LastInsertId()
	author.AuthorID = int(id)

	return id, nil
}

// PutAuthor : inserts an author if that does not exist and update author if exists
func (s AuthorStore) PutAuthor(author entities.Author) (int64, error) {
	var a entities.Author

	Row := s.DB.QueryRow("select * from author where author_id=?", author.AuthorID)

	err := Row.Scan(&a.AuthorID)
	if err != nil {
		res, _ := s.DB.Exec("insert into author(author_id,first_name,last_name,DOB,pen_name)values(?,?,?,?,?)",
			author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName)
		id, _ := res.LastInsertId()

		return id, nil
	} else {
		res, _ := s.DB.Exec("update author set author_id=?,first_name=?,last_name=?,dob=?,pen_name=? where author_id=?",
			author.AuthorID, author.FirstName, author.LastName, author.DOB, author.DOB, a.AuthorID)
		id, _ := res.LastInsertId()
		author.AuthorID = int(id)

		return id, nil
	}

}

// DeleteAuthor :  deletes an author
func (s AuthorStore) DeleteAuthor(id int) (int64, error) {
	res, _ := s.DB.Exec("delete from author where author_id=?", id)
	count, err := res.RowsAffected()
	if err != nil {
		return count, err
	}

	return count, nil
}
