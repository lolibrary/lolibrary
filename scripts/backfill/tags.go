package main

import "github.com/lolibrary/lolibrary/service.tag/domain"

/*
COPY (select id, slug, name from tags) TO STDOUT WITH CSV;

60667b47-f661-4db0-ab56-74becad8d567,stripes,stripes
5287901a-111a-49a0-b3dd-a32052a8eef3,roses,Roses
3af89989-28c4-4e70-a08c-d4d8de2826c7,gardens,Gardens
e1448e60-ce4d-4291-83e7-65c8b154e7e3,architectural,Architectural
f500e424-9b4e-4a43-bb9e-3f84f76df7f9,chairs,Chairs
8371d116-6102-4855-95c0-75ccd2bafc85,swans,Swans
ea72fb7f-f55f-410a-a659-580d731b59af,bows,Bows
c115f273-3ddd-432b-b036-9d3c3823c3a4,perfumes,Perfumes
df261473-2a70-4ed5-8dd7-91452051eff4,bottles,Bottles
f8cdc971-1c61-4ac8-a185-2742420a71ff,pearls,Pearls
677ac1d2-3bed-40fd-8bb5-74dfbce18f88,jewelry,Jewelry
70c0964e-2f5c-4b79-9dea-b36fde2ceea1,writing,Writing
a9126e03-66ea-4b48-9b4e-187526246502,sweets,Sweets
98189ad6-9575-4c0f-a0d8-386e7b91fa8a,stars,Stars
1ff3901d-1ece-4e79-b35c-067450b7bb74,hearts,Hearts
cd1f0d47-08f4-4a32-9273-1ef117aaecf4,desserts,Desserts
bdf7a0ab-8d53-491d-a7cf-5760da8e7546,cakes,Cakes
3ec8747a-fee5-4ec8-811a-5f5c8f406c61,food,Food
ce42a0bd-7fe6-4042-9142-33179cca6a3e,crowns,Crowns
a244589a-1a2d-4977-86b3-c163600d05d7,coaches-carriages,Coaches Carriages
79a5d8b7-83ec-452e-b86b-de09b28af4ba,castles,Castles
148491eb-4c1d-4c84-9220-25d7c6a2a6c3,swan-lake,Swan Lake
200b47c8-3fb0-4cf4-8bda-f9ceeab03300,fairy-tales,Fairy Tales
b5834ea3-ae9c-4c73-ad86-5cde023c37ad,florals,Florals
8026e2b8-825f-4cf5-b061-7e9095c2c6fb,jacquard,jacquard
7a2f96ca-d48e-4b6c-bc29-15a42d1154a1,partial,partial
963f404f-02ac-4ca7-9a67-d3c354e1aec4,regimental-stripes,regimental stripes
fa468a33-4710-42d5-aabb-bb2b56a0bd35,alice-in-wonderland,Alice in Wonderland
247fe0c3-96a4-45cf-ab9b-b72e74a786d0,trumps,Trumps
409ac7eb-9b4a-455f-a21d-e659d22229d8,clocks,Clocks
9d6dd0a2-4895-441b-a15f-6d46f79948dd,rabbits,Rabbits
9bca6dec-3592-480f-b1c2-7e9ef90d2dd3,strawberries,Strawberries
09c8a2c1-36ac-4050-b4f8-9de25dde2114,polka-dots,polka dots
ee207d84-d31c-41fd-8f2a-d969aa561124,crosses,Crosses
62a1429a-38e1-4cda-b1fd-714c453a6eac,frames,Frames
30d5df5e-2306-445b-ada3-50000431ba72,figures,Figures
9b15d05a-acab-49e8-9eb6-1388c47bb6bf,incomplete-colorways,Incomplete-colorways
bc777600-18a9-4ec1-b935-fcdc843d71b5,butterflies,Butterflies
a02b26f4-7e97-4416-b6e2-4aa0f1975fb7,thorns,Thorns
e1cd8467-2bdc-42fb-81a9-5ce9508f7eca,windows,Windows
4d56a435-3e90-4bfc-aeed-ec9bbf7ca9e4,churches,Churches
4778ff65-9521-4f3c-9c79-17f896cf5e4d,lace,lace
e967aea5-ecc9-4661-ba28-c53ced6eb086,velveteen,velveteen
0c742330-c88e-41da-ad2f-4a8c260f57dd,gates,Gates
21c2f99b-c3a4-4524-9f6c-590444596104,chiffon,chiffon
10cb852f-faec-489d-8fae-50023abca73a,chandeliers,Chandeliers
833619de-fabc-409b-bf3d-fd986cccb52e,candles,Candles
37510362-82aa-42d4-882c-e9c6cc200d92,cathedrals,Cathedrals
2b00ee01-6459-41c6-ab59-6aec76a485ca,ornaments,Ornaments
94b2de49-b081-4dc9-acc3-f10ef8b6bf42,christmas,Christmas
b9a11a81-6e75-4664-bad8-74d7d7350dd9,abstract-decorative,Abstract Decorative
1ffe0b29-33e8-4ae8-8bef-0f9f1170b381,plaids,plaids
75f5b041-0f50-4f36-ba88-75e53255b8fa,fruits,Fruits
6e8844be-b360-49bf-81ed-1284f82ae109,cherries,Cherries
d275af53-605d-4f78-8c23-cc9b88a44fea,embroidery,embroidery
13c39cd7-9b84-49c6-b17f-ef8c2f9a9247,solid,solid
ab99b523-5a8b-464a-ac88-040f21c4ae62,snow-white,Snow White
582aaf5e-0656-4059-967f-0b6c73752a59,apples,Apples
ff5b6c81-5d78-4306-a554-799fa02744f7,deer,Deer
9dbf01d4-ebab-4ba2-916b-37cd9b13ad43,animals,Animals
164f79f5-d2ec-4f5a-aed1-a0be99a85795,forest,Forest
c2f00d40-f2cf-4ae7-81f2-4f1b2c7eab97,trees,Trees
d86a7b2f-2457-4018-a5ac-fecad9d1a209,mirrors,Mirrors
ec3d8362-6dfc-4c21-a161-6f3d5fda75fc,white-rabbit,White Rabbit
7130552e-4b7d-4fb7-b2a6-981a74ac172b,headbands,headbands
d3a21074-e48d-4443-acd5-e11cfbafdf96,cats,Cats
bf2736c4-0056-4517-82d8-cc05d2c2e4a8,bunnybears,Bunnybears
034a7d78-2c42-441a-8eb9-b9eb2a66919e,furniture,Furniture
94d96cf4-bd36-4614-b091-2f7de813aaed,glitter,glitter
6bddda82-3da0-4671-952c-da5d8b6dcbe2,daisies,Daisies
72e592f8-bcc0-4b6f-a36e-c71e97f26c5b,bouquets,Bouquets
959a0135-9d1a-4ed8-aa0f-8326dd261690,imai-kira,Imai Kira
c8dfbb9e-aac2-4a7d-94d4-23553d600819,birds,Birds
b0aa776c-eda7-482f-8d48-2da2db150e3d,birdcages,Birdcages
cc925d7c-707f-4d34-8a2f-dcca88622ea9,cookies,Cookies
04fd0dcd-e706-4b83-aab5-eb5500deba4a,applique,applique
29b84e71-56f0-4d95-8edd-948e545847e4,playing-cards,Playing Cards
85f32a89-7fea-4d07-b60b-dcf53f045578,letters,Letters
c738f3a4-ac46-4b9e-b4f9-89501fdd269f,rickrack,rickrack
2c1ab46d-4ab3-41c8-b5d5-afdc4749f622,pinstripes,pinstripes
42ea103a-ec7e-4fc9-8e94-cca96fc652df,maid,maid
f3d80995-3dd3-460d-ad39-1a042c7e82d5,sailor,sailor
b4b7051e-a22e-4012-9c20-1252515004c9,marine,Marine
d5fb88f6-8ee0-4719-9682-911f1aa1e9c2,pom-pom,pom-pom
6492f5b1-e815-43c0-a3d7-d1897bb84cd3,grass,Grass
31fc1b44-0788-48f6-9edc-f1826f7bbce0,clouds,Clouds
6aab9ef6-e7e1-4457-b819-3f0a6fc5a165,animal-ears,animal ears
14f4fd24-cf61-454e-95ea-61d9a8d76432,collaboration,collaboration
10e8bc10-e26a-478a-b7f9-263fb286ce3a,whipped-cream,Whipped Cream
4f05986b-e3ef-4ee2-ad27-b255883e4b64,ice-creams,Ice Creams
bb2b6a63-aa52-4235-bd8f-21e332d0994e,knitted,knitted
9e86b439-c4f5-495a-8b7f-954af61d13db,bears,Bears
3010ad34-21c5-4df0-a4a5-cb6e1eee7a07,candy,Candy
13f05c97-ddce-4a9c-bd48-fa1877f2b852,beds,Beds
adf4de04-f73d-4e59-884b-54b0661a7f01,rhinestones,rhinestones
54f6e0dd-bc60-4e25-bacc-4a1bbd3b87d1,treasures,Treasures
5c634064-b597-4c1b-8b18-ccbf67aee6f5,pirates,Pirates
cae850b4-811f-4c21-9d87-222c3d7fb80f,musical-notes,Musical Notes
82d98b84-7ca4-49d7-b09c-39cbe93707cb,jewels,Jewels
e26cea64-2b6f-4a67-bb06-08177a6db0a9,keys-non-piano,Keys (non-piano)
93042fe7-b712-4174-b2a7-ff345e0e4aa6,skulls,Skulls
27b9a3b1-46c7-4439-bf86-21f6e7f18f0e,apron,apron
ab69d91c-67f6-469f-967e-4ce21f5ae939,pocket-watches,Pocket Watches
b3d0457b-ba5b-495a-8f26-afae76e9c7e8,anchors,Anchors
5ceb8a7f-7d98-4197-ad07-1a4edaee6e7a,heraldry,Heraldry
d4f0390a-595c-4870-aa7a-feee080d4a85,diamonds,Diamonds
d806236c-3854-46bc-a474-7ac99aeb8a61,instruments,Instruments
84c1d5c8-7ba8-4f64-8aa1-c2cb95649471,staircases,Staircases
54b3df74-269f-40ba-adfc-dc8c19aca53a,angels,Angels
fbc9d213-e837-4c45-9ebc-1d9cc334a398,stained-glass,Stained Glass
774c4752-36d5-4bda-a358-de2bdfcd552c,circus,Circus
758fd1d9-4248-4446-9448-1096ba8757c9,balloons,Balloons
2535508b-c990-4666-a7b3-da6f85829566,fleur-de-lis,Fleur-de-lis
a65f1189-e16c-43b7-8d10-fa1ddcda6be2,wings,Wings
54d0c06d-faeb-4b24-b0ee-74d04d7d0a5b,witches-magic,Witches Magic
d88902b0-9164-4595-a15f-ffdce28f8c0f,moons,Moons
a2fbc8ea-5831-4fe1-921a-c15230a87683,theatre,Theatre
190a6975-e382-4ad1-b20a-92c3d60c8007,poodles,Poodles
8a545c55-0d30-4124-8c8c-bf9ef8c7dc76,ponies,Ponies
2af2b193-c7b4-4781-bdf7-d1ef9f5ee194,carousels,Carousels
5b36abc1-61d0-4153-b7f3-8f91936897c1,replica,Replica
09997f04-4899-4074-9f9d-031466302f89,mushrooms,Mushrooms
5e4dcb60-545e-4191-bd57-24a685552618,pianos,Pianos
9175810c-90fe-4acb-8b85-42d064b8e4f5,music,Music
f531e387-e8d0-4f33-84de-3d562a6732f5,headdress-hair-accessories-not-a-tag,Headdress Hair Accessories (Not A Tag)
655a77a7-0d31-4fcb-808a-ec1bc3391175,dogs,Dogs
761f2b8a-af3e-4634-82fe-58d7d9bdaf21,spiderwebs,Spiderwebs
8e43623d-26fe-4dc2-ba13-ff0200b0fb98,gobelin,gobelin
b26835ea-c268-47b4-801b-0171ef111da2,chains,Chains
b7a6c010-6e2b-4478-aad6-60a13c7174b5,bats,Bats
8b30ccf3-b11d-4067-a877-cc9b1b6ab482,military,military
c807ad8e-86e3-4352-8c46-63303a2c2983,faux-fur,faux fur
77a18fde-ee4c-4ebd-a84d-9cab54f69526,hats,Hats
558f4b51-2094-47f0-a558-4415fbc79264,wrist-cuffs,wrist cuffs
b471d546-4626-42f4-9e99-351fcf54595a,camo,camo
4df05feb-9688-4c7b-9974-0d40ac2bfc15,houndstooth,houndstooth
2a678ffe-0422-4f30-ba53-0dde1fbc7528,spiders,Spiders
d12916ca-4ba4-4b37-9e6a-5e4f9527fb5d,argyle,argyle
97726e25-dd5e-4dff-a852-4f599497db61,pansies,Pansies
aa666862-3dc3-48bc-929c-666ee9889285,crown-label,Crown Label
f4292a67-e8e6-4c5a-83d8-5ad559761713,tea,Tea
38990061-0827-45a9-89d3-ef949c89ae50,cups,Cups
ae6d2068-4e7b-40c1-a10f-2a73b1289613,violins,Violins
254fd547-57bd-4624-948a-e8a97f92422d,squirrels,Squirrels
7e541d9c-ae73-4351-bc0b-d194dc9f5d62,sheep,Sheep
dc8f3636-63e9-4c2b-b9f0-99b58b4f9331,lions,Lions
45db235f-2761-4755-9cd8-ea632b1f79ce,lacy,lacy
7b072ca1-3556-49c7-b118-cc62bdbfc8cc,plants,Plants
216ab08e-d139-410d-9f14-cc6b7027e06b,dolls,Dolls
209fa323-d81f-4a14-862d-b181e2df9411,bonnets,bonnets
8cb3486a-3c32-4cd5-b48a-ebc1f047d40a,sleeping-beauty,Sleeping Beauty
a85026f7-a04f-40fd-a578-95bd475e539f,cupcakes,Cupcakes
f025198a-76b9-43d6-9fce-1b0607e490b4,feathers,Feathers
9430230c-f046-435a-bf03-732773251928,cameos,Cameos
cc84e20d-75cd-4b81-87ee-c08d9b5c3ff8,clubs,Clubs
93444f7a-f41e-43e6-9c31-e04e75a9369c,spades,Spades
904872ab-3d8c-404c-aa97-d0019945b6fe,numbers,Numbers
f3d6dae7-fa1f-4652-ac96-e41604e9819a,horses,Horses
724b2379-3daa-4cd2-a7ac-2f2290861d78,organdy,organdy
9e2ede1e-6847-4bd4-b2b8-2e0fef3acdc8,animal-print,animal print
0fe23d8f-74bf-46d6-b8d8-26c3a4c24834,denim,denim
078d36ba-7645-4a50-8cbb-47555c4c0300,toys,Toys
8b6265e5-a2f8-4099-a7d3-2365d89ee879,halloween,Halloween
a30b090a-3f7e-48c6-ab45-378cfa6e27bc,elephants,Elephants
34793d93-33ca-400b-b057-ff45ac786694,ballerinas,Ballerinas
fb3bf2f7-f2b0-4281-9659-752ee8e35495,ships,Ships
8f593035-b0cc-482f-b600-b51027bf558e,seashells,Seashells
84238c3f-1388-4c4c-b3cb-f5dc47b9ad04,nursery-rhymes,Nursery Rhymes
2de9e121-5651-478c-8a39-4a4edc049ebe,books,Books
e3c2268f-6977-4aaf-9106-1685bcf4cc16,hairclip,hairclip
fe69cd09-d2c4-4182-a241-cac26f2335e8,gifts,Gifts
ce947ad8-aaf1-42e0-8b24-7bc82274798f,chess,Chess
5cfcf9e5-a425-4541-9587-ab532506f644,macarons,Macarons
85bfca68-43ba-4b66-89e7-8ef77e524dd7,pies,Pies
70072323-c877-4251-8ac3-f2d3609d133f,mermaids,Mermaids
90298b53-5960-4b91-9571-7ff0545a5ba4,canotier,canotier
d155164c-d71c-4370-abdb-b0ed0e988427,jams,Jams
a136fdf2-702e-427e-b5cb-5096a5a26cad,novala-takemoto,Novala Takemoto
6a9d0dcb-d98d-4787-9dfa-f1ddf95e86c5,chocolates,Chocolates
7172edc2-bdfe-4008-af88-4e495bff82bb,eiffel-tower,Eiffel Tower
e4437025-c44f-40d1-abca-31e7d8aa6261,choker,choker
ceab0b21-0aef-4908-bfe6-8dba6a63999a,cutlery,Cutlery
1b14b999-2f4a-4d3d-9bfa-a536ca339bfa,beads,beads
3674697a-ae1e-4188-81b4-8dc699894a35,acrylic-resin,acrylic resin
71ee2101-150b-4f58-a2db-365029f22331,spoons,Spoons
f32b774a-d0c1-475e-8c9c-17a2e5f3fdae,unicorns,Unicorns
b561e47d-a6d7-463c-b723-735b59444001,cosmetics,Cosmetics
df668af3-5c82-4a25-8076-bd053d8798cd,umbrellas,Umbrellas
14f4cf79-1d20-49fb-8d78-937f458dac31,lucky-pack,lucky pack
858c1323-7299-4414-8efa-b2bd782d3f4f,pina-sweetcollection,Pina sweetcollection
04dd7ad4-d36d-4926-8915-1c52704d1c2d,japanese-indie,Japanese Indie
b04d9fe3-f892-447a-aa71-39115de5e540,chocochip-cookie,Chocochip Cookie
ac29b233-3faa-4925-97d3-14d3ee739ffb,bubbles,Bubbles
b25a08b2-8abb-4d3c-837c-6f651735b0aa,clothing-and-shoe-prints,Clothing and Shoe Prints
d337f75a-4ca2-4312-b1fd-484179feaf8c,doughnuts,Doughnuts
984ef551-6a4d-4602-8c44-5e4ffd17b329,royalty,Royalty
8c213370-4984-43ed-9f75-d389f39f57e5,cinderella,Cinderella
a0085398-5d31-4d00-ab04-7f5b27e75808,lipsticks,Lipsticks
17d38bc2-6f4f-47f3-98e6-4cf3e767579f,astronomy-space,Astronomy Space
2c0e20ff-09cb-4542-a5ab-5a1798e7c32b,snowflakes,Snowflakes
258fe6f0-9224-4cb9-91f4-8187fc69f0f3,fans,Fans
1e615461-f8e5-4627-ac8b-c27afe84ea54,plastic,plastic
d5869c48-5d65-43ea-8c73-179745cb6172,necklace,necklace
3c6e9b58-d567-4d96-a289-44fce7e0adf6,rings,Rings
ac069808-12bd-4339-bd18-0a944d73da9d,bracelet,bracelet
5e4e4472-c68d-4681-b023-cb14c217c9f2,toile,toile
fcbd1189-918a-41d2-9e96-761b4bd4de0f,buildings,Buildings
48c1912d-7660-4c32-8033-85e690d86417,fairies,Fairies
72b53869-f3da-4f88-8126-056df30e1057,tassels,tassels
589e877b-bcbc-410a-a303-54af1ff0f629,masks,Masks
6d991b39-5b78-42f9-9490-3188b2e3d26e,shoes-print,Shoes Print
9a4d8906-98fa-4a8e-91be-f011d7134b71,scrunchie,scrunchie
cfe1fb80-23b2-4238-8125-3b9bbb6de577,shoe-clips,shoe clips
f61c7697-6ab4-437d-b3e8-f6c0b6bf43f5,tarot,Tarot
84f19366-1076-4cbb-ae26-b30dd6f893e6,gloomy-bear,Gloomy Bear
9d8d799a-6e93-424b-a0ba-4cd0a8cfb776,coffins,Coffins
abf6935b-8032-4e52-9e50-e1189c629512,earrings,Earrings
f4fd2742-16a4-4fef-827d-ec57fd313c3c,gloves,gloves
759fb641-c6ab-4dee-80c1-0557b9676d4b,arm-warmers,arm warmers
0184ac6f-3d84-4b7f-bd84-af44b07bf376,kuragehime,Kuragehime
3ab7fa9e-0219-4418-84dd-28898c3cda44,religious-motifs,Religious Motifs
edd99c4e-caa2-46ae-85b5-86f1879be3a5,plates,Plates
b0896a36-cf23-497a-881a-9911b37e37be,lustyn-wonderland,Lusty'n Wonderland
5b55d2ef-d295-41b2-bcd8-cba97003a7c1,western-indie,Western Indie
a9c2cbbf-f0a9-4f18-a5e8-021c6af29866,gilet,gilet
67f0bb4f-d4ea-454c-bc65-0e77e7d68308,raspberries,Raspberries
4e536381-e94a-4250-bf61-478c8ef0e0c8,dolphins,Dolphins
84faadf0-73a0-4fda-9312-8d8d9a44ff13,disney,Disney
959a70e3-06fd-4acc-9a65-1ff931c03a32,paintings,Paintings
5d13dd7f-5301-4338-855a-8e090d79a943,arts,Arts
0f393374-49bc-47f1-83aa-496b86f6d242,print-replica,Print Replica
903c3cf3-536a-4442-b3b9-7f54c3fbde33,forks,Forks
e10da7a8-277a-4d9e-80a7-7cab5020c027,purses-print,Purses Print
f965ffb7-74d5-4fd8-ac5b-7180404248c1,classical-lolita-indie-brand-name,Classical Lolita (Indie Brand Name)
a1e8f6d6-8256-4272-a43f-6007f7fbd604,leather,leather
545feed3-b37d-4865-8879-c3dc0c630b39,foxes,Foxes
657bdb91-7157-4242-81eb-928f7a2d333b,baroque,Baroque
05610265-fcde-4efd-89ed-7a413d634420,mew,Mew
0eec7034-789e-4bba-9199-41457c13dd65,dear-margaret,Dear Margaret
88d14fb4-709f-4a2f-be46-c42042f99e38,puppies,Puppies
eaa90388-6f97-4883-a9b1-8134d63f7b26,cherubs,Cherubs
137c117d-f0b5-4057-8507-322a4105f86e,lapin-agill,Lapin Agill
a6a4f9ca-0567-4110-a9e0-ed8089e061a9,laurels,Laurels
0db4ef1f-effb-4d29-b917-574a470e51b3,korean-indie,Korean Indie
d841132c-c3ea-4a8e-95a3-f718a21752c4,doors,Doors
1e21c4ee-a745-4609-ae4d-f716871c3abb,spica,Spica
fbb1a252-e442-4e8e-b188-9fa767423d72,ptmy,PtMY
7e4be614-7dff-4582-be52-90baacdc54a6,recalled,recalled
46934c25-bdae-4472-b1ab-c6d170d9487a,gramm,GRAMM
f2f15f12-e333-41a4-954c-af6c2e75ae14,hands,Hands
189fcdc1-067e-4c8f-9a2c-e83e35236849,shotgun-wedding,Shotgun Wedding
54de7f53-958c-488c-9f03-2b0f3049069b,eyes,Eyes
c804f702-8b58-42d0-a7a9-89e8f8780e2b,nurse,nurse
0576b53a-c59f-4682-9b2f-1415b043265b,brooches,brooches
7d32c077-4354-45c4-bce1-b98a7dbf892f,r-series,R Series
8d0cb50e-2a4b-4636-afac-a9bc1002fc90,creamy-mami,Creamy Mami
62e4ee22-e10a-470a-bb79-2294c85dcebe,marchenmerry,Marchenmerry
8a439fa9-423c-4216-ba72-13b401db3e44,kidsyoyo,KidsYoyo
cfa01ebf-44a9-44fc-a914-4f555e17ecd5,mam,MAM
afca0316-fdad-48de-a207-a77dd8464ad0,mice,Mice
05ce0cf3-8358-4465-b408-c9e0daab384f,fairy-wish,Fairy Wish
9cbb71a1-a868-4fbe-a7e1-24feaa2b9512,boguta,Boguta
f6aa50fc-c64e-46ce-8abf-b9ab47251a89,magic-potion,Magic Potion
aafdb820-cbef-4f2e-a23c-b263c2b07c8e,chinese-indie,Chinese Indie
510c43cb-2618-4d37-abed-1fe660631f17,snow-globes,Snow Globes
f62d5937-8b6d-4f5e-9fb8-e0f90703d598,strawberry-on-the-shortcake,Strawberry on the Shortcake
7b3dd3ff-198c-4e06-a9fc-483eb8a7f650,lethes-castle,Lethe's Castle
9aad6f2c-a07e-492f-99a0-84151f2f2f53,little-chili-shop,Little Chili Shop
2187df4b-282c-4186-aa2a-3f8011f27ffc,sunglasses,Sunglasses
6a1a4ace-d5e0-45c6-8e7a-cbe766abe71c,peter-pan,Peter Pan
99aeec19-3bf3-4e27-8257-3dfbbe603dda,violets,Violets
dc0a8b4b-083a-468a-b01d-7196e43a1139,phantom-of-the-opera,Phantom of the Opera
95b48c4b-d6d2-4208-9247-8e61adf39a85,pumpkin-cat,Pumpkin Cat
a54ffece-5faa-46c9-a1af-9a62db5fb735,the-snow-field,The Snow Field
2b129174-161e-4919-a986-8b21485da293,dark-box,Dark Box
79645b26-f427-4b47-9bc6-920ce48dcd72,lady-sloth,Lady Sloth
5ffe3a83-30e8-44aa-b161-539625659856,mystery-garden,Mystery Garden
2fe741d5-7743-407b-bb44-64240ef5fd96,puppets,Puppets
9286103e-97b9-4e2a-b454-12f0553855be,dear-celine,Dear Celine
de974322-dcd1-4ca8-93bc-0684c729cbb0,patterns-not-a-tag,Patterns (Not A Tag)
7c0a6ccf-56ae-4fcc-b299-c9073d4620b6,rose-trianon,Rose Trianon
ba308f5d-3492-40a9-a622-40a141402c6b,weddings,Weddings
be8b93cd-f53a-44ae-8e58-76c9198ef34a,magic-tea-party,Magic Tea Party
18e282f2-dc5a-4222-a03d-4320d7863b93,cloud-chamber,Cloud Chamber
f019c054-1904-49d7-ae4e-332bd2946ccc,fanplusfriend,Fanplusfriend
86cfb6fc-767e-41ac-9fdb-ce81fc98f104,hmhm,HMHM
6e6620da-bc81-416c-bae1-955a3974979d,white-moon,White Moon
8661330c-e99e-4ba1-b80e-c3189fc2e6d4,leur-getter,Leur Getter
71f32f96-2573-4ca5-bed5-9f89f67d6810,were-all-mad-here,We're All Mad Here
dceb49ec-f180-412c-9669-57b686fe7759,cherie-cerise,Cherie Cerise
07988043-c876-4295-9110-264244c0bc48,elpress-l,Elpress L
96ebbd67-ddbf-4d64-98f9-a6a24891fca8,policy,policy
1635480b-17d9-49a6-9cfd-2c1a39226829,pegasus,Pegasus
fb63e6a0-956f-470c-b473-778fbdbbd6fe,triple-fortune,Triple Fortune
713a3325-daa6-4193-adb2-0e2d54271ee8,little-red-riding-hood,Little Red Riding Hood
2ff33588-c576-4919-b765-8b8022e92f89,swimmer,SWIMMER
81d371a5-ed25-47db-9d4b-ecb285ae5872,harps,Harps
4931a57b-e5df-4463-8d9c-f393f696e173,gears,Gears
1dcd303c-ab56-41c0-a668-5daa73a86460,veils,veils
f824cf09-ccc4-4d13-8f56-ec51c32572f1,marchen-die-prinzessin,Marchen die Prinzessin
da1fc3d9-fe79-4e37-b74e-280d20b427bf,the-nutcracker,The Nutcracker
42a02dcc-7aba-4db3-96f8-f63dcb3fb25f,the-little-mermaid,The Little Mermaid
5352a8a1-ae65-425e-a1df-d72b8e53a043,beauty-and-the-beast,Beauty and the Beast
69768313-fab6-40ae-a862-66bd8a1d0b5c,rapunzel,Rapunzel
5550d4db-d881-4c16-936d-1a98128c20d9,soufflesong,Soufflesong
c4c439f6-bdbf-43dd-9efe-ebf535332fa7,lollipops,Lollipops
1618d9ed-cff1-411c-b328-a748b2faaf60,orchestra,Orchestra
c7299dbe-21cd-4d22-81eb-c86c7d33d43e,chocomint,Chocomint
*/

var tags []*domain.Tag

func main() {
	//
}