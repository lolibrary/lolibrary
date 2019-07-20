create table if not exists images
(
    id         uuid                   not null,
    name       character varying(255) not null,
    filename   character varying(255) not null,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    primary key (id)
);

create table if not exists image_item
(
    id         bigserial primary key,
    image_id   uuid not null,
    item_id    uuid not null,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);

-- create the unique index - you can only have an image on an item once.
create unique index if not exists image_item_compound_unique on image_item (item_id, image_id);

-- create a separate index used for searching.
create index if not exists image_item_item_id_index on image_item (item_id);
