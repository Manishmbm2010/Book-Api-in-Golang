package driver

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"os"
)

func ConnectDb() *sql.DB {
	dbUrl, err := pq.ParseURL(os.Getenv("DB_URL"))
	logError(err)
	//panicOnError(err)
	db, err := sql.Open("postgres", dbUrl)
	logError(err)
	//panicOnError(err)
	err = db.Ping()
	logError(err)
	//panicOnError(err)
	return db
}

func logError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
