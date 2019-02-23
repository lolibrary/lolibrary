package database

import (
	"os"

	"github.com/jmoiron/sqlx"
)

// New returns a new database instance.
// It should be instantiated in the main function and shared by all goroutines.
// This function reads directly from the environment as a shared service.
func New() *sqlx.DB {
	conn := os.Getenv("DATABASE_URL")

	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		panic(err)
	}

	return db
}
