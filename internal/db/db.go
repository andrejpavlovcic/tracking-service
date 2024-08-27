package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func InitDB() error {
	var err error
	db, err = sql.Open("postgres", getDatabaseStringConnection())
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	return nil
}

func GetDB() *sql.DB {
	if db == nil {
		InitDB()
	}

	return db
}

func getDatabaseStringConnection() string {
	var (
		dbHost             = os.Getenv("DB_HOST")
		dbPort             = os.Getenv("DB_PORT")
		dbUser             = os.Getenv("DB_USER")
		dbPassword         = os.Getenv("DB_PASSWORD")
		dbName             = os.Getenv("DB_NAME")
		dbConnectionString = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbHost,
			dbPort,
			dbUser,
			dbPassword,
			dbName,
		)
	)

	return dbConnectionString
}
