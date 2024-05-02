package postgres

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbPath := os.Getenv("DB_PATH")

	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
