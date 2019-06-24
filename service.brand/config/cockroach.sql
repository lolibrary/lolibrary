create table if not exists brands
(
    id          uuid                   not null,
    image_id    uuid                   not null,
    slug        character varying(255) not null,
    short_name  character varying(255) not null,
    name        character varying(255) not null,
    description text,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    primary key (id)
);

create unique index brands_slug_unique on brands (slug);
create unique index brands_short_name_unique on brands (short_name);
