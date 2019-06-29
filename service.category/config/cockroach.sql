create table if not exists categories
(
    id   uuid                   not null,
    slug character varying(255) not null,
    name character varying(255) not null,
    primary key (id)
);

create unique index if not exists categories_slug_unique on categories (slug);
