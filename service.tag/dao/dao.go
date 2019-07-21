package dao

import (
	"context"

	"cloud.google.com/go/firestore"
)

var (
	Firestore *firestore.Client

	tagsByID   *firestore.CollectionRef
	tagsBySlug *firestore.CollectionRef
)

// Init starts up the database access object package and configures a database.
func Init() {
	if Firestore == nil {
		// add this to database package later
		client, err := firestore.NewClient(context.Background(), "lolibrary-180111")
		if err != nil {
			panic(err)
		}

		Firestore = client
	}

	tagsByID = Firestore.Collection("tags")
	tagsBySlug = Firestore.Collection("tags-by-slug")
}
