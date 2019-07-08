package main

import (
	"context"

	"github.com/lolibrary/lolibrary/libraries/backfill"
	categoryproto "github.com/lolibrary/lolibrary/service.category/proto"
)

// COPY (select id, slug, name from categories) TO STDOUT WITH CSV;
var categoryQuery = `
3da3a1c0-2d58-4a69-be39-666c7c4646ab,jsk,JSK
835a990c-9a7c-4a14-9fd0-893647111e43,hair-accessories,Hair Accessories
146cd101-9c8b-4999-b564-acc1e2915bb4,skirt,Skirt
0b1919d0-64a3-46f7-8ff4-b94f4da41131,sets,Sets
f93b4b74-9cac-4710-bf99-1947f491f035,op,OP
9c4c1c5f-a312-472f-8364-3b3df212a8dd,blouse,Blouse
c8936a0d-d2ca-4149-a98f-1c2516f82927,socks,Socks
820c8f2d-ce76-4cb9-b2d5-82f24640fe60,jewelry,Jewelry
5287e746-4b9b-44ca-9c6f-09efee95173a,bags,Bags
7a4348c9-e489-4a28-9590-221de863ec44,coats,Coats
ca5f50b3-3555-4957-8b9c-a0e4e5084066,bolero,Bolero
f583db67-75ed-4665-a7ef-3b08ee85d73b,pants,Pants
d3764f5a-d63e-4fdc-944a-4c56e563c1ca,bloomers,Bloomers / Undergarments
90f54b22-9016-4892-87ff-05d1f9e67807,salopette,Salopette
e5537f89-c107-4792-ba3a-d50dcc9352f1,unmentionables,Unmentionables
59fa782d-1a79-4d25-8557-1a6301c04981,other,Other
697d4051-16cb-420e-ac5a-f63729a06fcd,cardigan,Cardigan
9c32c4e7-299c-48b4-b530-c2c22f696888,accessories,Accessories
f95079db-6f4e-4678-8c22-1b07405672d2,corsetbustier,Corset / Bustier
a9eacacd-a626-494f-8940-ff10ea22eb0c,cape,Cape
6fa1c61e-ee2a-4e61-b8ce-ff78b72e5c59,vest,Vest
46ccb1ff-b607-4492-9a28-cddd778b319c,petticoat,Petticoat
0ee298f2-3fcc-4a0a-b580-163cef338d06,parasols,Parasols
bde69ada-c7a8-4069-875e-ef73f379e1d7,cutsew,Cutsew
36b221b0-aa63-469b-ac99-bcd17dca527c,shoes,Shoes
`

// in: ^([^,]+),([^,]+),(.*)$
// out: {\n\tId: "$1",\n\tSlug: "$2",\n\tName: "$3",\n},
var categories = []*categoryproto.POSTCreateCategoryRequest{
	{
		Id:   "3da3a1c0-2d58-4a69-be39-666c7c4646ab",
		Slug: "jsk",
		Name: "JSK",
	},
	{
		Id:   "835a990c-9a7c-4a14-9fd0-893647111e43",
		Slug: "hair-accessories",
		Name: "Hair Accessories",
	},
	{
		Id:   "146cd101-9c8b-4999-b564-acc1e2915bb4",
		Slug: "skirt",
		Name: "Skirt",
	},
	{
		Id:   "0b1919d0-64a3-46f7-8ff4-b94f4da41131",
		Slug: "sets",
		Name: "Sets",
	},
	{
		Id:   "f93b4b74-9cac-4710-bf99-1947f491f035",
		Slug: "op",
		Name: "OP",
	},
	{
		Id:   "9c4c1c5f-a312-472f-8364-3b3df212a8dd",
		Slug: "blouse",
		Name: "Blouse",
	},
	{
		Id:   "c8936a0d-d2ca-4149-a98f-1c2516f82927",
		Slug: "socks",
		Name: "Socks",
	},
	{
		Id:   "820c8f2d-ce76-4cb9-b2d5-82f24640fe60",
		Slug: "jewelry",
		Name: "Jewelry",
	},
	{
		Id:   "5287e746-4b9b-44ca-9c6f-09efee95173a",
		Slug: "bags",
		Name: "Bags",
	},
	{
		Id:   "7a4348c9-e489-4a28-9590-221de863ec44",
		Slug: "coats",
		Name: "Coats",
	},
	{
		Id:   "ca5f50b3-3555-4957-8b9c-a0e4e5084066",
		Slug: "bolero",
		Name: "Bolero",
	},
	{
		Id:   "f583db67-75ed-4665-a7ef-3b08ee85d73b",
		Slug: "pants",
		Name: "Pants",
	},
	{
		Id:   "d3764f5a-d63e-4fdc-944a-4c56e563c1ca",
		Slug: "bloomers",
		Name: "Bloomers / Undergarments",
	},
	{
		Id:   "90f54b22-9016-4892-87ff-05d1f9e67807",
		Slug: "salopette",
		Name: "Salopette",
	},
	{
		Id:   "e5537f89-c107-4792-ba3a-d50dcc9352f1",
		Slug: "unmentionables",
		Name: "Unmentionables",
	},
	{
		Id:   "59fa782d-1a79-4d25-8557-1a6301c04981",
		Slug: "other",
		Name: "Other",
	},
	{
		Id:   "697d4051-16cb-420e-ac5a-f63729a06fcd",
		Slug: "cardigan",
		Name: "Cardigan",
	},
	{
		Id:   "9c32c4e7-299c-48b4-b530-c2c22f696888",
		Slug: "accessories",
		Name: "Accessories",
	},
	{
		Id:   "f95079db-6f4e-4678-8c22-1b07405672d2",
		Slug: "corsetbustier",
		Name: "Corset / Bustier",
	},
	{
		Id:   "a9eacacd-a626-494f-8940-ff10ea22eb0c",
		Slug: "cape",
		Name: "Cape",
	},
	{
		Id:   "6fa1c61e-ee2a-4e61-b8ce-ff78b72e5c59",
		Slug: "vest",
		Name: "Vest",
	},
	{
		Id:   "46ccb1ff-b607-4492-9a28-cddd778b319c",
		Slug: "petticoat",
		Name: "Petticoat",
	},
	{
		Id:   "0ee298f2-3fcc-4a0a-b580-163cef338d06",
		Slug: "parasols",
		Name: "Parasols",
	},
	{
		Id:   "bde69ada-c7a8-4069-875e-ef73f379e1d7",
		Slug: "cutsew",
		Name: "Cutsew",
	},
	{
		Id:   "36b221b0-aa63-469b-ac99-bcd17dca527c",
		Slug: "shoes",
		Name: "Shoes",
	},
}

func main() {
	ctx := context.Background()

	backfill.Start("color", len(categories))
	defer backfill.Stop()

	for _, category := range categories {
		backfill.Request(category.Request(ctx))
	}
}
