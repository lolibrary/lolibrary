create table if not exists features
(
    id   uuid                   not null,
    slug character varying(255) not null,
    name character varying(255) not null,
    primary key (id)
);

create unique index if not exists features_slug_unique on features (slug);
