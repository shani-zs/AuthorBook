package authorbook

import (
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
	AuthorID  int    `json:"authorID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DOB       string `json:"DOB"`
	PenName   string `json:"penName"`
}

type Book struct {
	BookID        string  `json:"bookID"`
	AuthorID      int     `json:"authorID"`
	Title         string  `json:"title"`
	Publication   string  `json:"publication"`
	PublishedDate string  `json:"publishedDate"`
	Author        *Author `json:"author"`
}

//GetAllBook : returns all books to the client
func GetAllBook(w http.ResponseWriter, req *http.Request) {

	title := req.URL.Query().Get("title")
	includeAuthor := req.URL.Query().Get("includeAuthor")

	var books []Book

	switch {
	case title != "" && includeAuthor == "true":
		books = BooksWithAuthor(title)

		data, err := json.Marshal(books)
		if err != nil {
			log.Print(w, "does not exist")
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		_, _ = w.Write(data)
		w.WriteHeader(http.StatusOK)
		return

	case title != "" && includeAuthor == "false" || includeAuthor == "":
		books = FetchingAllBooks(title)

		data, err := json.Marshal(books)
		if err != nil {
			log.Print(w, "does not exist")
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		_, _ = w.Write(data)
		w.WriteHeader(http.StatusOK)
		return

	case includeAuthor == "true":
		books = BooksWithAuthor("")

		data, err := json.Marshal(books)
		if err != nil {
			log.Print(w, "does not exist")
			w.WriteHeader(http.StatusBadRequest)

			return
		}

		_, _ = w.Write(data)
		w.WriteHeader(http.StatusOK)
		return

	default:
		fmt.Println("hello")
		books = FetchingAllBooks("")
		data, err := json.Marshal(books)
		if err != nil {
			log.Print(w, "does not exist")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, _ = w.Write(data)
		w.WriteHeader(http.StatusOK)
		return
	}
}

// GetBookByID : returns a single book
func GetBookByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	if req.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	StringErr := strings.ToLower(params["id"])
	log.Print(StringErr)
	id, _ := strconv.Atoi(params["id"])
	if id <= 0 {
		log.Printf("invalid id")
		w.WriteHeader(http.StatusBadRequest)
	}

	DB := Connection()

	row := DB.QueryRow("select * from book where bookId=?", params["id"])
	var b Book
	if err := row.Scan(&b.BookID, &b.AuthorID, &b.Title, &b.Publication, &b.PublishedDate); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)

		return
	}

	_, author := FetchingAuthor(b.AuthorID)
	b.Author = &author

	EncodedBook, err := json.Marshal(b)
	if err != nil {
		log.Print(err)
		return
	}

	_, _ = w.Write(EncodedBook)
	w.WriteHeader(http.StatusOK)
}

// PostBook : post a book to the database if author exist
func PostBook(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}

	var book Book
	err = json.Unmarshal(body, &book)
	log.Print(err)

	if book.BookID == "" || book.AuthorID <= 0 || book.Author.FirstName == "" || book.Title == "" || book.BookID[0] == '-' {
		log.Print("not valid constraints!")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	if !checkDob(book.Author.DOB) || !checkPublishDate(book.PublishedDate) || !checkPublication(book.Publication) {
		log.Print("not valid constraints!")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	DB := Connection()

	row := DB.QueryRow("select * from book where bookId=?", book.BookID)
	var checkExitingID Book
	err = row.Scan(&checkExitingID.BookID)
	if checkExitingID.BookID == book.BookID {
		log.Print("failed")
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	a, _ := FetchingAuthor(book.AuthorID)
	if a != book.AuthorID {
		log.Print("author does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = DB.Exec("insert into book(bookId,authorId,title,publication,publishedDate)values (?,?,?,?,?)", book.BookID,
		book.AuthorID, book.Title, book.Publication, book.PublishedDate)
	if err != nil {
		log.Print(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, _ = fmt.Fprintf(w, "successfully created!")
}

// PostAuthor : post the author to the database
func PostAuthor(w http.ResponseWriter, req *http.Request) {
	body := req.Body
	ReqData, err := io.ReadAll(body)
	if err != nil {
		log.Printf("failed:%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	var author Author
	err = json.Unmarshal(ReqData, &author)
	log.Print(err)

	_, existingAuthor := FetchingAuthor(author.AuthorID)
	if existingAuthor.AuthorID == author.AuthorID || author.FirstName == "" || author.AuthorID <= 0 || !checkDob(author.DOB) {
		log.Print("invalid constraints!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	DB := Connection()
	_, err = DB.Exec("insert into author(authorId,firstname,lastName,DOB, penName)values(?,?,?,?,?)", author.AuthorID,
		author.FirstName, author.LastName, author.DOB, author.PenName)
	if err != nil {
		log.Print(err)
		return
	}
	_, _ = fmt.Fprintf(w, "successfully created!")
	w.WriteHeader(http.StatusCreated)
}

// PutBook : updates the particular field in book table and if not exits then creates
func PutBook(w http.ResponseWriter, req *http.Request) {
	body := req.Body
	params := mux.Vars(req)
	ReqBody, err := io.ReadAll(body)
	if err != nil {
		log.Print(err)
		return
	}

	var book Book
	_ = json.Unmarshal(ReqBody, &book)

	id, author := FetchingAuthor(book.AuthorID)
	if id != book.AuthorID {
		log.Print("author does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	book.Author = &author

	DB := Connection()

	if !checkPublishDate(book.PublishedDate) || !checkPublication(book.Publication) || book.Title == "" || !checkDob(book.Author.DOB) {
		log.Print("invalid constraints!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var checkExistingBook Book
	row := DB.QueryRow("select * from book where bookId=?", params["id"])

	if err = row.Scan(&checkExistingBook.BookID, &checkExistingBook.AuthorID, &checkExistingBook.Title,
		&checkExistingBook.Publication, &checkExistingBook.PublishedDate); err == nil {
		_, _ = DB.Exec("update book set bookId=?,authorId=?,title=?,publication=?,publishedDate=? where bookId=?",
			book.BookID, book.AuthorID, book.Title, book.Publication, book.PublishedDate, params["id"])

		_, _ = fmt.Fprintf(w, "successfull updated !")
		w.WriteHeader(http.StatusCreated)
		return

	} else {
		_, _ = DB.Exec("insert into book(bookId,authorId,title,publication,publishedDate)values(?,?,?,?,?) ",
			book.BookID, book.AuthorID, book.Title, book.Publication, book.PublishedDate)

		_, _ = fmt.Fprintf(w, "successful inserted ")
		w.WriteHeader(http.StatusCreated)
		return
	}
}

// PutAuthor : updates the particular field in author table
func PutAuthor(w http.ResponseWriter, req *http.Request) {
	ReqData, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("failed:%v\n", err)
		return
	}
	var author Author
	err = json.Unmarshal(ReqData, &author)
	log.Print(err)

	params := mux.Vars(req)
	DB := Connection()

	if !checkDob(author.DOB) {
		log.Print("no valid DOB")
		w.WriteHeader(http.StatusBadRequest)
	}

	id, _ := strconv.Atoi(params["id"])
	var checkExistingAuthor Author

	row := DB.QueryRow("select * from author where authorId=?", id)
	if err = row.Scan(&checkExistingAuthor.AuthorID, &checkExistingAuthor.FirstName, &checkExistingAuthor.LastName,
		&checkExistingAuthor.DOB, &checkExistingAuthor.PenName); err == nil {
		_, _ = DB.Exec("update author set authorId=?,firstName=?,lastName=?,DOB=?,penName=? where authorId=?",
			author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName, id)

		_, _ = fmt.Fprintf(w, "successful updated ")
		w.WriteHeader(http.StatusCreated)
		return
	} else {
		_, _ = DB.Exec("insert into author(authorId,firstName,lastName,DOB, penName)values(?,?,?,?,?)",
			author.AuthorID, author.FirstName, author.LastName, author.DOB, author.PenName)

		_, _ = fmt.Fprintf(w, "successful inserted ")
		w.WriteHeader(http.StatusCreated)
		return
	}
}

// DeleteBook : deletes a book send in the path parameter
func DeleteBook(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Print("invalid id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if id <= 0 {
		_, _ = fmt.Println("invalid id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	DB := Connection()
	_ = DB.QueryRow("delete from book where bookId=?", params["id"])

	_, _ = fmt.Fprintf(w, "successfully deleted ")
	w.WriteHeader(http.StatusNoContent)
}

// DeleteAuthor : removes an author corresponding id
func DeleteAuthor(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Print("invalid id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data := params["id"]
	if data[0] == '-' {
		log.Print("negative id!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	DB := Connection()
	_ = DB.QueryRow("delete from author where authorId=?", id)

	_, _ = fmt.Fprintf(w, "successfully deleted ")
	w.WriteHeader(http.StatusNoContent)
}

// checkDob : Validate the DOB of the author
func checkDob(dob string) bool {

	Dob := strings.Split(dob, "/")
	day, _ := strconv.Atoi(Dob[0])
	month, _ := strconv.Atoi(Dob[1])
	year, _ := strconv.Atoi(Dob[2])

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
	_ = strings.ToLower(publication)

	return !(publication == "penguin" || publication == "scholastic" || publication == "arihant")
}

// checkPublishDate : validate the published date
func checkPublishDate(publishDate string) bool {
	p := strings.Split(publishDate, "/")
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

// FetchingAllBooks : fetches all books from the database
func FetchingAllBooks(title string) []Book {
	DB := Connection()
	var Rows *sql.Rows
	var err error

	if title == "" {
		Rows, err = DB.Query("SELECT * FROM book")
		if err != nil {
			log.Print(err)
		}

	} else {
		Rows, err = DB.Query("SELECT * FROM book where title=?", title)
		if err != nil {
			log.Print(err)
		}
	}

	var bk []Book

	for Rows.Next() {
		var b Book
		err = Rows.Scan(&b.BookID, &b.AuthorID, &b.Title, &b.Publication, &b.PublishedDate)
		if err != nil {
			log.Print(err)
		}

		bk = append(bk, b)
	}

	return bk
}

// Connection : makes the connection to the database
func Connection() *sql.DB {
	Db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/authorbook")
	if err != nil {
		log.Fatal("failed to connect with database:\n", err)
	}

	pingErr := Db.Ping()
	if pingErr != nil {
		log.Fatal("failed to ping", pingErr)
	}

	return Db
}

// FetchingAuthor : gets the author detail from the database
func FetchingAuthor(id int) (int, Author) {
	Db := Connection()
	defer Db.Close()

	Row := Db.QueryRow("SELECT * FROM author where authorId=?", id)
	var author Author
	if err := Row.Scan(&author.AuthorID, &author.FirstName, &author.LastName, &author.DOB, &author.PenName); err != nil {
		_ = fmt.Errorf("failed: %v\n", err)
		return 0, Author{}
	}
	return author.AuthorID, author
}

// BooksWithAuthor :returns all books with author
func BooksWithAuthor(title string) []Book {
	Db := Connection()
	var Rows *sql.Rows
	var err error

	if title == "" {
		Rows, err = Db.Query("SELECT * FROM book")
		if err != nil {
			log.Print(err)
			return []Book{}
		}
	} else {
		Rows, err = Db.Query("SELECT * FROM book where title=?", title)
		if err != nil {
			log.Print(err)
			return []Book{}
		}
	}

	var bk []Book

	for Rows.Next() {
		var b Book
		err := Rows.Scan(&b.BookID, &b.AuthorID, &b.Title, &b.Publication, &b.PublishedDate)
		if err != nil {
			log.Print(err)
			return []Book{}
		}

		_, a := FetchingAuthor(b.AuthorID)
		b.Author = &a
		bk = append(bk, b)
	}

	return bk
}
