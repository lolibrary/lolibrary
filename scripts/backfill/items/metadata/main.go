package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/lolibrary/lolibrary/libraries/backfill"
	itemproto "github.com/lolibrary/lolibrary/service.item/proto"
)

var metaPath = path.Join(
	os.Getenv("GOPATH"), "src/github.com/lolibrary/lolibrary",
	"scripts/backfill/items/metadata",
	"metadata.json",
)

type metadata struct {
	OriginalID   string `json:"original_id"`
	OriginalSlug string `json:"original_slug"`
}

func main() {
	b, err := ioutil.ReadFile(metaPath)
	if err != nil {
		log.Fatalf("⚠️  Error reading metadata.json:\n\t > %v\n", err)
	}

	var meta map[string]metadata
	if err := json.Unmarshal(b, &meta); err != nil {
		log.Fatalf("⚠️  Error json decoding metadata.json:\n\t > %v\n", err)
	}

	backfill.Start("item.metadata", len(meta))
	defer backfill.Stop()

	for id, m := range meta {
		backfill.Request(itemproto.PUTUpdateItemRequest{
			Id: id,
			Metadata: map[string]string{
				"original_id": m.OriginalID,
				"original_slug": m.OriginalSlug,
			},
		}.Request(context.Background()))
	}
}
