package driver

import (
	"os"
	"log"
	"database/sql"

	"github.com/lib/pq"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB{
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err := sql.Open("postgres", pgURL)
	logFatal(err)

	errorConn := db.Ping()
	logFatal(errorConn)

	return db
}