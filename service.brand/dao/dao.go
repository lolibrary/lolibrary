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

	brandsByID = Firestore.Collection("attributes")
	brandsBySlug = Firestore.Collection("attributes-by-slug")
	brandsByShortName = Firestore.Collection("attributes-by-item-id")

	return Firestore
}
