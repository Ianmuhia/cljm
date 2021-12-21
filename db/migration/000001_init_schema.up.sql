CREATE TABLE "users" (
  "id" bigserial,
  "name" varchar PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "dp" varchar,
  "access_level" int ,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "news" (
  "id" bigserial PRIMARY KEY,
  "user" varchar,
  "cover_image" varchar NOT NULL,
  "images" varchar,
  "title" varchar NOT NULL,
  "subtitle" varchar NOT NULL,
  "content" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "news_images" (
  "image" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "blog" (
  "id" bigserial PRIMARY KEY,
  "user" varchar,
  "cover_image" varchar NOT NULL,
  "title" varchar NOT NULL,
  "subtitle" varchar NOT NULL,
  "content" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "blog_images" (
  "image" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "books" (
  "id" bigserial PRIMARY KEY,
  "synopsis" varchar NOT NULL,
  "author" varchar NOT NULL,
  "genre" varchar NOT NULL,
  "file" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "genre" (
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "gallery" (
  "id" bigserial PRIMARY KEY,
  "image" varchar,
  "caption" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "gallery_images" (
  "image" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sermon" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "subtitle" varchar NOT NULL,
  "content" varchar NOT NULL,
  "video_data" varchar NOT NULL
);

ALTER TABLE "news" ADD FOREIGN KEY ("user") REFERENCES "users" ("name") ON DELETE CASCADE;
ALTER TABLE "blog" ADD FOREIGN KEY ("user") REFERENCES "users" ("name") ON DELETE CASCADE;
ALTER TABLE "news" ADD FOREIGN KEY ("images") REFERENCES "news_images" ("image");
ALTER TABLE "blog" ADD FOREIGN KEY ("images") REFERENCES "blog_images" ("image");
ALTER TABLE "gallery" ADD FOREIGN KEY ("images") REFERENCES "gallery_images" ("image");
ALTER TABLE "books" ADD FOREIGN KEY ("genre") REFERENCES "genre" ("name");