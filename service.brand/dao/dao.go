package dao

import (
	"io"

	"cloud.google.com/go/firestore"
	"github.com/lolibrary/lolibrary/libraries/database"
)

var (
	Firestore *firestore.Client

	brandsByID        *firestore.CollectionRef
	brandsBySlug      *firestore.CollectionRef
	brandsByShortName *firestore.CollectionRef
)

// Init starts up the database access object package and configures a database.
func Init() io.Closer {
	if Firestore == nil {
		Firestore = database.NewClient()
	}

	brandsByID = Firestore.Collection("brands")
	brandsBySlug = Firestore.Collection("brands-by-slug")
	brandsByShortName = Firestore.Collection("brands-by-short-name")

	return Firestore
}
