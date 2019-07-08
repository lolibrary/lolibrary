package main

import (
	"context"

	"github.com/lolibrary/lolibrary/libraries/backfill"
	featureproto "github.com/lolibrary/lolibrary/service.feature/proto"
)

// COPY (select id, slug, name from features) TO STDOUT WITH CSV;
var featureQuery = `
5936cf73-5d8b-4545-832d-38ab93f80c13,lining,Lining
1b240899-2e1a-4531-b45a-525a5f97073c,no-shirring,No shirring
de015110-805c-4fc0-9c51-a9351f283895,corset-lacing,Corset lacing
58cb6890-7f3d-4f70-8841-222cef92dda0,long-sleeves,Long sleeves
acff794f-1c7b-4093-80df-2c06d674fce7,detachable-bow,Detachable bow
9e93eb7c-b819-4266-8553-7f752caada8a,partial-shirring,Partial shirring
e2cb096e-bd2e-402e-a597-02083ef2fc0c,side-zip,Side zip
ccfa8f72-797d-44a9-bd0b-f0281c888a34,detachable-waist-ties,Detachable waist ties
a1fa33c8-9708-41b9-bd89-e96b02283027,tiered-skirt,Tiered skirt
353a0f70-f93a-49f7-8b72-d6e35e2af92b,short-sleeves,Short sleeves
003de53c-fe4c-4096-b73c-49d4b2442605,high-waist,High waist
b8ef6a0d-3547-495c-a352-1a44fe77381d,pockets,Pockets
6a17ade6-173c-4849-9b65-bed17747f419,back-shirring,Back shirring
ef7c1c4e-fd71-4a25-9e93-1153f0a4b647,pintucks,Pintucks
2301684f-60cb-40f0-92f0-5d499f97452e,peter-pan-collar,Peter pan collar
c85d68f9-2fa3-4232-a347-0dc371f0caf5,full-shirring,Full shirring
b22223d7-6ad1-4d9a-98fe-8ea3e3b8a273,adjustable-straps,Adjustable straps
ff7f5968-db54-4439-b5a1-d39bd4eb0a95,detachable-trim,Detachable trim
e387d3a3-0748-480a-b428-948b50c52006,pleats,Pleats
25e00d3b-71c6-4243-8a86-3b393f40fb90,high-neck-collar,High neck collar
b96c126b-1569-404f-8af8-7d27d5afab05,neck-ties,Neck ties
f6c244a4-34f7-4c71-84e5-5f0fd86d51bf,bustled,Bustled
6a58020e-a6cc-4af7-ba89-77b839930460,empire-waist,Empire waist
1096e4aa-518d-4369-b218-2a9669f5bd94,removable-waist-ribbon,Removable waist ribbon
243afa26-fef3-446a-91b4-fe6ee5cd15d4,scalloped,Scalloped
d7155e56-dcae-435e-bb55-e68edb4920c0,detachable-sleeves,Detachable sleeves
0f106ab4-d72c-416a-8dfa-81ac414433db,dropped-waist,Dropped waist
215c6ed0-bb6d-407e-a01e-d7ce2f724dac,removable-collar,Removable collar
559ee9d9-b202-4b6c-8750-d61a1550659e,boning,Boning
4df859a5-5dcd-46ea-ba8c-e6f1f86392b1,tucks,Tucks
4dcabb42-8d90-4e62-b15b-160638718ce2,removable-belt,Removable belt
3de93969-d329-4260-a7f6-68b3cef55e33,princess-sleeves,Princess sleeves
ed0f6875-4ff8-426c-8e7a-66ba0d0ffb9f,halter-neckline,Halter neckline
46e9e476-979f-4a1a-ba7c-8421f2999970,jabot,Jabot
df360fe6-6ba8-42ec-b369-4de5bf96a712,capelet,Capelet
3157e7d1-d18d-4812-9509-39f2d106f913,removable-sash,Removable sash
bcf83fbb-1fc4-483a-9282-7cd33a1257cd,built-in-petticoat,Built-in petticoat
6b1c8cf3-89de-44d7-a2e7-9792900b8742,convertible-straps,Convertible straps
41d1d5b8-10d8-4689-8451-36672951973f,underbust,Underbust
6464d4a9-5655-4ff4-9fe5-d9b217aa49b9,detachable-apron,Detachable apron
82379390-f647-4a8b-84a6-871079705347,detachable-yoke,Detachable yoke
d4f4b763-e713-4790-b96a-44a178ff8023,sailor-collar,Sailor Collar
`

// in: ^([^,]+),([^,]+),(.*)$
// out: {\n\tId: "$1",\n\tSlug: "$2",\n\tName: "$3",\n},
var features = []*featureproto.POSTCreateFeatureRequest{
	{
		Id:   "5936cf73-5d8b-4545-832d-38ab93f80c13",
		Slug: "lining",
		Name: "Lining",
	},
	{
		Id:   "1b240899-2e1a-4531-b45a-525a5f97073c",
		Slug: "no-shirring",
		Name: "No shirring",
	},
	{
		Id:   "de015110-805c-4fc0-9c51-a9351f283895",
		Slug: "corset-lacing",
		Name: "Corset lacing",
	},
	{
		Id:   "58cb6890-7f3d-4f70-8841-222cef92dda0",
		Slug: "long-sleeves",
		Name: "Long sleeves",
	},
	{
		Id:   "acff794f-1c7b-4093-80df-2c06d674fce7",
		Slug: "detachable-bow",
		Name: "Detachable bow",
	},
	{
		Id:   "9e93eb7c-b819-4266-8553-7f752caada8a",
		Slug: "partial-shirring",
		Name: "Partial shirring",
	},
	{
		Id:   "e2cb096e-bd2e-402e-a597-02083ef2fc0c",
		Slug: "side-zip",
		Name: "Side zip",
	},
	{
		Id:   "ccfa8f72-797d-44a9-bd0b-f0281c888a34",
		Slug: "detachable-waist-ties",
		Name: "Detachable waist ties",
	},
	{
		Id:   "a1fa33c8-9708-41b9-bd89-e96b02283027",
		Slug: "tiered-skirt",
		Name: "Tiered skirt",
	},
	{
		Id:   "353a0f70-f93a-49f7-8b72-d6e35e2af92b",
		Slug: "short-sleeves",
		Name: "Short sleeves",
	},
	{
		Id:   "003de53c-fe4c-4096-b73c-49d4b2442605",
		Slug: "high-waist",
		Name: "High waist",
	},
	{
		Id:   "b8ef6a0d-3547-495c-a352-1a44fe77381d",
		Slug: "pockets",
		Name: "Pockets",
	},
	{
		Id:   "6a17ade6-173c-4849-9b65-bed17747f419",
		Slug: "back-shirring",
		Name: "Back shirring",
	},
	{
		Id:   "ef7c1c4e-fd71-4a25-9e93-1153f0a4b647",
		Slug: "pintucks",
		Name: "Pintucks",
	},
	{
		Id:   "2301684f-60cb-40f0-92f0-5d499f97452e",
		Slug: "peter-pan-collar",
		Name: "Peter pan collar",
	},
	{
		Id:   "c85d68f9-2fa3-4232-a347-0dc371f0caf5",
		Slug: "full-shirring",
		Name: "Full shirring",
	},
	{
		Id:   "b22223d7-6ad1-4d9a-98fe-8ea3e3b8a273",
		Slug: "adjustable-straps",
		Name: "Adjustable straps",
	},
	{
		Id:   "ff7f5968-db54-4439-b5a1-d39bd4eb0a95",
		Slug: "detachable-trim",
		Name: "Detachable trim",
	},
	{
		Id:   "e387d3a3-0748-480a-b428-948b50c52006",
		Slug: "pleats",
		Name: "Pleats",
	},
	{
		Id:   "25e00d3b-71c6-4243-8a86-3b393f40fb90",
		Slug: "high-neck-collar",
		Name: "High neck collar",
	},
	{
		Id:   "b96c126b-1569-404f-8af8-7d27d5afab05",
		Slug: "neck-ties",
		Name: "Neck ties",
	},
	{
		Id:   "f6c244a4-34f7-4c71-84e5-5f0fd86d51bf",
		Slug: "bustled",
		Name: "Bustled",
	},
	{
		Id:   "6a58020e-a6cc-4af7-ba89-77b839930460",
		Slug: "empire-waist",
		Name: "Empire waist",
	},
	{
		Id:   "1096e4aa-518d-4369-b218-2a9669f5bd94",
		Slug: "removable-waist-ribbon",
		Name: "Removable waist ribbon",
	},
	{
		Id:   "243afa26-fef3-446a-91b4-fe6ee5cd15d4",
		Slug: "scalloped",
		Name: "Scalloped",
	},
	{
		Id:   "d7155e56-dcae-435e-bb55-e68edb4920c0",
		Slug: "detachable-sleeves",
		Name: "Detachable sleeves",
	},
	{
		Id:   "0f106ab4-d72c-416a-8dfa-81ac414433db",
		Slug: "dropped-waist",
		Name: "Dropped waist",
	},
	{
		Id:   "215c6ed0-bb6d-407e-a01e-d7ce2f724dac",
		Slug: "removable-collar",
		Name: "Removable collar",
	},
	{
		Id:   "559ee9d9-b202-4b6c-8750-d61a1550659e",
		Slug: "boning",
		Name: "Boning",
	},
	{
		Id:   "4df859a5-5dcd-46ea-ba8c-e6f1f86392b1",
		Slug: "tucks",
		Name: "Tucks",
	},
	{
		Id:   "4dcabb42-8d90-4e62-b15b-160638718ce2",
		Slug: "removable-belt",
		Name: "Removable belt",
	},
	{
		Id:   "3de93969-d329-4260-a7f6-68b3cef55e33",
		Slug: "princess-sleeves",
		Name: "Princess sleeves",
	},
	{
		Id:   "ed0f6875-4ff8-426c-8e7a-66ba0d0ffb9f",
		Slug: "halter-neckline",
		Name: "Halter neckline",
	},
	{
		Id:   "46e9e476-979f-4a1a-ba7c-8421f2999970",
		Slug: "jabot",
		Name: "Jabot",
	},
	{
		Id:   "df360fe6-6ba8-42ec-b369-4de5bf96a712",
		Slug: "capelet",
		Name: "Capelet",
	},
	{
		Id:   "3157e7d1-d18d-4812-9509-39f2d106f913",
		Slug: "removable-sash",
		Name: "Removable sash",
	},
	{
		Id:   "bcf83fbb-1fc4-483a-9282-7cd33a1257cd",
		Slug: "built-in-petticoat",
		Name: "Built-in petticoat",
	},
	{
		Id:   "6b1c8cf3-89de-44d7-a2e7-9792900b8742",
		Slug: "convertible-straps",
		Name: "Convertible straps",
	},
	{
		Id:   "41d1d5b8-10d8-4689-8451-36672951973f",
		Slug: "underbust",
		Name: "Underbust",
	},
	{
		Id:   "6464d4a9-5655-4ff4-9fe5-d9b217aa49b9",
		Slug: "detachable-apron",
		Name: "Detachable apron",
	},
	{
		Id:   "82379390-f647-4a8b-84a6-871079705347",
		Slug: "detachable-yoke",
		Name: "Detachable yoke",
	},
	{
		Id:   "d4f4b763-e713-4790-b96a-44a178ff8023",
		Slug: "sailor-collar",
		Name: "Sailor Collar",
	},
}

func main() {
	ctx := context.Background()

	backfill.Start("feature", len(features))
	defer backfill.Stop()

	for _, feature := range features {
		backfill.Request(feature.Request(ctx))
	}
}
