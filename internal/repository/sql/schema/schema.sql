create table users
(
    id            bigserial
        primary key,
    created_at    timestamp with time zone,
    updated_at    timestamp with time zone,
    deleted_at    timestamp with time zone,
    user_name     varchar(255) not null
        unique,
    full_name     varchar(255) not null
        unique,
    email         varchar(100) not null
        unique,
    password_hash varchar(100) not null,
    profile_image text    default ''::text,
    is_verified   boolean default false
);

alter table users
    owner to clj;

create index idx_users_deleted_at
    on users (deleted_at);



create table news
(
    id          bigserial
        primary key,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,
    author_id   bigint
        constraint fk_news_author
            references users
            on update cascade on delete set null,
    cover_image text,
    title       text,
    sub_title   text,
    content     text
);

alter table news
    owner to clj;

create index idx_news_author_id
    on news (author_id);

create index idx_news_deleted_at
    on news (deleted_at);

