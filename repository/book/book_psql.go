package bookRepository

import (
	"log"
	"database/sql"
	"../../models"
)

type BookRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
	rows, err := db.Query(`SELECT
									id,
									title,
									author,
									year
								FROM book
								ORDER BY id ASC
								`)
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Year,
		)
		logFatal(err)

		books = append(books, book)
	}

	return books
}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) models.Book{
	row := db.QueryRow("SELECT id,title,author,year FROM book WHERE id=$1", id)

	err := row.Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Year,
	)
	logFatal(err)

	return book
}

func (b BookRepository) AddBook(db *sql.DB, book models.Book) int{
	sqlStatement := `INSERT INTO book (title, author, year) 
					VALUES ($1, $2, $3) 
					RETURNING id;`
	err := db.QueryRow(sqlStatement, book.Title, book.Author, book.Year).Scan(&book.ID)
	logFatal(err)

	return book.ID
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) int64{
	sqlStatement := `UPDATE book set 
					title = $2,
					author = $3,
					year = $4
					WHERE id = $1
					RETURNIng id;`

	res, err := db.Exec(sqlStatement, book.ID, book.Title, book.Author, book.Year)
	logFatal(err)

	count, err := res.RowsAffected()
	logFatal(err)

	return count
}

func (b BookRepository) RemoveBook(db *sql.DB, id int) int64{
	sqlStatement := `DELETE from book WHERE id=$1;`
	res, err := db.Exec(sqlStatement, id)
	logFatal(err)

	count, err := res.RowsAffected()
	logFatal(err)

	return count
}
