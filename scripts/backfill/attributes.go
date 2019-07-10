package main

import attributeproto "github.com/lolibrary/lolibrary/service.attribute/proto"

// COPY (select id, slug, name from attributes) TO STDOUT WITH CSV;
var attributesQuery = `
9ddb8f74-b103-43bb-aee4-aa3273a84818,bust,Bust
73e5168c-6aa7-497e-b8ea-c1f14ee71494,length,Length
0b5bb1ce-7818-434c-9396-336b7d548c5f,price,Price
f2ed336b-8ef5-4a07-833e-ac85f6bbd45f,waist,Waist
98f6e9e0-6515-44e3-88d8-18036c72e62e,owner-height,Owner Height
ca4d2b96-4fbe-484e-a478-c9e6e4cc6a05,owner-length,Owner Length
ce3358d7-05b4-4e8a-9a17-02039239760c,owner-waist,Owner Waist
788cf361-e119-4e15-bb66-bdc886bb93cd,cuff,Cuff
7fb26f02-fe4b-44d6-be1d-8906355738ef,shoulder-width,Shoulder Width
b78fa491-044c-42d7-b8f8-248fe9399d75,sleeve-length,Sleeve Length
13376ba0-6cb1-41a3-929c-0179d405b881,owner-notes,Owner Notes
9bf9205e-46b8-4b60-8b31-c734e35e0176,owner-bust,Owner Bust
697b2dad-8a60-4e58-8fbc-5e7ea39d7e6b,owner-underbust,Owner Underbust
22aa74f9-9cc2-4075-8123-5396606720e6,heel-height,Heel Height
7f6d1524-57c9-47d7-becc-ef3ca4ed3d2a,material,Material
2f2f4e54-b222-44a5-bba4-e86510afec38,soles,Soles
dfcbdd1e-b24a-4f2f-b4c3-19aff5092e8c,finishes,Finishes
be3d27e2-a5e5-4662-9691-52fa47cfc312,skirt-length,Skirt Length
dd0f0124-7cce-4dce-991a-4cf75ab8eebd,hip,Hip
b49203f7-eb3f-4179-8a4e-1bb94a805110,inseam,Inseam
e9fa9aaa-d5ec-4ff0-ae98-aba41b658517,thigh,Thigh
`

// in: ^([^,]+),([^,]+),(.*)$
// out: {\n\tId: "$1",\n\tSlug: "$2",\n\tName: "$3",\n},
var attributes = []*attributeproto.POSTCreateAttributeRequest{
	{
		Id:   "9ddb8f74-b103-43bb-aee4-aa3273a84818",
		Slug: "bust",
		Name: "Bust",
	},
	{
		Id:   "73e5168c-6aa7-497e-b8ea-c1f14ee71494",
		Slug: "length",
		Name: "Length",
	},
	{
		Id:   "0b5bb1ce-7818-434c-9396-336b7d548c5f",
		Slug: "price",
		Name: "Price",
	},
	{
		Id:   "f2ed336b-8ef5-4a07-833e-ac85f6bbd45f",
		Slug: "waist",
		Name: "Waist",
	},
	{
		Id:   "98f6e9e0-6515-44e3-88d8-18036c72e62e",
		Slug: "owner-height",
		Name: "Owner Height",
	},
	{
		Id:   "ca4d2b96-4fbe-484e-a478-c9e6e4cc6a05",
		Slug: "owner-length",
		Name: "Owner Length",
	},
	{
		Id:   "ce3358d7-05b4-4e8a-9a17-02039239760c",
		Slug: "owner-waist",
		Name: "Owner Waist",
	},
	{
		Id:   "788cf361-e119-4e15-bb66-bdc886bb93cd",
		Slug: "cuff",
		Name: "Cuff",
	},
	{
		Id:   "7fb26f02-fe4b-44d6-be1d-8906355738ef",
		Slug: "shoulder-width",
		Name: "Shoulder Width",
	},
	{
		Id:   "b78fa491-044c-42d7-b8f8-248fe9399d75",
		Slug: "sleeve-length",
		Name: "Sleeve Length",
	},
	{
		Id:   "13376ba0-6cb1-41a3-929c-0179d405b881",
		Slug: "owner-notes",
		Name: "Owner Notes",
	},
	{
		Id:   "9bf9205e-46b8-4b60-8b31-c734e35e0176",
		Slug: "owner-bust",
		Name: "Owner Bust",
	},
	{
		Id:   "697b2dad-8a60-4e58-8fbc-5e7ea39d7e6b",
		Slug: "owner-underbust",
		Name: "Owner Underbust",
	},
	{
		Id:   "22aa74f9-9cc2-4075-8123-5396606720e6",
		Slug: "heel-height",
		Name: "Heel Height",
	},
	{
		Id:   "7f6d1524-57c9-47d7-becc-ef3ca4ed3d2a",
		Slug: "material",
		Name: "Material",
	},
	{
		Id:   "2f2f4e54-b222-44a5-bba4-e86510afec38",
		Slug: "soles",
		Name: "Soles",
	},
	{
		Id:   "dfcbdd1e-b24a-4f2f-b4c3-19aff5092e8c",
		Slug: "finishes",
		Name: "Finishes",
	},
	{
		Id:   "be3d27e2-a5e5-4662-9691-52fa47cfc312",
		Slug: "skirt-length",
		Name: "Skirt Length",
	},
	{
		Id:   "dd0f0124-7cce-4dce-991a-4cf75ab8eebd",
		Slug: "hip",
		Name: "Hip",
	},
	{
		Id:   "b49203f7-eb3f-4179-8a4e-1bb94a805110",
		Slug: "inseam",
		Name: "Inseam",
	},
	{
		Id:   "e9fa9aaa-d5ec-4ff0-ae98-aba41b658517",
		Slug: "thigh",
		Name: "Thigh",
	},
}


