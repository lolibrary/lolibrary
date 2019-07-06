package main

import (
	"context"
	"fmt"

	"github.com/gosuri/uiprogress"
	"github.com/logrusorgru/aurora"
	"github.com/lolibrary/lolibrary/libraries/filters"
	"github.com/lolibrary/lolibrary/libraries/portforward"
	"github.com/lolibrary/lolibrary/service.category/domain"
	categoryproto "github.com/lolibrary/lolibrary/service.category/proto"
	"github.com/monzo/typhon"
)

/*
 jsk              | JSK
 hair-accessories | Hair Accessories
 skirt            | Skirt
 sets             | Sets
 op               | OP
 blouse           | Blouse
 socks            | Socks
 jewelry          | Jewelry
 bags             | Bags
 coats            | Coats
 bolero           | Bolero
 pants            | Pants
 bloomers         | Bloomers / Undergarments
 salopette        | Salopette
 unmentionables   | Unmentionables
 other            | Other
 cardigan         | Cardigan
 accessories      | Accessories
 corsetbustier    | Corset/Bustier
 cape             | Cape
 vest             | Vest
 petticoat        | Petticoat
 parasols         | Parasols
 cutsew           | Cutsew
 shoes            | Shoes
*/
var categories = []*domain.Category{
	{Slug: "jsk", Name: "JSK"},
	{Slug: "hair-accessories", Name: "Hair Accessories"},
	{Slug: "skirt", Name: "Skirt"},
	{Slug: "sets", Name: "Sets"},
	{Slug: "op", Name: "OP"},
	{Slug: "blouse", Name: "Blouse"},
	{Slug: "socks", Name: "Socks"},
	{Slug: "jewelry", Name: "Jewelry"},
	{Slug: "bags", Name: "Bags"},
	{Slug: "coats", Name: "Coats"},
	{Slug: "bolero", Name: "Bolero"},
	{Slug: "pants", Name: "Pants"},
	{Slug: "bloomers", Name: "Bloomers / Undergarments"},
	{Slug: "salopette", Name: "Salopette"},
	{Slug: "unmentionables", Name: "Unmentionables"},
	{Slug: "other", Name: "Other"},
	{Slug: "cardigan", Name: "Cardigan"},
	{Slug: "accessories", Name: "Accessories"},
	{Slug: "corsetbustier", Name: "Corset / Bustier"},
	{Slug: "cape", Name: "Cape"},
	{Slug: "vest", Name: "Vest"},
	{Slug: "petticoat", Name: "Petticoat"},
	{Slug: "parasols", Name: "Parasols"},
	{Slug: "cutsew", Name: "Cutsew"},
	{Slug: "shoes", Name: "Shoes"},
}

func main() {
	fmt.Printf("⚙️  Backfilling %v Category records via the API.\n", aurora.Green(len(categories)))

	ctx := context.Background()
	port, cmd := portforward.Enable()
	defer cmd.Close()

	client := typhon.Client.
		Filter(filters.EdgeProxyFilter(fmt.Sprintf("127.0.0.1:%d", port))).
		Filter(typhon.ErrorFilter)

	uiprogress.Start()
	defer uiprogress.Stop()

	bar := uiprogress.AddBar(len(categories))
	bar.AppendCompleted()

	errors := 0

	for _, category := range categories {
		_, err := categoryproto.POSTCreateCategoryRequest{
			Slug: category.Slug,
			Name: category.Name,
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
