package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/gammazero/workerpool"
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

	pool := workerpool.New(10)

	for id, m := range meta {
		pool.Submit(func () {
			id := id
			m := m

			//
		})
	}

}
