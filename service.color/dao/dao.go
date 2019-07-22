package dao

import (
	"io"

	"cloud.google.com/go/firestore"
	"github.com/lolibrary/lolibrary/libraries/database"
)

var (
	Firestore *firestore.Client

	colorsByID   *firestore.CollectionRef
	colorsBySlug *firestore.CollectionRef
)

// Init starts up the database access object package and configures a database.
func Init() io.Closer {
	if Firestore == nil {
		Firestore = database.NewClient()
	}

	colorsByID = Firestore.Collection("colors")
	colorsBySlug = Firestore.Collection("colors-by-slug")

	return Firestore
}
