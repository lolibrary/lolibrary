package main

import (
	"context"
	"fmt"

	"github.com/gosuri/uiprogress"
	"github.com/logrusorgru/aurora"
	"github.com/lolibrary/lolibrary/libraries/filters"
	"github.com/lolibrary/lolibrary/libraries/portforward"
	"github.com/lolibrary/lolibrary/service.brand/domain"
	brandproto "github.com/lolibrary/lolibrary/service.brand/proto"
	"github.com/monzo/typhon"
)

var brands = []*domain.Brand{
	{
		Slug:      "angelic-pretty",
		ShortName: "ap",
		Name:      "Angelic Pretty",
	},
	{
		Slug:      "baby-the-stars-shine-bright",
		ShortName: "btssb",
		Name:      "Baby, the Stars Shine Bright",
	},
	{
		Slug:      "innocent-world",
		ShortName: "iw",
		Name:      "Innocent World",
	},
	{
		Slug:      "alice-and-the-pirates",
		ShortName: "aatp",
		Name:      "Alice and the Pirates",
	},
	{
		Slug:      "metamorphose-temps-de-fille",
		ShortName: "meta",
		Name:      "Metamorphose Temps de Fille",
	},
	{
		Slug:      "jane-marple",
		ShortName: "jane-marple",
		Name:      "Jane Marple",
	},
	{
		Slug:      "victorian-maiden",
		ShortName: "victorian-maiden",
		Name:      "Victorian Maiden",
	},
	{
		Slug:      "indie-brand",
		ShortName: "indie",
		Name:      "Indie Brand",
	},
	{
		Slug:      "atelier-boz",
		ShortName: "atelier-boz",
		Name:      "Atelier Boz",
	},
	{
		Slug:      "emily-shirley-temple-cute",
		ShortName: "emily-temple-cute",
		Name:      "Emily & Shirley Temple Cute",
	},
	{
		Slug:      "excentrique",
		ShortName: "excentrique",
		Name:      "Excentrique",
	},
	{
		Slug:      "bodyline",
		ShortName: "bodyline",
		Name:      "Bodyline",
	},
	{
		Slug:      "taobao",
		ShortName: "taobao",
		Name:      "TaoBao",
	},
	{
		Slug:      "moi-meme-moitie",
		ShortName: "moitie",
		Name:      "Moi-même-Moitié",
	},
	{
		Slug:      "juliette-et-justine",
		ShortName: "j-et-j",
		Name:      "Juliette et Justine",
	},
	{
		Slug:      "mary-magdalene",
		ShortName: "mary-magdalene",
		Name:      "Mary Magdalene",
	},
	{
		Slug:      "putumayo",
		ShortName: "putumayo",
		Name:      "Putumayo",
	},
	{
		Slug:      "h-naoto",
		ShortName: "h-naoto",
		Name:      "h. Naoto",
	},
	{
		Slug:      "antique-beast",
		ShortName: "antique-beast",
		Name:      "Antique Beast",
	},
	{
		Slug:      "atelier-pierrot",
		ShortName: "atelier-pierrot",
		Name:      "Atelier Pierrot",
	},
	{
		Slug:      "black-peace-now",
		ShortName: "bpn",
		Name:      "Black Peace Now",
	},
	{
		Slug:      "beth",
		ShortName: "beth",
		Name:      "Beth",
	},
	{
		Slug:      "maxicimam",
		ShortName: "max",
		Name:      "MAXICIMAM",
	},
	{
		Slug:      "cornet",
		ShortName: "cornet",
		Name:      "Cornet",
	},
	{
		Slug:      "grimoire",
		ShortName: "grimoire",
		Name:      "Grimoire",
	},
	{
		Slug:      "heart-e",
		ShortName: "heart-e",
		Name:      "Heart E",
	},
	{
		Slug:      "6dokidoki",
		ShortName: "dokidoki",
		Name:      "6%DOKIDOKI",
	},
	{
		Slug:      "offbrand",
		ShortName: "offbrand",
		Name:      "Offbrand",
	},
	{
		Slug:      "milk",
		ShortName: "milk",
		Name:      "MILK",
	},
	{
		Slug:      "physical-drop",
		ShortName: "physical-drop",
		Name:      "Physical Drop",
	},
	{
		Slug:      "millefleurs",
		ShortName: "millefleurs",
		Name:      "Millefleurs",
	},
	{
		Slug:      "pink-house",
		ShortName: "pink-house",
		Name:      "Pink House",
	},
	{
		Slug:      "vivienne-westwood",
		ShortName: "vivienne-westwood",
		Name:      "Vivienne Westwood",
	},
	{
		Slug:      "haenuli",
		ShortName: "haenuli",
		Name:      "Haenuli",
	},
	{
		Slug:      "chess-story",
		ShortName: "chess-story",
		Name:      "Chess Story",
	},
	{
		Slug:      "infanta",
		ShortName: "infanta",
		Name:      "Infanta",
	},
	{
		Slug:      "lief",
		ShortName: "lief",
		Name:      "Lief",
	},
	{
		Slug:      "surface-spell",
		ShortName: "surface-spell",
		Name:      "SurfaceSpell",
	},
	{
		Slug:      "chantilly",
		ShortName: "chantilly",
		Name:      "Enchantlic Enchantilly",
	},
	{
		Slug:      "krad-lanrete",
		ShortName: "krad-lanrete",
		Name:      "Krad Lanrete",
	},
	{
		Slug:      "sheglit",
		ShortName: "sheglit",
		Name:      "Sheglit",
	},
}

func main() {
	fmt.Printf("⚙️  Backfilling %v brand records via the API.\n", aurora.Green(len(brands)))

	ctx := context.Background()
	port, cmd := portforward.Enable()
	defer cmd.Close()

	client := typhon.Client.
		Filter(filters.EdgeProxyFilter(fmt.Sprintf("127.0.0.1:%d", port))).
		Filter(typhon.ErrorFilter)

	uiprogress.Start()
	defer uiprogress.Stop()

	bar := uiprogress.AddBar(len(brands))
	errors := 0

	for _, brand := range brands {
		_, err := brandproto.POSTCreateBrandRequest{
			Slug:        brand.Slug,
			ShortName:   brand.ShortName,
			Name:        brand.Name,
			Description: brand.Name,
		}.SendVia(ctx, client).DecodeResponse()
		if err != nil {
			errors++
		}

		bar.Incr()
	}

	if errors > 0 {
		fmt.Printf("⚠️  Completed with %v errors\n", aurora.Red(errors))
	} else {
		fmt.Println("✅ All done!")
	}
}
