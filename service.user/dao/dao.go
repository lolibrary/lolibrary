package dao

import (
	"io"

	"cloud.google.com/go/firestore"
	"github.com/lolibrary/lolibrary/libraries/database"
)

var (
	Firestore *firestore.Client

	usersByID       *firestore.CollectionRef
	usersByUsername *firestore.CollectionRef
	usersByEmail    *firestore.CollectionRef
)

// Init starts up the database access object package and configures a database.
func Init() io.Closer {
	if Firestore == nil {
		Firestore = database.NewClient()
	}

	usersByID = Firestore.Collection("users")
	usersByUsername = Firestore.Collection("users-by-username")
	usersByEmail = Firestore.Collection("users-by-email")

	return Firestore
}
