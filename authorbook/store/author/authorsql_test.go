package author

import (
	"errors"
	"log"
	"projects/GoLang-Interns-2022/authorbook/entities"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestPost(t *testing.T) {
	testcases := []struct {
		desc string
		body entities.Author

		expectedErr  error
		RowAffected  int64
		LastInserted int64
	}{
		{desc: "valid authorhttp", body: entities.Author{
			AuthorID: 11, FirstName: "vinod", LastName: "pal", DOB: "20/05/1990", PenName: "Dh"},
			expectedErr: nil, RowAffected: 1, LastInserted: 11},
		{desc: "exiting authorhttp", body: entities.Author{
			AuthorID: 1, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"},
			expectedErr: errors.New("already exists"), RowAffected: 0, LastInserted: 0},
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("error during the opening of database:%v\n", err)
	}

	defer db.Close()

	for _, tc := range testcases {
		mock.ExpectExec("insert into authorhttp(author_id,first_name,last_name,dob,pen_name)values(?,?,?,?,?)").
			WithArgs(tc.body.AuthorID, tc.body.FirstName, tc.body.LastName,
				tc.body.DOB, tc.body.PenName).WillReturnResult(sqlmock.NewResult(tc.LastInserted, tc.RowAffected)).
			WillReturnError(tc.expectedErr)

		s := New(db)
		_, err = s.Post(tc.body)

		if err != tc.expectedErr {
			t.Errorf("failed for %s", tc.desc)
		}
	}
}
func TestPut(t *testing.T) {
	testcases := []struct {
		desc         string
		body         entities.Author
		id           int
		RowAffected  int64
		LastInserted int64

		expectedErr error
	}{
		{desc: "invalid authorhttp", body: entities.Author{
			AuthorID: 4, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"}, id: 20,
			RowAffected: 0, LastInserted: 0, expectedErr: errors.New("does not exist")},
		{desc: "exiting authorhttp", body: entities.Author{
			AuthorID: 3, FirstName: "nilotpal", LastName: "mrinal", DOB: "20/05/1990", PenName: "Dark horse"}, id: 4,
			RowAffected: 1, LastInserted: 0, expectedErr: nil},
	}

	for _, tc := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Print(err)
		}
		s := New(db)

		mock.ExpectExec("select count(author_id) from authorhttp where author_id=?").WithArgs(tc.id).
			WillReturnResult(sqlmock.NewResult(tc.LastInserted, tc.RowAffected)).WillReturnError(tc.expectedErr)

		mock.ExpectExec("update authorhttp set author_id=?,first_name=?,last_name=?,dob=?,pen_name=? where author_id=?").
			WithArgs(tc.body.AuthorID, tc.body.FirstName, tc.body.LastName, tc.body.DOB, tc.body.PenName, tc.id).
			WillReturnResult(sqlmock.NewResult(tc.LastInserted, tc.RowAffected)).WillReturnError(tc.expectedErr)

		_, err = s.Put(tc.body, tc.id)

		if err != tc.expectedErr {
			t.Errorf("failed for %v\n, expected: %v, got: %v", tc.desc, tc.expectedErr, err)
		}
		db.Close()
	}
}
func TestDelete(t *testing.T) {
	testcases := []struct {
		// input
		desc   string
		target int
		// output
		rowsAffected   int64
		lastInsertedID int64
		expectedErr    error
	}{
		{"valid authorId", 4, 1, 0, nil},
		{"invalid authorId", -1, 0, 0, errors.New("invalid authorID")},
	}

	for _, tc := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Print(err)
		}
		as := New(db)

		mock.ExpectExec("delete from author where author_id=?").WithArgs(tc.target).
			WillReturnResult(sqlmock.NewResult(tc.lastInsertedID, tc.rowsAffected)).WillReturnError(tc.expectedErr)
		_, err = as.Delete(tc.target)

		if err != tc.expectedErr {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
func TestIncludeAuthor(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Print(err)
	}
	var (
		author  = entities.Author{AuthorID: 1, FirstName: "shani", LastName: "kumar", DOB: "20/06/2000", PenName: "sk"}
		author1 = sqlmock.NewRows([]string{"author_id", "first_name", "last_name", "dob", "pen_name"}).AddRow(author.AuthorID,
			author.FirstName, author.LastName, author.DOB, author.PenName)
	)

	Testcases := []struct {
		desc     string
		targetID int

		expected    entities.Author
		expectedErr error
	}{
		{desc: "fetching book by id",
			targetID: 1, expected: entities.Author{AuthorID: 1, FirstName: "shani", LastName: "kumar", DOB: "20/06/2000", PenName: "sk"},
		},
		{"invalid id", -1, entities.Author{}, errors.New("invalid")},
	}

	for _, tc := range Testcases {
		bs := New(db)
		mock.ExpectQuery("SELECT * FROM author where author_id=?").WithArgs(tc.targetID).WillReturnRows(author1).WillReturnError(tc.expectedErr)
		a, err := bs.IncludeAuthor(tc.targetID)
		if err != nil {
			log.Print(err)
		}

		if a != tc.expected || err != tc.expectedErr {
			t.Errorf("failed for %v\n", tc.desc)
		}
	}
}
