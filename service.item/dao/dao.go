package dao

import (
	"io"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/lolibrary/lolibrary/libraries/database"
)

var (
	DB *gorm.DB
)

// Init starts up the database access object package and configures a database.
func Init() io.Closer {
	if DB == nil {
		DB = database.Connect()
		DB.SetLogger(log.New(os.Stdout, "\n", 0))
	}

	return DB
}
