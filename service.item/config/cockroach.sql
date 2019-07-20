create table if not exists items
(
    id             uuid                   not null PRIMARY KEY,
    brand_id       uuid                   not null,
    category_id    uuid                   not null,
    user_id        uuid                   not null,
    image_id       uuid                   not null,
    publisher_id   uuid,
    slug           character varying(255) not null,
    english_name   character varying(300) not null,
    foreign_name   character varying(300) not null,
    product_number character varying(255) not null,
    notes          character varying      not null,
    price          character varying(255),
    currency       character varying(255),
    year           integer                not null,
    status         integer                not null default 0,
    metadata       jsonb                           default '{}',
    created_at     timestamp with time zone,
    updated_at     timestamp with time zone,
    published_at   timestamp with time zone
);

-- slug has to be unique
create unique index if not exists items_slug_unique on items (slug);

-- we search on these fields a lot.
create index if not exists items_brand_id_index on items (brand_id);
create index if not exists items_category_id_index on items (category_id);
create index if not exists items_user_id_index on items (user_id);

-- index created/updated as we sort on them a *lot*.
create index if not exists items_created_at_index on items (created_at);
create index if not exists items_updated_at_index on items (updated_at);

-- create an inverted index to search json metadata.
create index if not exists items_metadata_inverted on items using gin (metadata);
