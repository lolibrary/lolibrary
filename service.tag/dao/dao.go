package dao

import (
	"io"

	"cloud.google.com/go/firestore"
	"github.com/lolibrary/lolibrary/libraries/database"
)

var (
	Firestore *firestore.Client

	tagsByID   *firestore.CollectionRef
	tagsBySlug *firestore.CollectionRef
)

// Init starts up the database access object package and configures a database.
func Init() io.Closer {
	if Firestore == nil {
		Firestore = database.NewClient()
	}

	tagsByID = Firestore.Collection("tags")
	tagsBySlug = Firestore.Collection("tags-by-slug")

	return Firestore
}
