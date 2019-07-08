package main

import (
	"context"

	"github.com/lolibrary/lolibrary/libraries/backfill"
	brandproto "github.com/lolibrary/lolibrary/service.brand/proto"
)

// COPY (select id, slug, short_name, name, image_id from brands) TO STDOUT WITH CSV;
var brandQuery = `
ecb829b3-8a9e-4348-8945-b985ede7f03a,angelic-pretty,ap,Angelic Pretty,87f6ed35-8cf6-460b-92ac-7b57d1c946ad
0e29621e-5cd5-4814-957c-903808652c6c,baby-the-stars-shine-bright,btssb,"Baby, the Stars Shine Bright",34a7ef76-38a9-4f7f-88d3-6c4db23b638e
1d849c42-a991-48e5-a1cc-dde246782651,innocent-world,iw,Innocent World,4240e5dd-95f9-4617-8967-f00386a369a4
d6f7698e-bdf2-4b91-825b-4ebabbca75d0,alice-and-the-pirates,aatp,Alice and the Pirates,854b1455-a84d-4517-8a16-9c079fac586e
8c0a5cd7-15a2-4c36-a59a-19062c1ceb9b,metamorphose-temps-de-fille,meta,Metamorphose Temps de Fille,3596a82d-c6d0-4cc3-9337-4bd124ff4f5b
33378176-432a-4676-8f99-4e1167de9bdf,jane-marple,jane-marple,Jane Marple,7b23cbc3-91c4-448d-95f6-19cd970a791e
ef5cf9ed-b62e-47fb-a2ed-ad25c3c35d17,victorian-maiden,victorian-maiden,Victorian Maiden,d0bb7b25-0831-4386-9f72-014bcc428151
b24938af-de30-4f81-8b57-c2192ca6c507,indie-brand,indie,Indie Brand,1ef4feab-23bb-49cc-9a1b-e8396dc73883
f60edf85-4466-4369-ae4a-753073f11b58,atelier-boz,atelier-boz,Atelier Boz,7834a3d2-4eb7-44e5-a706-b2d3deb4b861
ef82cf8e-38af-4ddc-8cf3-f43b1c3b29a1,emily-shirley-temple-cute,emily-temple-cute,Emily & Shirley Temple Cute,ab1211a7-067e-4f5e-aabb-1066ac889853
bf16c13c-1775-47c5-8d95-0c57895de258,excentrique,excentrique,Excentrique,6f39056c-1ec9-4b7d-9469-e5956bbe3cd7
9c9dab7d-cbe4-4817-8f0b-8f3a2001cac2,bodyline,bodyline,Bodyline,95270616-aac0-44ca-a7e4-749c1988264d
be220d8b-bffc-42a4-a073-9c886e6fe79f,taobao,taobao,TaoBao,53fcfec3-bd65-4804-8404-9801cf226c81
53da5b54-735a-4912-ae27-b1a0fc1659ff,moi-meme-moitie,moitie,Moi-même-Moitié,876a5104-b4fd-48b8-a918-8dada69adc76
19cd58ae-7ea2-41f1-a145-9caf83cde401,juliette-et-justine,j-et-j,Juliette et Justine,82821d32-4966-4ab2-8f02-40ab8dc6f5d4
4da985ce-4d52-40ad-ba7f-18d6650926bc,mary-magdalene,mary-magdalene,Mary Magdalene,4f86b676-b548-4b98-9efa-e5474e978749
8a4bfb01-81e7-4b19-994f-48b40716ea93,putumayo,putumayo,Putumayo,4c9ce683-bd05-48a9-addd-33dfb66e9529
e8c88299-fa80-468a-a719-e86501237c88,h-naoto,h-naoto,h. Naoto,c158f4d0-3c6c-49ab-ae23-1f76166fd77e
0751c2d7-bc5b-4ec2-a69a-315cf8305162,antique-beast,antique-beast,Antique Beast,f27d6af5-b9c9-4b13-9e79-e3546f900672
1205a7b6-c82b-4744-baaf-4465be8317d7,atelier-pierrot,atelier-pierrot,Atelier Pierrot,70c741c2-89a4-4342-8cf7-e33a37610615
6ed4636f-561b-4161-ba78-14e58ca83ca2,black-peace-now,bpn,Black Peace Now,bb4f11af-196c-44f2-9d92-0dcbcdc3ff15
642e97e5-2dd8-4a5f-bd54-c6ce7c846053,beth,beth,Beth,0b745363-d6f7-4e49-9abc-16f4ad6f7334
6b0cb2ca-8f53-4320-b2d7-232c3e333497,maxicimam,max,MAXICIMAM,c75657d7-1f25-4f3b-b75b-82d5fd0208c2
f454192f-93b8-4635-9409-b066c8d24c69,cornet,cornet,Cornet,170f7ced-5556-4722-96b7-7c4b3f790556
66bc41c7-11f8-4dd2-a43c-8a0312d12a3f,grimoire,grimoire,Grimoire,a49df53d-b098-435d-bde5-dd670b943bdb
5f422ee1-11b8-4593-a6d3-9eef471ae10e,heart-e,heart-e,Heart E,9e4acf29-72cb-4992-83fd-92849e5b3460
8025f54f-05a3-4ea5-bfa9-08dc4a4957f8,6dokidoki,dokidoki,6%DOKIdOKI,44c038da-87d3-4f2e-8c43-4e1db8543921
1117e0da-4f1c-4300-a266-f78fcb4770cd,offbrand,offbrand,Offbrand,701671bb-cbd7-46d5-95b5-3439bdfe3fc3
dd82e2a2-8ed8-444f-a75f-85174eacfe71,milk,milk,MILK,ca50d686-e9cd-441f-a92c-88e973d3e5b2
d6702e87-d724-4b2f-9b2d-3a454d0f0dc3,physical-drop,physical-drop,Physical Drop,e82f33b5-5d7c-45be-bfbb-81bc85a61232
bc40549b-db09-4c03-8461-e5b37766009e,millefleurs,millefleurs,Millefleurs,d47cf504-17a0-4059-8207-98df30c8f4af
68d9a41d-0297-4f01-96be-b2c6d2b359a1,pink-house,pink-house,Pink House,eeb21c64-1043-4154-9017-71e97d0aed36
d7f9dfb9-49a1-475c-825a-6d49d7fe6af7,vivienne-westwood,vivienne-westwood,Vivienne Westwood,2ece1a48-dc4e-435c-a6dd-3b2198a95aec
5aaa0ed0-7203-4ead-965c-46009ded4225,haenuli,haenuli,Haenuli,c2a31b55-4ff8-4e8a-9d32-d52b28d958a4
7b0096c0-59f8-4e05-a80c-a7701639072d,chess-story,chess-story,Chess Story,0841e332-e908-4f2d-bdba-b35bf6f643ec
03eaf969-7a92-4696-8e3c-0957a1b6d56c,infanta,infanta,Infanta,00c7fd51-e548-4488-8a68-1ba26ca76172
7c476701-e8de-4c28-83bc-e70faef88256,lief,lief,Lief,51adf68d-b630-4d52-a315-1d8495da1d11
db054a75-a12c-49c3-b97d-7589f238b31e,surface-spell,surface-spell,SurfaceSpell,b1498af4-6d7b-4ff9-a9d9-03b0678b34f4
00d02505-12c9-4785-94f7-a57eacd70fd0,chantilly,chantilly,Enchantlic Enchantilly,758170ad-6d99-4b06-8557-60dc7a002f85
5e5d054c-ddc1-49f4-8603-e8477a16049d,krad-lanrete,krad-lanrete,Krad Lanrete,a61a2c83-68de-4e0e-8a7f-54b69432624d
3d36fee9-b2e6-4eae-a486-6f89b78306e4,sheglit,sheglit,Sheglit,d356e02d-f680-46b8-abb2-f731520d7aef
`

// in: ^([^,]+),([^,]+),([a-z\-]+),"?([a-zA-Z, ]+|[^,]+)"?,([^,]+)$
// out: {\n\tId: "$1",\n\tSlug: "$2",\n\tShortName: "$3",\n\tName: "$4",\n\tImageId: "$5",\n},
var brands = []*brandproto.POSTCreateBrandRequest{
	{
		Id:        "ecb829b3-8a9e-4348-8945-b985ede7f03a",
		Slug:      "angelic-pretty",
		ShortName: "ap",
		Name:      "Angelic Pretty",
		ImageId:   "87f6ed35-8cf6-460b-92ac-7b57d1c946ad",
	},
	{
		Id:        "0e29621e-5cd5-4814-957c-903808652c6c",
		Slug:      "baby-the-stars-shine-bright",
		ShortName: "btssb",
		Name:      "Baby, the Stars Shine Bright",
		ImageId:   "34a7ef76-38a9-4f7f-88d3-6c4db23b638e",
	},
	{
		Id:        "1d849c42-a991-48e5-a1cc-dde246782651",
		Slug:      "innocent-world",
		ShortName: "iw",
		Name:      "Innocent World",
		ImageId:   "4240e5dd-95f9-4617-8967-f00386a369a4",
	},
	{
		Id:        "d6f7698e-bdf2-4b91-825b-4ebabbca75d0",
		Slug:      "alice-and-the-pirates",
		ShortName: "aatp",
		Name:      "Alice and the Pirates",
		ImageId:   "854b1455-a84d-4517-8a16-9c079fac586e",
	},
	{
		Id:        "8c0a5cd7-15a2-4c36-a59a-19062c1ceb9b",
		Slug:      "metamorphose-temps-de-fille",
		ShortName: "meta",
		Name:      "Metamorphose Temps de Fille",
		ImageId:   "3596a82d-c6d0-4cc3-9337-4bd124ff4f5b",
	},
	{
		Id:        "33378176-432a-4676-8f99-4e1167de9bdf",
		Slug:      "jane-marple",
		ShortName: "jane-marple",
		Name:      "Jane Marple",
		ImageId:   "7b23cbc3-91c4-448d-95f6-19cd970a791e",
	},
	{
		Id:        "ef5cf9ed-b62e-47fb-a2ed-ad25c3c35d17",
		Slug:      "victorian-maiden",
		ShortName: "victorian-maiden",
		Name:      "Victorian Maiden",
		ImageId:   "d0bb7b25-0831-4386-9f72-014bcc428151",
	},
	{
		Id:        "b24938af-de30-4f81-8b57-c2192ca6c507",
		Slug:      "indie-brand",
		ShortName: "indie",
		Name:      "Indie Brand",
		ImageId:   "1ef4feab-23bb-49cc-9a1b-e8396dc73883",
	},
	{
		Id:        "f60edf85-4466-4369-ae4a-753073f11b58",
		Slug:      "atelier-boz",
		ShortName: "atelier-boz",
		Name:      "Atelier Boz",
		ImageId:   "7834a3d2-4eb7-44e5-a706-b2d3deb4b861",
	},
	{
		Id:        "ef82cf8e-38af-4ddc-8cf3-f43b1c3b29a1",
		Slug:      "emily-shirley-temple-cute",
		ShortName: "emily-temple-cute",
		Name:      "Emily & Shirley Temple Cute",
		ImageId:   "ab1211a7-067e-4f5e-aabb-1066ac889853",
	},
	{
		Id:        "bf16c13c-1775-47c5-8d95-0c57895de258",
		Slug:      "excentrique",
		ShortName: "excentrique",
		Name:      "Excentrique",
		ImageId:   "6f39056c-1ec9-4b7d-9469-e5956bbe3cd7",
	},
	{
		Id:        "9c9dab7d-cbe4-4817-8f0b-8f3a2001cac2",
		Slug:      "bodyline",
		ShortName: "bodyline",
		Name:      "Bodyline",
		ImageId:   "95270616-aac0-44ca-a7e4-749c1988264d",
	},
	{
		Id:        "be220d8b-bffc-42a4-a073-9c886e6fe79f",
		Slug:      "taobao",
		ShortName: "taobao",
		Name:      "TaoBao",
		ImageId:   "53fcfec3-bd65-4804-8404-9801cf226c81",
	},
	{
		Id:        "53da5b54-735a-4912-ae27-b1a0fc1659ff",
		Slug:      "moi-meme-moitie",
		ShortName: "moitie",
		Name:      "Moi-même-Moitié",
		ImageId:   "876a5104-b4fd-48b8-a918-8dada69adc76",
	},
	{
		Id:        "19cd58ae-7ea2-41f1-a145-9caf83cde401",
		Slug:      "juliette-et-justine",
		ShortName: "j-et-j",
		Name:      "Juliette et Justine",
		ImageId:   "82821d32-4966-4ab2-8f02-40ab8dc6f5d4",
	},
	{
		Id:        "4da985ce-4d52-40ad-ba7f-18d6650926bc",
		Slug:      "mary-magdalene",
		ShortName: "mary-magdalene",
		Name:      "Mary Magdalene",
		ImageId:   "4f86b676-b548-4b98-9efa-e5474e978749",
	},
	{
		Id:        "8a4bfb01-81e7-4b19-994f-48b40716ea93",
		Slug:      "putumayo",
		ShortName: "putumayo",
		Name:      "Putumayo",
		ImageId:   "4c9ce683-bd05-48a9-addd-33dfb66e9529",
	},
	{
		Id:        "e8c88299-fa80-468a-a719-e86501237c88",
		Slug:      "h-naoto",
		ShortName: "h-naoto",
		Name:      "h. Naoto",
		ImageId:   "c158f4d0-3c6c-49ab-ae23-1f76166fd77e",
	},
	{
		Id:        "0751c2d7-bc5b-4ec2-a69a-315cf8305162",
		Slug:      "antique-beast",
		ShortName: "antique-beast",
		Name:      "Antique Beast",
		ImageId:   "f27d6af5-b9c9-4b13-9e79-e3546f900672",
	},
	{
		Id:        "1205a7b6-c82b-4744-baaf-4465be8317d7",
		Slug:      "atelier-pierrot",
		ShortName: "atelier-pierrot",
		Name:      "Atelier Pierrot",
		ImageId:   "70c741c2-89a4-4342-8cf7-e33a37610615",
	},
	{
		Id:        "6ed4636f-561b-4161-ba78-14e58ca83ca2",
		Slug:      "black-peace-now",
		ShortName: "bpn",
		Name:      "Black Peace Now",
		ImageId:   "bb4f11af-196c-44f2-9d92-0dcbcdc3ff15",
	},
	{
		Id:        "642e97e5-2dd8-4a5f-bd54-c6ce7c846053",
		Slug:      "beth",
		ShortName: "beth",
		Name:      "Beth",
		ImageId:   "0b745363-d6f7-4e49-9abc-16f4ad6f7334",
	},
	{
		Id:        "6b0cb2ca-8f53-4320-b2d7-232c3e333497",
		Slug:      "maxicimam",
		ShortName: "max",
		Name:      "MAXICIMAM",
		ImageId:   "c75657d7-1f25-4f3b-b75b-82d5fd0208c2",
	},
	{
		Id:        "f454192f-93b8-4635-9409-b066c8d24c69",
		Slug:      "cornet",
		ShortName: "cornet",
		Name:      "Cornet",
		ImageId:   "170f7ced-5556-4722-96b7-7c4b3f790556",
	},
	{
		Id:        "66bc41c7-11f8-4dd2-a43c-8a0312d12a3f",
		Slug:      "grimoire",
		ShortName: "grimoire",
		Name:      "Grimoire",
		ImageId:   "a49df53d-b098-435d-bde5-dd670b943bdb",
	},
	{
		Id:        "5f422ee1-11b8-4593-a6d3-9eef471ae10e",
		Slug:      "heart-e",
		ShortName: "heart-e",
		Name:      "Heart E",
		ImageId:   "9e4acf29-72cb-4992-83fd-92849e5b3460",
	},
	{
		Id:        "8025f54f-05a3-4ea5-bfa9-08dc4a4957f8",
		Slug:      "6dokidoki",
		ShortName: "dokidoki",
		Name:      "6%DOKIdOKI",
		ImageId:   "44c038da-87d3-4f2e-8c43-4e1db8543921",
	},
	{
		Id:        "1117e0da-4f1c-4300-a266-f78fcb4770cd",
		Slug:      "offbrand",
		ShortName: "offbrand",
		Name:      "Offbrand",
		ImageId:   "701671bb-cbd7-46d5-95b5-3439bdfe3fc3",
	},
	{
		Id:        "dd82e2a2-8ed8-444f-a75f-85174eacfe71",
		Slug:      "milk",
		ShortName: "milk",
		Name:      "MILK",
		ImageId:   "ca50d686-e9cd-441f-a92c-88e973d3e5b2",
	},
	{
		Id:        "d6702e87-d724-4b2f-9b2d-3a454d0f0dc3",
		Slug:      "physical-drop",
		ShortName: "physical-drop",
		Name:      "Physical Drop",
		ImageId:   "e82f33b5-5d7c-45be-bfbb-81bc85a61232",
	},
	{
		Id:        "bc40549b-db09-4c03-8461-e5b37766009e",
		Slug:      "millefleurs",
		ShortName: "millefleurs",
		Name:      "Millefleurs",
		ImageId:   "d47cf504-17a0-4059-8207-98df30c8f4af",
	},
	{
		Id:        "68d9a41d-0297-4f01-96be-b2c6d2b359a1",
		Slug:      "pink-house",
		ShortName: "pink-house",
		Name:      "Pink House",
		ImageId:   "eeb21c64-1043-4154-9017-71e97d0aed36",
	},
	{
		Id:        "d7f9dfb9-49a1-475c-825a-6d49d7fe6af7",
		Slug:      "vivienne-westwood",
		ShortName: "vivienne-westwood",
		Name:      "Vivienne Westwood",
		ImageId:   "2ece1a48-dc4e-435c-a6dd-3b2198a95aec",
	},
	{
		Id:        "5aaa0ed0-7203-4ead-965c-46009ded4225",
		Slug:      "haenuli",
		ShortName: "haenuli",
		Name:      "Haenuli",
		ImageId:   "c2a31b55-4ff8-4e8a-9d32-d52b28d958a4",
	},
	{
		Id:        "7b0096c0-59f8-4e05-a80c-a7701639072d",
		Slug:      "chess-story",
		ShortName: "chess-story",
		Name:      "Chess Story",
		ImageId:   "0841e332-e908-4f2d-bdba-b35bf6f643ec",
	},
	{
		Id:        "03eaf969-7a92-4696-8e3c-0957a1b6d56c",
		Slug:      "infanta",
		ShortName: "infanta",
		Name:      "Infanta",
		ImageId:   "00c7fd51-e548-4488-8a68-1ba26ca76172",
	},
	{
		Id:        "7c476701-e8de-4c28-83bc-e70faef88256",
		Slug:      "lief",
		ShortName: "lief",
		Name:      "Lief",
		ImageId:   "51adf68d-b630-4d52-a315-1d8495da1d11",
	},
	{
		Id:        "db054a75-a12c-49c3-b97d-7589f238b31e",
		Slug:      "surface-spell",
		ShortName: "surface-spell",
		Name:      "SurfaceSpell",
		ImageId:   "b1498af4-6d7b-4ff9-a9d9-03b0678b34f4",
	},
	{
		Id:        "00d02505-12c9-4785-94f7-a57eacd70fd0",
		Slug:      "chantilly",
		ShortName: "chantilly",
		Name:      "Enchantlic Enchantilly",
		ImageId:   "758170ad-6d99-4b06-8557-60dc7a002f85",
	},
	{
		Id:        "5e5d054c-ddc1-49f4-8603-e8477a16049d",
		Slug:      "krad-lanrete",
		ShortName: "krad-lanrete",
		Name:      "Krad Lanrete",
		ImageId:   "a61a2c83-68de-4e0e-8a7f-54b69432624d",
	},
	{
		Id:        "3d36fee9-b2e6-4eae-a486-6f89b78306e4",
		Slug:      "sheglit",
		ShortName: "sheglit",
		Name:      "Sheglit",
		ImageId:   "d356e02d-f680-46b8-abb2-f731520d7aef",
	},
}

func main() {
	ctx := context.Background()

	backfill.Start("brand", len(brands))
	defer backfill.Stop()

	for _, brand := range brands {
		backfill.Request(brand.Request(ctx))
	}
}
