package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gosuri/uiprogress"
	"github.com/lolibrary/lolibrary/libraries/api"
	"github.com/lolibrary/lolibrary/service.category/domain"
	categoryproto "github.com/lolibrary/lolibrary/service.category/proto"
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
	api.Run(backfill)

	fmt.Println("✅ All done!")
}

func backfill() {
	ctx := context.Background()
	uiprogress.Start()
	defer uiprogress.Stop()

	bar := uiprogress.AddBar(len(categories))
	bar.PrependElapsed()
	bar.AppendCompleted()

	for _, category := range categories {
		req := categoryproto.POSTCreateCategoryRequest{
			Slug: category.Slug,
			Name: category.Name,
		}
		_, err := req.Send(ctx).DecodeResponse()
		if err != nil {
			fmt.Printf("❌ Error: %v (%v)\n", err, req.Request(ctx).URL)
			os.Exit(1)
		}

		bar.Incr()
	}
}
