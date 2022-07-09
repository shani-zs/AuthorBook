package author

import (
	"context"
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
func (s Store) Post(ctx context.Context, author entities.Author) (int, error) {
	_, err := s.DB.ExecContext(ctx, "insert into authorhttp(author_id,first_name,last_name,dob,pen_name)values(?,?,?,?,?)",
		author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName)
	if err != nil {
		log.Print(err)
		return -1, err
	}

	return author.AuthorID, nil
}

// Put : inserts an author if that does not exist and update author if exists
func (s Store) Put(ctx context.Context, author entities.Author, id int) (int, error) {
	res, err := s.DB.ExecContext(ctx, "select count(author_id) from authorhttp where author_id=?", id)
	if err != nil {
		log.Print(err)
		return -1, err
	}

	rA, _ := res.RowsAffected()
	if rA > 0 {
		res, err := s.DB.ExecContext(ctx, "update authorhttp set author_id=?,first_name=?,last_name=?,dob=?,pen_name=? where author_id=?",
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

// Delete :  deletes an authorhttp
func (s Store) Delete(ctx context.Context, id int) (int, error) {
	res, err := s.DB.Exec("delete from authorhttp where author_id=?", id)
	if err != nil {
		return 0, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return int(count), err
	}

	return int(count), nil
}
