create table if not exists colors
(
    id         uuid                   not null,
    slug       character varying(255) not null,
    name       character varying(255) not null,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    primary key (id)
);

create unique index if not exists colors_slug_unique on colors (slug);
