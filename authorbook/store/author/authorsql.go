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

// Post : insert an author
func (s Store) Post(author entities.Author) (int, error) {
	res, err := s.DB.Exec("insert into author(first_name,last_name,dob,pen_name)values(?,?,?,?)",
		author.FirstName, author.LastName, author.DOB, author.PenName)
	if err != nil {
		log.Print(err)
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

// Put : inserts an author if that does not exist and update author if exists
func (s Store) Put(author entities.Author, id int) (int, error) {
	_, err := s.DB.Exec("update author set author_id=?,first_name=?,last_name=?,dob=?,pen_name=? where author_id=?",
		author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName, id)
	if err != nil {
		log.Print(err)
		return -1, err
	}

	return author.AuthorID, nil
}

// Delete :  deletes an author
func (s Store) Delete(id int) (int, error) {
	res, err := s.DB.Exec("delete from author where author_id=?", id)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return int(count), err
	}

	return int(count), nil
}

func (s Store) IncludeAuthor(id int) (entities.Author, error) {
	var author entities.Author

	Row := s.DB.QueryRow("SELECT * FROM author where author_id=?", id)

	if err := Row.Scan(&author.AuthorID, &author.FirstName, &author.LastName, &author.DOB, &author.PenName); err != nil {
		log.Printf("failed: %v\n", err)
		return entities.Author{}, err
	}

	return author, nil
}
