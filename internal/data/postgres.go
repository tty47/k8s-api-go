package data

import (
	"database/sql"
	"io/ioutil"
	"os"
)

type PostgresDB struct {
	db *sql.DB
}

// getConnection open connection against PSQL database
func getConnection() (*sql.DB, error) {
	// generate the uri using env vars
	uri := os.Getenv("DB_DRIVER") + "://" + os.Getenv("DB_USER") +":" +
		os.Getenv("DB_PASSWORD") +"@" +os.Getenv("DB_HOST") +":"+ os.Getenv("DB_PORT") +
		"/"+os.Getenv("DB_NAME")+"?&sslmode=disable"
	return sql.Open(os.Getenv("DB_DRIVER") , uri)
}

// MakeMigration setup the database structure
func MakeMigration(db *sql.DB) error {
	b, err := ioutil.ReadFile("./internal/data/init.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}
