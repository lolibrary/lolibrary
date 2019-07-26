package dao

import (
	"io"

	"cloud.google.com/go/firestore"
	"github.com/lolibrary/lolibrary/libraries/database"
)

var (
	Firestore *firestore.Client

	categoriesByID   *firestore.CollectionRef
	categoriesBySlug *firestore.CollectionRef
)

// Init starts up the database access object package and configures a database.
func Init() io.Closer {
	if Firestore == nil {
		Firestore = database.NewClient()
	}

	categoriesByID = Firestore.Collection("categories")
	categoriesBySlug = Firestore.Collection("categories-by-slug")

	return Firestore
}
