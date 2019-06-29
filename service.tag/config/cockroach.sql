create table if not exists tags
(
    id   uuid                   not null,
    slug character varying(255) not null,
    name character varying(255) not null,
    primary key (id)
);

create unique index if not exists tags_slug_unique on tags (slug);
