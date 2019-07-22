package dao

import (
	"io"

	"cloud.google.com/go/firestore"
	"github.com/lolibrary/lolibrary/libraries/database"
)

var (
	Firestore *firestore.Client

	featuresByID   *firestore.CollectionRef
	featuresBySlug *firestore.CollectionRef
)

// Init starts up the database access object package and configures a database.
func Init() io.Closer {
	if Firestore == nil {
		Firestore = database.NewClient()
	}

	featuresByID = Firestore.Collection("features")
	featuresBySlug = Firestore.Collection("features-by-slug")

	return Firestore
}
