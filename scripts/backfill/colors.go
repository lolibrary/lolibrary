package main

import (
	"context"

	"github.com/lolibrary/lolibrary/libraries/backfill"
	colorproto "github.com/lolibrary/lolibrary/service.color/proto"
)

// COPY (select id, slug, name from colors) TO STDOUT WITH CSV;
var query = `
480112b1-5c69-494c-92fe-5074b078fd61,black,Black
01c5418e-b4a6-46e6-8b3b-5c208e1b5232,pink,Pink
38d13588-b9b5-40d0-8f16-24106e754156,white,White
6ea2597b-4c5b-4988-9695-b867b4f12014,offwhite,Offwhite
c78ab9b5-9b49-4d61-8040-e0b181179ca6,ivory,Ivory
81448b74-56a4-431f-a1bf-f763f73665fa,navy,Navy
ef1cf482-5ddc-4f15-8548-3df03fe7ee90,brown,Brown
e6a00b63-e173-48b4-ad20-dbbc0716eced,red,Red
82a10bf0-ed01-47bb-b426-02cb2528403e,winebordeaux,Wine/Bordeaux
3a547686-d0fb-483a-bfc6-ec6d6ac3df29,sax,Sax
ac470398-8ca1-4cf3-a12b-221783058fae,blue,Blue
78e79ff0-f34f-4a8a-8e59-42908e19e2d4,beige,Beige
fefee163-372c-4911-abb2-90369443025a,lavender,Lavender
358ad676-0490-4379-8e31-246113ff0f1d,green,Green
095dad29-3e3d-43e8-aeaf-dc9b4c70c1f3,mint,Mint
150c5a2d-096a-4aaa-8802-c240e6e0c0f6,black-x-white,Black x White
ea28e93d-00d4-44f2-81c6-ab0b3e716488,gray,Gray
a3aa88d6-9587-463f-9e28-3e4d0a9998dd,black-x-offwhite,Black x Offwhite
83458461-6784-43c6-b135-b30a5719b6d0,yellow,Yellow
9b78bfc6-f398-4e67-a716-a2587de5c2e9,purple,Purple
96db0661-4119-4c1b-87fd-106cb3a8c2ef,rose,Rose
d40726c6-3e74-47af-8367-c57cfae79416,pink-x-offwhite,Pink x Offwhite
455d08e3-b871-459f-8b11-33235735ac55,cream,Cream
cabbd55f-c407-412b-a953-18cca0eca6b3,gold,Gold
28770e8e-e92e-4014-9ed2-30c22ef2aa02,pink-x-white,Pink x White
3f1d0be8-9307-47ce-bd84-398357352dc2,white-x-pink,White x Pink
cf3e0df8-8ac8-4d81-ba74-d71aa87c0699,black-x-pink,Black x Pink
60b5ad9d-55f2-4ca2-bb13-73b24b65ffd7,black-x-red,Black x Red
293293d4-643e-4fff-8743-e65cf5ad769b,antique-gold,Antique Gold
db7d838f-47b5-4339-890e-e454088fc5f1,silver,Silver
048b81e4-743c-4648-b47a-ba9c77588c5a,white-x-black,White x Black
3adc3bb7-7b11-479f-a169-f52851ca6290,beige-x-brown,Beige x Brown
2ead53d0-fec4-455f-9ae5-438f72d896c4,sax-x-white,Sax x White
2decd39b-3924-4c67-8113-776f914c9c50,black-x-navy,Black x Navy
919f522e-c2bf-4f28-87c9-44a2411b2e68,dark-pink,Dark Pink
ab0d226a-4319-4ea6-881d-3f6254e1007b,antique-silver,Antique Silver
012d08f0-d5cf-4a3f-b3bc-8521065d1982,milk-tea,Milk tea
318ccd79-4cd4-4698-9f6c-32b28f61e45f,orange,Orange
967b8c2e-e7d3-4554-bd48-d1fbdff4511f,red-x-white,Red x White
79c9934d-b7f3-4643-b739-2f8a2d611766,black-x-beige,Black x Beige
d63ce2cc-c409-44aa-93f5-9bb4b4927681,brown-x-beige,Brown x Beige
db3c33ad-3c49-4841-9a94-5354c08b94ab,black-x-gray,Black x Gray
fa65a6ea-9f7d-42a3-bee2-b26f3a053bf8,black-x-silver,Black x Silver
b3653f71-cd10-4825-ac4a-4b057985f51c,red-x-offwhite,Red x Offwhite
2d2aef41-5720-41b4-9418-9b85559243d9,black-x-gold,Black x Gold
d89ee8a5-a037-41fa-92c6-ce074118361a,offwhite-x-black,Offwhite x Black
92ddd18a-0f3a-4d26-a926-49a6ef3611d1,black-x-blue,Black x Blue
138c9b1a-a7d3-4649-b3d8-6805737ae7e9,navy-x-offwhite,Navy x Offwhite
f62fe35a-a012-411e-9ed1-58679eaea2f4,sax-x-offwhite,Sax x Offwhite
59e71c69-868d-48cf-b837-25988a97d14e,brown-x-pink,Brown x Pink
8d059289-3793-4ec2-b556-5a128a5c1b64,black-x-wine,Black x Wine
db1e19ab-e75e-41a9-bcc2-8dcf4b9a680f,black-x-purple,Black x Purple
8bb5b45d-8ad3-4efa-89e3-894ef0ca143a,pink-x-sax,Pink x Sax
cc3d71d0-6644-4349-822a-1a45246c9658,navy-x-white,Navy x White
f6023d08-cec1-4642-b7c7-c6bad4d94697,lavender-x-white,Lavender x White
1c974af8-c49c-4f4c-9a3c-b8580fe377b5,olive,Olive
e6c0e6e4-77c2-43e3-b4b1-43d0033de5a7,gray-x-black,Gray x Black
f9bdb0cd-4eef-403f-b5ba-19db6c32ef08,ivory-x-brown,Ivory x Brown
dd261a43-6f7a-43de-a6cb-51cdeb8b8ebb,red-x-pink,Red x Pink
69758125-19d0-4692-a2c4-28f138ebeb3a,white-x-red,White x Red
2e3356a7-9a57-4bd4-897a-302a307e2eef,winebordeaux-x-offwhite,Wine/Bordeaux x Offwhite
a157f14e-1920-4612-ad6a-18cce23422d0,pink-x-black,Pink x Black
a1655996-08a3-4cb7-a420-224e95ca1389,white-x-navy,White x Navy
aea7288f-aad1-4533-95bf-21afd0de462a,blue-x-white,Blue x White
2ca0ed10-a5a7-4fa0-a17a-e558a64ce9c7,black-x-ivory,Black x Ivory
2c34bdf0-415e-4dd6-8f11-a634d011fc40,offwhite-x-navy,Offwhite x Navy
b8189edc-77a1-4b86-a1aa-0ca6ae8bb5bd,mint-x-white,Mint x White
3efe65f9-eb60-49fa-9c1a-bfc4e38508fe,white-x-gray,White x Gray
b9a1f231-2af2-4aff-8273-988c420eea25,black-x-green,Black x Green
756725e3-f9e0-4ce6-8448-09b7adde2b9b,greenmint-x-brown,Green/Mint x Brown
165c5f0f-e522-4960-8aa2-5f5572d5653b,navy-x-black,Navy x Black
5fba2c5d-da13-4e9e-92a6-aee3636129ff,pink-gold,Pink Gold
4ec0e82f-d600-4cc3-acd3-9499f4769159,ivory-x-black,Ivory x Black
43169e5b-fd99-4f80-8d58-aceb77d72c44,wine-x-offwhite,Wine x Offwhite
4bd56904-9a85-4115-927e-f401d3cb992b,black-x-plum,Black x Plum
a1c2ea36-02df-4045-9666-3e1b17b7daa5,white-x-green,White x Green
0a4e7498-d182-43ef-9a6d-67fecc8ba2e0,lavender-x-black,Lavender x Black
9804e495-074b-4bea-92dd-b1775546a311,offwhitecream,Offwhite/Cream
`

// in: ^([^,]+),([^,]+),(.*)$
// out: {\n\tId: "$1",\n\tSlug: "$2",\n\tName: "$3",\n},
var colors = []*colorproto.POSTCreateColorRequest{
	{
		Id:   "480112b1-5c69-494c-92fe-5074b078fd61",
		Slug: "black",
		Name: "Black",
	},
	{
		Id:   "01c5418e-b4a6-46e6-8b3b-5c208e1b5232",
		Slug: "pink",
		Name: "Pink",
	},
	{
		Id:   "38d13588-b9b5-40d0-8f16-24106e754156",
		Slug: "white",
		Name: "White",
	},
	{
		Id:   "6ea2597b-4c5b-4988-9695-b867b4f12014",
		Slug: "offwhite",
		Name: "Offwhite",
	},
	{
		Id:   "c78ab9b5-9b49-4d61-8040-e0b181179ca6",
		Slug: "ivory",
		Name: "Ivory",
	},
	{
		Id:   "81448b74-56a4-431f-a1bf-f763f73665fa",
		Slug: "navy",
		Name: "Navy",
	},
	{
		Id:   "ef1cf482-5ddc-4f15-8548-3df03fe7ee90",
		Slug: "brown",
		Name: "Brown",
	},
	{
		Id:   "e6a00b63-e173-48b4-ad20-dbbc0716eced",
		Slug: "red",
		Name: "Red",
	},
	{
		Id:   "82a10bf0-ed01-47bb-b426-02cb2528403e",
		Slug: "winebordeaux",
		Name: "Wine/Bordeaux",
	},
	{
		Id:   "3a547686-d0fb-483a-bfc6-ec6d6ac3df29",
		Slug: "sax",
		Name: "Sax",
	},
	{
		Id:   "ac470398-8ca1-4cf3-a12b-221783058fae",
		Slug: "blue",
		Name: "Blue",
	},
	{
		Id:   "78e79ff0-f34f-4a8a-8e59-42908e19e2d4",
		Slug: "beige",
		Name: "Beige",
	},
	{
		Id:   "fefee163-372c-4911-abb2-90369443025a",
		Slug: "lavender",
		Name: "Lavender",
	},
	{
		Id:   "358ad676-0490-4379-8e31-246113ff0f1d",
		Slug: "green",
		Name: "Green",
	},
	{
		Id:   "095dad29-3e3d-43e8-aeaf-dc9b4c70c1f3",
		Slug: "mint",
		Name: "Mint",
	},
	{
		Id:   "150c5a2d-096a-4aaa-8802-c240e6e0c0f6",
		Slug: "black-x-white",
		Name: "Black x White",
	},
	{
		Id:   "ea28e93d-00d4-44f2-81c6-ab0b3e716488",
		Slug: "gray",
		Name: "Gray",
	},
	{
		Id:   "a3aa88d6-9587-463f-9e28-3e4d0a9998dd",
		Slug: "black-x-offwhite",
		Name: "Black x Offwhite",
	},
	{
		Id:   "83458461-6784-43c6-b135-b30a5719b6d0",
		Slug: "yellow",
		Name: "Yellow",
	},
	{
		Id:   "9b78bfc6-f398-4e67-a716-a2587de5c2e9",
		Slug: "purple",
		Name: "Purple",
	},
	{
		Id:   "96db0661-4119-4c1b-87fd-106cb3a8c2ef",
		Slug: "rose",
		Name: "Rose",
	},
	{
		Id:   "d40726c6-3e74-47af-8367-c57cfae79416",
		Slug: "pink-x-offwhite",
		Name: "Pink x Offwhite",
	},
	{
		Id:   "455d08e3-b871-459f-8b11-33235735ac55",
		Slug: "cream",
		Name: "Cream",
	},
	{
		Id:   "cabbd55f-c407-412b-a953-18cca0eca6b3",
		Slug: "gold",
		Name: "Gold",
	},
	{
		Id:   "28770e8e-e92e-4014-9ed2-30c22ef2aa02",
		Slug: "pink-x-white",
		Name: "Pink x White",
	},
	{
		Id:   "3f1d0be8-9307-47ce-bd84-398357352dc2",
		Slug: "white-x-pink",
		Name: "White x Pink",
	},
	{
		Id:   "cf3e0df8-8ac8-4d81-ba74-d71aa87c0699",
		Slug: "black-x-pink",
		Name: "Black x Pink",
	},
	{
		Id:   "60b5ad9d-55f2-4ca2-bb13-73b24b65ffd7",
		Slug: "black-x-red",
		Name: "Black x Red",
	},
	{
		Id:   "293293d4-643e-4fff-8743-e65cf5ad769b",
		Slug: "antique-gold",
		Name: "Antique Gold",
	},
	{
		Id:   "db7d838f-47b5-4339-890e-e454088fc5f1",
		Slug: "silver",
		Name: "Silver",
	},
	{
		Id:   "048b81e4-743c-4648-b47a-ba9c77588c5a",
		Slug: "white-x-black",
		Name: "White x Black",
	},
	{
		Id:   "3adc3bb7-7b11-479f-a169-f52851ca6290",
		Slug: "beige-x-brown",
		Name: "Beige x Brown",
	},
	{
		Id:   "2ead53d0-fec4-455f-9ae5-438f72d896c4",
		Slug: "sax-x-white",
		Name: "Sax x White",
	},
	{
		Id:   "2decd39b-3924-4c67-8113-776f914c9c50",
		Slug: "black-x-navy",
		Name: "Black x Navy",
	},
	{
		Id:   "919f522e-c2bf-4f28-87c9-44a2411b2e68",
		Slug: "dark-pink",
		Name: "Dark Pink",
	},
	{
		Id:   "ab0d226a-4319-4ea6-881d-3f6254e1007b",
		Slug: "antique-silver",
		Name: "Antique Silver",
	},
	{
		Id:   "012d08f0-d5cf-4a3f-b3bc-8521065d1982",
		Slug: "milk-tea",
		Name: "Milk tea",
	},
	{
		Id:   "318ccd79-4cd4-4698-9f6c-32b28f61e45f",
		Slug: "orange",
		Name: "Orange",
	},
	{
		Id:   "967b8c2e-e7d3-4554-bd48-d1fbdff4511f",
		Slug: "red-x-white",
		Name: "Red x White",
	},
	{
		Id:   "79c9934d-b7f3-4643-b739-2f8a2d611766",
		Slug: "black-x-beige",
		Name: "Black x Beige",
	},
	{
		Id:   "d63ce2cc-c409-44aa-93f5-9bb4b4927681",
		Slug: "brown-x-beige",
		Name: "Brown x Beige",
	},
	{
		Id:   "db3c33ad-3c49-4841-9a94-5354c08b94ab",
		Slug: "black-x-gray",
		Name: "Black x Gray",
	},
	{
		Id:   "fa65a6ea-9f7d-42a3-bee2-b26f3a053bf8",
		Slug: "black-x-silver",
		Name: "Black x Silver",
	},
	{
		Id:   "b3653f71-cd10-4825-ac4a-4b057985f51c",
		Slug: "red-x-offwhite",
		Name: "Red x Offwhite",
	},
	{
		Id:   "2d2aef41-5720-41b4-9418-9b85559243d9",
		Slug: "black-x-gold",
		Name: "Black x Gold",
	},
	{
		Id:   "d89ee8a5-a037-41fa-92c6-ce074118361a",
		Slug: "offwhite-x-black",
		Name: "Offwhite x Black",
	},
	{
		Id:   "92ddd18a-0f3a-4d26-a926-49a6ef3611d1",
		Slug: "black-x-blue",
		Name: "Black x Blue",
	},
	{
		Id:   "138c9b1a-a7d3-4649-b3d8-6805737ae7e9",
		Slug: "navy-x-offwhite",
		Name: "Navy x Offwhite",
	},
	{
		Id:   "f62fe35a-a012-411e-9ed1-58679eaea2f4",
		Slug: "sax-x-offwhite",
		Name: "Sax x Offwhite",
	},
	{
		Id:   "59e71c69-868d-48cf-b837-25988a97d14e",
		Slug: "brown-x-pink",
		Name: "Brown x Pink",
	},
	{
		Id:   "8d059289-3793-4ec2-b556-5a128a5c1b64",
		Slug: "black-x-wine",
		Name: "Black x Wine",
	},
	{
		Id:   "db1e19ab-e75e-41a9-bcc2-8dcf4b9a680f",
		Slug: "black-x-purple",
		Name: "Black x Purple",
	},
	{
		Id:   "8bb5b45d-8ad3-4efa-89e3-894ef0ca143a",
		Slug: "pink-x-sax",
		Name: "Pink x Sax",
	},
	{
		Id:   "cc3d71d0-6644-4349-822a-1a45246c9658",
		Slug: "navy-x-white",
		Name: "Navy x White",
	},
	{
		Id:   "f6023d08-cec1-4642-b7c7-c6bad4d94697",
		Slug: "lavender-x-white",
		Name: "Lavender x White",
	},
	{
		Id:   "1c974af8-c49c-4f4c-9a3c-b8580fe377b5",
		Slug: "olive",
		Name: "Olive",
	},
	{
		Id:   "e6c0e6e4-77c2-43e3-b4b1-43d0033de5a7",
		Slug: "gray-x-black",
		Name: "Gray x Black",
	},
	{
		Id:   "f9bdb0cd-4eef-403f-b5ba-19db6c32ef08",
		Slug: "ivory-x-brown",
		Name: "Ivory x Brown",
	},
	{
		Id:   "dd261a43-6f7a-43de-a6cb-51cdeb8b8ebb",
		Slug: "red-x-pink",
		Name: "Red x Pink",
	},
	{
		Id:   "69758125-19d0-4692-a2c4-28f138ebeb3a",
		Slug: "white-x-red",
		Name: "White x Red",
	},
	{
		Id:   "2e3356a7-9a57-4bd4-897a-302a307e2eef",
		Slug: "winebordeaux-x-offwhite",
		Name: "Wine/Bordeaux x Offwhite",
	},
	{
		Id:   "a157f14e-1920-4612-ad6a-18cce23422d0",
		Slug: "pink-x-black",
		Name: "Pink x Black",
	},
	{
		Id:   "a1655996-08a3-4cb7-a420-224e95ca1389",
		Slug: "white-x-navy",
		Name: "White x Navy",
	},
	{
		Id:   "aea7288f-aad1-4533-95bf-21afd0de462a",
		Slug: "blue-x-white",
		Name: "Blue x White",
	},
	{
		Id:   "2ca0ed10-a5a7-4fa0-a17a-e558a64ce9c7",
		Slug: "black-x-ivory",
		Name: "Black x Ivory",
	},
	{
		Id:   "2c34bdf0-415e-4dd6-8f11-a634d011fc40",
		Slug: "offwhite-x-navy",
		Name: "Offwhite x Navy",
	},
	{
		Id:   "b8189edc-77a1-4b86-a1aa-0ca6ae8bb5bd",
		Slug: "mint-x-white",
		Name: "Mint x White",
	},
	{
		Id:   "3efe65f9-eb60-49fa-9c1a-bfc4e38508fe",
		Slug: "white-x-gray",
		Name: "White x Gray",
	},
	{
		Id:   "b9a1f231-2af2-4aff-8273-988c420eea25",
		Slug: "black-x-green",
		Name: "Black x Green",
	},
	{
		Id:   "756725e3-f9e0-4ce6-8448-09b7adde2b9b",
		Slug: "greenmint-x-brown",
		Name: "Green/Mint x Brown",
	},
	{
		Id:   "165c5f0f-e522-4960-8aa2-5f5572d5653b",
		Slug: "navy-x-black",
		Name: "Navy x Black",
	},
	{
		Id:   "5fba2c5d-da13-4e9e-92a6-aee3636129ff",
		Slug: "pink-gold",
		Name: "Pink Gold",
	},
	{
		Id:   "4ec0e82f-d600-4cc3-acd3-9499f4769159",
		Slug: "ivory-x-black",
		Name: "Ivory x Black",
	},
	{
		Id:   "43169e5b-fd99-4f80-8d58-aceb77d72c44",
		Slug: "wine-x-offwhite",
		Name: "Wine x Offwhite",
	},
	{
		Id:   "4bd56904-9a85-4115-927e-f401d3cb992b",
		Slug: "black-x-plum",
		Name: "Black x Plum",
	},
	{
		Id:   "a1c2ea36-02df-4045-9666-3e1b17b7daa5",
		Slug: "white-x-green",
		Name: "White x Green",
	},
	{
		Id:   "0a4e7498-d182-43ef-9a6d-67fecc8ba2e0",
		Slug: "lavender-x-black",
		Name: "Lavender x Black",
	},
	{
		Id:   "9804e495-074b-4bea-92dd-b1775546a311",
		Slug: "offwhitecream",
		Name: "Offwhite/Cream",
	},
}

func main() {
	ctx := context.Background()

	backfill.Start("color", len(colors))
	defer backfill.Stop()

	for _, color := range colors {
		backfill.Request(color.Request(ctx))
	}
}
