// This backfills metadata for items.
// The sources to fill are:
//  - original_id from the very first version of lolibrary
//  - original_slug, from the very first version of lolibrary (the /apparel/* links)
//
// Metadata is in an inverted index and so you can search it very effectively.
package main

import (
	"os"
	"path"
)

var (
	basePath        = path.Join(os.Getenv("GOPATH"), "src/github.com/lolibrary/lolibrary/scripts/backfill/items")
	nodeLinkPath    = path.Join(basePath, "node_links.json")
	apparelLinkPath = path.Join(basePath, "apparel_links.json")
)

func main() {

}
