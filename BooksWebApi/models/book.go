package models

import (
	"BooksWebApi/db"
	"fmt"
)

type Book struct {
	Id          int64
	Title       string
	Isbn        string
	Author      string
	ReleaseYear int
}

var bookCollection []Book

func (b *Book) Save() error {

	statment, err := db.GetDb().Prepare(`
		INSERT INTO 
		books
		    (title, isbn, author, release_year)
		VALUES
		    (?, ?, ?, ?)
	`)

	defer statment.Close()

	if err != nil {
		return err
	}

	result, err := statment.Exec(b.Title, b.Isbn, b.Author, b.ReleaseYear)
	if err != nil {
		return err
	}

	bookId, err := result.LastInsertId()
	b.Id = bookId

	return err
}

func GetAllBooks() ([]Book, error) {

	dbCursor, err := db.GetDb().Query(`SELECT * FROM books`)
	if err != nil {
		return nil, err
	}

	for dbCursor.Next() {

		var bookObject Book
		err := dbCursor.Scan(
			&bookObject.Id,
			&bookObject.Title,
			&bookObject.Isbn,
			&bookObject.Author,
			&bookObject.ReleaseYear,
		)

		if err != nil {
			return nil, err
		}

		bookCollection = append(bookCollection, bookObject)
	}

	return bookCollection, nil
}

func GetBookById(id int64) (Book, error) {
	dbCursor := db.GetDb()
	var bookObject Book

	row := dbCursor.QueryRow(`SELECT * FROM books WHERE id = ?`, id)
	err := row.Scan(
		&bookObject.Id,
		&bookObject.Title,
		&bookObject.Isbn,
		&bookObject.Author,
		&bookObject.ReleaseYear,
	)
	if err != nil {
		return bookObject, err
	}

	return bookObject, nil
}

func (b *Book) UpdateBook() error {
	statement, err := db.GetDb().Prepare(`
		UPDATE books 
		SET title = ?, isbn = ?, author = ?, release_year = ?
		WHERE id = ?
	`)

	defer statement.Close()

	if err != nil {
		return err
	}

	_, err = statement.Exec(b.Title, b.Isbn, b.Author, b.ReleaseYear, b.Id)
	if err != nil {
		return err
	}
	return err
}

func DeleteBook(id int64) error {
	dbCursor := db.GetDb()

	statement, err := dbCursor.Prepare(`DELETE FROM books WHERE id = ?`)

	defer statement.Close()

	result, err := statement.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no book found with id %d", id)
	}

	return nil

}
