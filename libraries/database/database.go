package database

import (
	"context"

	"cloud.google.com/go/firestore"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/proullon/ramsql/driver"
)

func NewClient() *firestore.Client {
	client, err := firestore.NewClient(context.Background(), "lolibrary-180111")
	if err != nil {
		panic(err)
	}

	return client
}
