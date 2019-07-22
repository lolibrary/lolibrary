package dao

import (
	"io"

	"cloud.google.com/go/firestore"
	"github.com/lolibrary/lolibrary/libraries/database"
)

var (
	Firestore *firestore.Client

	attributesByID     *firestore.CollectionRef
	attributesBySlug   *firestore.CollectionRef
	attributesByItemID *firestore.CollectionRef
)

// Init starts up the database access object package and configures a database.
func Init() io.Closer {
	if Firestore == nil {
		Firestore = database.NewClient()
	}

	attributesByID = Firestore.Collection("attributes")
	attributesBySlug = Firestore.Collection("attributes-by-slug")
	attributesByItemID = Firestore.Collection("attributes-by-item-id")

	return Firestore
}
