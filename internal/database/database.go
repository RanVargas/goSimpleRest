package database

import (
	"booklib/pkg/model"
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const (
	pageSize = 50
	dbType   = "postgres"
)

var (
	db          *sql.DB
	dbUser      string
	dbPassword  string
	dbName      string
	sqlFilePath string
)
var dbURL = "postgresql://" + dbUser + ":" + dbPassword + "@localhost:5432/" + dbName + "?sslmode=disable"

func GetDbToMigrate() (*sql.DB, string) {
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	sqlFilePath = os.Getenv("PATH_TO_SQL_FILE_MIGRATION")
	host := "localhost"
	port := 5432
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, dbUser, dbPassword, dbName)
	return db, psqlInfo
}

// InitDB initializes the database connection.
func InitDB() {
	connectErr := StartConnection()
	if connectErr != nil {
		fmt.Println(connectErr)
		panic(connectErr)
	}

	triggerDbMigration, err := isMigrationNeeded(db)
	if err != nil {
		fmt.Println(err)
	}
	if triggerDbMigration {
		panic("DB has not been setup")
	}
	defer CloseConnection()
}

func GetBooks(page int) ([]model.Book, error) {
	// Implement your logic to fetch books from the database.
	fmt.Println(dbURL)
	connectErr := StartConnection()
	if connectErr != nil {
		fmt.Println(connectErr)
		panic(connectErr)
	}
	offset := (page - 1) * pageSize
	var books []model.Book
	rows, err := db.Query("SELECT id, title, author FROM book ORDER BY id LIMIT $1 OFFSET $2", pageSize, offset)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.Author, &book.Title)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	defer CloseConnection()
	return books, nil
}

func GetAllBooks() ([]model.Book, error) {
	// Implement your logic to fetch books from the database.
	connectErr := StartConnection()
	if connectErr != nil {
		fmt.Println(connectErr)
		panic(connectErr)
	}
	var books []model.Book
	rows, err := db.Query("SELECT id, title, author FROM book ORDER BY id")

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.Author, &book.Title)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	defer CloseConnection()
	return books, nil
}

func GetBookByID(bookID int) (model.Book, error) {
	connectErr := StartConnection()
	if connectErr != nil {
		fmt.Println(connectErr)
		panic(connectErr)
	}
	var book model.Book
	err := db.QueryRow("SELECT id, title, author FROM book WHERE id = $1", bookID).
		Scan(&book.ID, &book.Title, &book.Author)

	if err != nil {
		return model.Book{}, err
	}

	defer CloseConnection()
	return book, nil
}

// GetBookContentByID retrieves the content of a book by pages from the database.
func GetBookContentByID(bookID, page int) (model.BookPage, error) {
	connectErr := StartConnection()
	if connectErr != nil {
		fmt.Println(connectErr)
		panic(connectErr)
	}
	// Implement your logic to fetch book content by ID and page from the database.
	var bookContent model.BookPage
	err := db.QueryRow("SELECT id, book_id, page_num, content FROM book_content WHERE book_id = $1 AND page_num = $2", bookID, page).
		Scan(&bookContent.ID, &bookContent.BookID, &bookContent.PageNum, &bookContent.Content)

	if err != nil {
		return model.BookPage{}, err
	}

	defer CloseConnection()
	return bookContent, nil
}

func GetBookContentByIdUnpaginated(bookID int) ([]model.BookPage, error) {
	connectErr := StartConnection()
	if connectErr != nil {
		fmt.Println(connectErr)
		panic(connectErr)
	}
	var bookContent []model.BookPage

	rows, err := db.Query("SELECT id, book_id, page_num, content FROM book_content WHERE book_id = $1", bookID)

	/*
			id SERIAL PRIMARY KEY,
		    book_id INT,
		    page_num INT,
		    content TEXT
	*/
	if err != nil {
		return []model.BookPage{}, err
	}

	for rows.Next() {
		var bookPage model.BookPage
		err := rows.Scan(&bookPage.ID, &bookPage.BookID, &bookPage.PageNum, &bookPage.Content)
		if err != nil {
			continue
		}
		bookContent = append(bookContent, bookPage)
	}

	defer CloseConnection()
	return bookContent, nil
}

func isMigrationNeeded(db *sql.DB) (bool, error) {

	// Check if the "users" table exists as an example
	var exists bool
	err := db.QueryRow(`SELECT EXISTS (
                           SELECT FROM information_schema.tables 
                           WHERE  table_schema = 'public'
                           AND    table_name   = 'book');`).Scan(&exists)
	if err != nil {
		return false, err
	}

	// Doesn't exist -> migration needed
	if !exists {
		return true, nil
	}

	// Table exists -> migration not needed
	return false, nil
}

func CloseConnection() error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}

func StartConnection() error {
	var err error

	err = godotenv.Load("D:\\Practice\\ran-vargas-naithergrand-2024-1-16-dev-backend-coding-challenge-library-api\\local.env")
	if err != nil {
		panic(err)
	}

	dbUser = os.Getenv("DB_USER")

	var drive pq.Driver
	fmt.Println(drive)
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	sqlFilePath = os.Getenv("PATH_TO_SQL_FILE_MIGRATION")
	fmt.Println(dbUser, " This is the password ", dbPassword, " DBName is: ", dbName, " sqlfilepath ", sqlFilePath)
	host := "localhost"
	port := 5432
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, dbUser, dbPassword, dbName)
	//anotherPsql := "host=" + host + " port=" + port + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"

	fmt.Println(dbURL)
	db, err = sql.Open("postgres", psqlInfo)
	return err
}
