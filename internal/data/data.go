package data

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"sync"
)

var (
	data *Data
	once sync.Once
)

// Data manages the connection to the database.
type Data struct {
	DB *sql.DB
}

// New returns a new instance of Data with the database connection ready.
func New() *Data {
	once.Do(initDB)

	return data
}

// initialize the data variable with the connection to the database.
func initDB() {
	log.Println("Testing connection against database...")
	db, err := getConnection()
	if err != nil {
		log.Panic(err)
	}

	// get the value from .env file, if true, execute migrations
	databaseMigrations := os.Getenv("DATABASE_MIGRATIONS")
	migrate, _ := strconv.ParseBool(databaseMigrations)
	if migrate {
		log.Println("Creating migrations on database...")
		err = MakeMigration(db)
		if err != nil {
			log.Panic(err)
		}
	}

	data = &Data{
		DB: db,
	}
}
