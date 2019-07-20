// This backfills metadata for items.
// The sources to fill are:
//  - original_id from the very first version of lolibrary
//  - original_slug, from the very first version of lolibrary (the /apparel/* links)
//
// Metadata is in an inverted index and so you can search it very effectively.
package links

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	basePath        = path.Join(os.Getenv("GOPATH"), "src/github.com/lolibrary/lolibrary/scripts/backfill/items")
	nodeLinkPath    = path.Join(basePath, "node_links.json")
	apparelLinkPath = path.Join(basePath, "apparel_links.json")
	metadataPath    = path.Join(basePath, "metadata.json")
)

type nodeLink struct {
	ID     string `json:"id"`
	NodeID string `json:"node_id"`
}

type apparelLink struct {
	ID           string `json:"id"`
	OriginalSlug string `json:"original_slug"`
}

type meta struct {
	OriginalID   string `json:"original_id"`
	OriginalSlug string `json:"original_slug"`
}

var items = make(map[string]meta)

func main() {
	// first, open files.
	nl, err := ioutil.ReadFile(nodeLinkPath)
	if err != nil {
		log.Fatalf("⚠️  Error opening node_links.json:\n\t > %v\n", err)
	}

	al, err := ioutil.ReadFile(apparelLinkPath)
	if err != nil {
		log.Fatalf("⚠️  Error opening apparel_links.json:\n\t > %v\n", err)
	}

	// combine them
	var nodeLinks []nodeLink
	if err := json.Unmarshal(nl, &nodeLinks); err != nil {
		log.Fatalf("⚠️  Error json decoding node_links.json:\n\t > %v\n", err)
	}

	var apparelLinks []apparelLink
	if err := json.Unmarshal(al, &apparelLinks); err != nil {
		log.Fatalf("⚠️  Error json decoding apparel_links.json:\n\t > %v\n", err)
	}

	for _, link := range nodeLinks {
		items[link.ID] = meta{OriginalID: link.NodeID}
	}

	for _, link := range apparelLinks {
		if i, ok := items[link.ID]; ok {
			items[link.ID] = meta{
				OriginalID: i.OriginalID,
				OriginalSlug: link.OriginalSlug,
			}
		} else {
			// missing item from node_id.
			log.Printf("🚨 Missing item ID: %s (slug: %s)\n", link.ID, link.OriginalSlug)
			items[link.ID] = meta{OriginalSlug: link.OriginalSlug}
		}
	}

	// now generate a single json file called metadata.json, as a one-item-per-line deal.
	b, err := json.Marshal(items)
	if err != nil {
		log.Fatalf("⚠️  Error json encoding metadata.json:\n\t > %v\n", err)
	}

	err = ioutil.WriteFile(metadataPath, b, os.FileMode(os.O_RDWR | os.O_TRUNC))
	if err != nil {
		log.Fatalf("⚠️  Error writing metadata.json\n\t > %v\n", err)
	}

	fmt.Printf("✅ Wrote metadata.json")
}
