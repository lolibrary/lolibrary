package dao

import (
	"io"

	"cloud.google.com/go/firestore"
	"github.com/lolibrary/lolibrary/libraries/database"
)

var (
	Firestore *firestore.Client

	imagesByID       *firestore.CollectionRef
	imagesByFilename *firestore.CollectionRef
)

// Init starts up the database access object package and configures a database.
func Init() io.Closer {
	if Firestore == nil {
		Firestore = database.NewClient()
	}

	imagesByID = Firestore.Collection("images")
	imagesByFilename = Firestore.Collection("images-by-filename")

	return Firestore
}
