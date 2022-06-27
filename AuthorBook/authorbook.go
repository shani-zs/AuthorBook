package AuthorBook

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Author struct {
	AuthorId  int    `json:"authorId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DOB       string `json:"DOB"`
	PenName   string `json:"penName"`
}

type Book struct {
	BookId        string  `json:"bookId"`
	AuthorId      int     `json:"authorId"`
	Title         string  `json:"title"`
	Publication   string  `json:"publication"`
	PublishedDate string  `json:"publishedDate"`
	Author        *Author `json:"author"`
}

// Connection : makes the connection to the database
func Connection() *sql.DB {
	Db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/AuthorBook")
	if err != nil {
		log.Fatal("failed to connect with database:\n", err)
	}

	pingErr := Db.Ping()
	if pingErr != nil {
		log.Fatal("failed to ping", pingErr)
	}

	return Db
}

// FetchingAllBooks : fetches all books from the database
func FetchingAllBooks() []Book {
	Db := Connection()
	defer Db.Close()

	Rows, err := Db.Query("SELECT * FROM book")

	if err != nil {
		fmt.Errorf("%v\n", err)
	}
	defer Rows.Close()

	var bk []Book

	for Rows.Next() {
		var b Book
		err := Rows.Scan(&b.BookId, &b.AuthorId, &b.Title, &b.Publication, &b.PublishedDate)
		if err != nil {
			fmt.Errorf("%v\n", err)
		}

		_, author := FetchingAuthor(b.AuthorId)
		b.Author = &author
		bk = append(bk, b)
	}

	return bk
}

// FetchingAuthor : gets the author detail from the database
func FetchingAuthor(id int) (int, Author) {
	Db := Connection()
	defer Db.Close()

	Row := Db.QueryRow("SELECT * FROM author where authorId=?", id)
	var author Author
	if err := Row.Scan(&author.AuthorId, &author.FirstName, &author.LastName, &author.DOB, &author.PenName); err != nil {
		fmt.Errorf("failed: %v\n", err)
	}
	return author.AuthorId, author
}

//GetAllBook : returns all books to the client
func GetAllBook(w http.ResponseWriter, req *http.Request) {

	bk := FetchingAllBooks()

	mbk, err := json.Marshal(bk)
	if err != nil {
		fmt.Errorf("%v\n", err)
	}
	bytes.NewBuffer(mbk)

	_, err = w.Write(mbk)
	if err != nil {
		fmt.Errorf("%v", err)
	}

}

// GetBookById : returns a single book
func GetBookById(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	if req.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	strings.ToLower(params["id"])
	id, _ := strconv.Atoi(params["id"])
	if id <= 0 {
		fmt.Errorf("invalid id")
		w.WriteHeader(http.StatusBadRequest)
	}

	Db := Connection()

	row := Db.QueryRow("select * from book where bookId=?", params["id"])
	var b Book
	if err := row.Scan(&b.BookId, &b.AuthorId, &b.Title, &b.Publication, &b.PublishedDate); err != nil {
		fmt.Errorf("failed,%v\n", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, author := FetchingAuthor(b.AuthorId)
	b.Author = &author

	_, err := json.Marshal(b)
	if err != nil {
		fmt.Errorf("failed,%v\n", err)
		return
	}

	//w.Write(data)
	w.WriteHeader(http.StatusOK)

}

// checkDob : Validate the DOB of the author
func checkDob(DOB string) bool {

	dob := strings.Split(DOB, "/")
	day, _ := strconv.Atoi(dob[0])
	month, _ := strconv.Atoi(dob[1])
	year, _ := strconv.Atoi(dob[2])

	switch {
	case day <= 0 || day > 31:
		return false
	case month <= 0 || month > 12:
		return false
	case year > 2010:
		return false
	}
	return true
}

// checkPublication : validates particular publications only
func checkPublication(publication string) bool {
	strings.ToLower(publication)

	if !(publication == "penguin" || publication == "scholastic" || publication == "arihant") {
		return false
	}
	return true
}

// checkPublishDate : validate the published date
func checkPublishDate(PublishDate string) bool {
	p := strings.Split(PublishDate, "/")
	day, _ := strconv.Atoi(p[0])
	month, _ := strconv.Atoi(p[1])
	year, _ := strconv.Atoi(p[2])

	switch {
	case day < 0 || day > 31:
		return false
	case month < 0 || month > 12:
		return false
	case year > 2022 || year < 1880:
		return false
	}

	return true
}

// PostBook : post a book to the database if author exist
func PostBook(w http.ResponseWriter, req *http.Request) {

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Errorf("failed for %v\n", err)
	}
	var book Book
	json.Unmarshal(body, &book)

	if id, _ := strconv.Atoi(book.BookId); id <= 0 {
		fmt.Println("not valid constraints!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if book.BookId == "" || book.AuthorId <= 0 || book.Author.FirstName == "" || book.Title == "" {
		fmt.Println("not valid constraints!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !checkDob(book.Author.DOB) || !checkPublishDate(book.PublishedDate) || !checkPublication(book.Publication) {
		fmt.Println("not valid constraints!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Db := Connection()

	res := Db.QueryRow("select * from book where bookId=?", book.BookId)
	var checkExitingId Book
	_ = res.Scan(&checkExitingId.BookId)
	if checkExitingId.BookId == book.BookId {
		fmt.Errorf("failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	a, _ := FetchingAuthor(book.AuthorId)
	if a != book.AuthorId {
		fmt.Errorf("author doesnot exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = Db.Exec("insert into book(bookId,authorId,title,publication,publishedDate)values (?,?,?,?,?)", book.BookId,
		book.AuthorId, book.Title, book.Publication, book.PublishedDate)
	if err != nil {
		fmt.Errorf("error:%v\n", err)
		return
	}

	w.Write(body)
	w.WriteHeader(http.StatusCreated)
}

// PostAuthor : post the author to the database
func PostAuthor(w http.ResponseWriter, req *http.Request) {
	body := req.Body

	data, err := io.ReadAll(body)
	if err != nil {
		fmt.Errorf("failed:%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	var author Author
	json.Unmarshal(data, &author)

	a, _ := FetchingAuthor(author.AuthorId)
	if a == author.AuthorId || author.FirstName == "" {
		fmt.Errorf("failed : %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !checkDob(author.DOB) {
		fmt.Println("not valid DOB")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Db := Connection()

	_, err = Db.Exec("insert into author(authorId,firstname,lastName,DOB, penName)values(?,?,?,?,?)", author.AuthorId,
		author.FirstName, author.LastName, author.DOB, author.PenName)
	if err != nil {
		fmt.Errorf("failed %v\n", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// PutBook : updates the particular field in book table
func PutBook(w http.ResponseWriter, req *http.Request) {

	body := req.Body
	params := mux.Vars(req)
	data, err := io.ReadAll(body)
	if err != nil {
		fmt.Errorf("failed:%v\n", err)
		return
	}

	var book Book
	json.Unmarshal(data, &book)

	id, author := FetchingAuthor(book.AuthorId)
	if id != book.AuthorId {
		fmt.Errorf("author does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	book.Author = &author

	Db := Connection()

	if !checkPublishDate(book.PublishedDate) || !checkPublication(book.Publication) || book.Title == "" || !checkDob(book.Author.DOB) {
		fmt.Println("invalid constraints!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var checkExistingBook Book
	row := Db.QueryRow("select * from book where bookId=?", params["id"])
	if err = row.Scan(&checkExistingBook.BookId, &checkExistingBook.AuthorId, &checkExistingBook.Title, &checkExistingBook.Publication, &checkExistingBook.PublishedDate); err == nil {
		_ = Db.QueryRow("delete from book where bookId=?", params["id"])
		_, err = Db.Exec("insert into book(bookId,authorId,title,publication,publishedDate)values(?,?,?,?,?)",
			book.BookId, book.AuthorId, book.Title, book.Publication, book.PublishedDate)

		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	} else {
		_, err = Db.Exec("insert into book(bookId,authorId,title,publication,publishedDate)values(?,?,?,?,?) ",
			book.BookId, book.AuthorId, book.Title, book.Publication, book.PublishedDate)

		w.Write(data)
		w.WriteHeader(http.StatusCreated)
	}

}

// PutAuthor : updates the particular field in author table
func PutAuthor(w http.ResponseWriter, req *http.Request) {

	ReqData, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Errorf("failed:%v\n", err)
		return
	}
	var author Author
	json.Unmarshal(ReqData, &author)

	params := mux.Vars(req)
	Db := Connection()

	if !checkDob(author.DOB) {
		fmt.Println("no valid DOB")
		w.WriteHeader(http.StatusBadRequest)
	}
	id, _ := strconv.Atoi(params["id"])
	var checkExistingAuthor Author
	row := Db.QueryRow("select * from author where authorId=?", id)
	if err = row.Scan(&checkExistingAuthor.AuthorId, &checkExistingAuthor.FirstName, &checkExistingAuthor.LastName,
		&checkExistingAuthor.DOB, &checkExistingAuthor.PenName); err == nil {
		fmt.Println(checkExistingAuthor)
		_ = Db.QueryRow("delete from author where authorId=?", checkExistingAuthor.AuthorId)
		_, err = Db.Exec("insert into author(authorId,firstName,lastName,DOB, penName)values(?,?,?,?,?)",
			author.AuthorId, author.FirstName, author.LastName, author.DOB, author.PenName)

		w.WriteHeader(http.StatusCreated)
		w.Write(ReqData)
	} else {
		fmt.Println(checkExistingAuthor)
		_, err = Db.Exec("insert into author(authorId,firstName,lastName,DOB, penName)values(?,?,?,?,?)",
			author.AuthorId, author.FirstName, author.LastName, author.DOB, author.PenName)

		w.WriteHeader(http.StatusCreated)
		w.Write(ReqData)
	}
}

// DeleteBook : deletes a book send in the path parameter
func DeleteBook(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Errorf("invalid id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if id <= 0 {
		fmt.Println("invalid id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Db := Connection()
	_ = Db.QueryRow("delete from book where bookId=?", params["id"])
	w.WriteHeader(http.StatusNoContent)
}

// DeleteAuthor : removes an author corresponding id
func DeleteAuthor(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Errorf("invalid id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if id <= 0 {
		fmt.Println("invalid id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Db := Connection()
	_ = Db.QueryRow("delete from author where authorId=?", params["id"])
	w.WriteHeader(http.StatusNoContent)
}
