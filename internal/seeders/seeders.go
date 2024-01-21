package seeders

import (
	"booklib/internal/database"
	"booklib/internal/filehelper"
	"booklib/pkg/model"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func CaptureDbData() error {
	//TODO read all DB data to prepare Seed
	//TODO call file handler to have all data captured written to files
	err := godotenv.Load("D:\\Practice\\ran-vargas-naithergrand-2024-1-16-dev-backend-coding-challenge-library-api\\local.env")
	if err != nil {
		panic(err)
	}
	folder := os.Getenv("OUTPUT_FOLDER")
	var outputBooksFile = folder + "savedBooks.xlsx"
	var books []model.Book
	//var content model.BookPage
	books, dbErr := database.GetAllBooks()
	if dbErr != nil {
		return dbErr
	}
	//outputContentFiles := make([]string, 0)
	filehelper.WriteExcelFile(outputBooksFile, books)

	for _, book := range books {
		content, err := database.GetBookContentByIdUnpaginated(book.ID)
		if err != nil {
			log.Print(err)
			continue
		}
		for _, page := range content {
			path := folder + fmt.Sprint(page.BookID) + "-" + fmt.Sprint(page.PageNum) + ".txt"
			filehelper.WriteToPlainTextFile(path, page.Content)
		}

	}
	return nil
}

func SeedBooks() error {
	db, dbUrl := database.GetDbToMigrate()
	db, conErr := sql.Open("postgres", dbUrl)
	if conErr != nil {
		return conErr
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	books := make([]model.Book, 0)
	books = append(books, model.Book{
		Title:  "Crime and Punishment",
		Author: "Fyodor Dostoevsky",
	})
	books = append(books, model.Book{
		Title:  "1984",
		Author: "George Orwell",
	})
	books = append(books, model.Book{
		Title:  "The Happy Prince and Other Tales",
		Author: "Oscar Wilde",
	})
	books = append(books, model.Book{
		Title:  "The Birth of Tragedy",
		Author: "Friedrich Nietzsche",
	})
	books = append(books, model.Book{
		Title:  "Fear and Trembling",
		Author: "Søren Kierkegaard",
	})
	books = append(books, model.Book{
		Title:  "The Myth of Sisyphus",
		Author: "Albert Camus",
	})
	books = append(books, model.Book{
		Title:  "Neuromancer",
		Author: "William Gibson",
	})

	stmt, err := tx.Prepare(`INSERT INTO book (title, author) VALUES ($1, $2) RETURNING id`)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, book := range books {

		var bookId int
		err := stmt.QueryRow(book.Title, book.Author).Scan(&bookId)
		if err != nil {
			tx.Rollback()
			return err
		}
		placeholder := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Imperdiet sed euismod nisi porta lorem mollis aliquam. Sed sed risus pretium quam vulputate dignissim suspendisse in est. Tempus quam pellentesque nec nam aliquam sem et tortor consequat. Tellus mauris a diam maecenas sed enim ut sem. Fermentum posuere urna nec tincidunt praesent semper feugiat nibh. Massa eget egestas purus viverra. Aliquet nibh praesent tristique magna sit amet purus gravida. Ipsum consequat nisl vel pretium lectus quam id leo in. Sit amet nulla facilisi morbi tempus iaculis urna. Nulla facilisi etiam dignissim diam quis enim lobortis scelerisque.

		Rutrum quisque non tellus orci ac auctor. Morbi tristique senectus et netus et malesuada. Purus gravida quis blandit turpis cursus in hac habitasse. Scelerisque eu ultrices vitae auctor eu augue ut. Venenatis cras sed felis eget. Id faucibus nisl tincidunt eget nullam non. Ornare lectus sit amet est placerat in egestas erat. Dui sapien eget mi proin sed libero. Eget arcu dictum varius duis. Sem et tortor consequat id. Purus non enim praesent elementum facilisis. Facilisis sed odio morbi quis. Quis hendrerit dolor magna eget est lorem ipsum dolor. Quam quisque id diam vel quam elementum. Ultrices eros in cursus turpis massa tincidunt dui ut ornare. A condimentum vitae sapien pellentesque habitant morbi tristique senectus.
		
		Mauris nunc congue nisi vitae suscipit tellus mauris a. Mattis aliquam faucibus purus in massa tempor nec feugiat nisl. Velit euismod in pellentesque massa. Vel quam elementum pulvinar etiam non quam lacus. Nulla aliquet enim tortor at auctor urna. Purus non enim praesent elementum. Nunc sed augue lacus viverra vitae. Sit amet porttitor eget dolor morbi non arcu. Netus et malesuada fames ac turpis egestas integer eget. Sit amet facilisis magna etiam tempor orci. Morbi blandit cursus risus at ultrices. Ultrices neque ornare aenean euismod elementum. Ac turpis egestas integer eget aliquet nibh praesent tristique. Semper quis lectus nulla at volutpat diam.
		
		Morbi non arcu risus quis varius quam quisque. Sit amet tellus cras adipiscing enim eu turpis egestas pretium. Scelerisque eu ultrices vitae auctor eu augue ut lectus. Ut sem nulla pharetra diam sit amet nisl. Aliquam purus sit amet luctus venenatis lectus. Enim diam vulputate ut pharetra sit amet. Elit sed vulputate mi sit amet mauris. Dictum non consectetur a erat nam. Vel pharetra vel turpis nunc eget lorem dolor sed viverra. Turpis egestas integer eget aliquet nibh praesent tristique. Ac odio tempor orci dapibus. Praesent elementum facilisis leo vel fringilla est. Risus feugiat in ante metus dictum. Metus vulputate eu scelerisque felis imperdiet. Scelerisque viverra mauris in aliquam sem fringilla ut. Eu lobortis elementum nibh tellus. Lacus viverra vitae congue eu consequat ac felis donec et. Vulputate sapien nec sagittis aliquam malesuada bibendum arcu vitae. Velit egestas dui id ornare. Platea dictumst vestibulum rhoncus est pellentesque elit ullamcorper.
		
		Vel facilisis volutpat est velit egestas. Vulputate eu scelerisque felis imperdiet proin fermentum leo vel orci. Integer quis auctor elit sed vulputate mi sit. Iaculis at erat pellentesque adipiscing commodo elit at imperdiet dui. Ac auctor augue mauris augue. Maecenas volutpat blandit aliquam etiam erat velit scelerisque. Tristique senectus et netus et malesuada fames. Ipsum faucibus vitae aliquet nec ullamcorper sit amet risus. In metus vulputate eu scelerisque felis. Cras ornare arcu dui vivamus arcu felis bibendum ut. Tortor consequat id porta nibh venenatis cras sed. Cras sed felis eget velit aliquet sagittis id consectetur. Massa tincidunt nunc pulvinar sapien et ligula. Odio ut sem nulla pharetra diam sit amet.`
		_, err = tx.Exec(`INSERT INTO book_content (book_id, content, page_num) VALUES ($1, $2, $3)`,
			bookId, placeholder, 1)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	defer db.Close()
	return tx.Commit()
}

func SeedBooksFromFile() error {
	db, dbUrl := database.GetDbToMigrate()
	db, conErr := sql.Open("postgres", dbUrl)
	if conErr != nil {
		return conErr
	}
	var books []model.Book
	var booksContent []model.BookPage
	err := godotenv.Load("D:\\Practice\\ran-vargas-naithergrand-2024-1-16-dev-backend-coding-challenge-library-api\\local.env")
	if err != nil {
		panic(err)
	}
	folder := os.Getenv("OUTPUT_FOLDER")
	dirFiles := filehelper.GetFilesInThisDirectory(folder)
	for _, file := range dirFiles {
		extension := path.Ext(file)
		if extension == ".txt" {
			pureName := strings.TrimSuffix(file, path.Ext(file))
			ids := strings.Split(pureName, "-")

			content, err := filehelper.ReadFromPlainTextFile(file)
			if err != nil {
				log.Println(err)
			}
			bookId, err := strconv.Atoi(ids[0])

			pagNum, err := strconv.Atoi(ids[1])
			bookContent := model.BookPage{
				BookID:  bookId,
				PageNum: pagNum,
				Content: content,
			}
			booksContent = append(booksContent, bookContent)
		} else if extension == ".xlsx" {
			books = filehelper.ReadExcelFile(extension)
		}

	}

	insertBooks := "INSERT INTO book(author, title) VALUES ($1, $2)"

	for _, book := range books {
		_, err := db.Exec(insertBooks, book.Author, book.Title)
		if err != nil {
			log.Fatal(err)
			return err
		}

		insertBooksContent := "INSERT INTO book(book_id, page_num, content) VALUES ($1, $2, $2)"

		for _, bookContent := range booksContent {
			_, err := db.Exec(insertBooksContent, bookContent.BookID, bookContent.PageNum, bookContent.Content)
			if err != nil {
				log.Fatal(err)
				return err
			}
		}
	}
	defer db.Close()
	return nil
}
