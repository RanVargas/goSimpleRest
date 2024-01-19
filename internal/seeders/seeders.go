package seeders

import (
	"database/sql"
	"log"
)

func CaptureDbData(db *sql.DB) {
	//TODO read all DB data to prepare Seed
	//TODO call file handler to have all data captured written to files
}

func SeedBooks(db *sql.DB) error {
	_, err := db.Exec(`
        INSERT INTO books (title, author, content)
        VALUES
            ('Book 1', 'Author 1', 'Content 1'),
            ('Book 2', 'Author 2', 'Content 2'),
            ('Book 3', 'Author 3', 'Content 3');
    `)
	if err != nil {
		log.Printf("Error seeding books: %v", err)
		return err
	}
	return nil
}
