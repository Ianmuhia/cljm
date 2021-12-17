CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "email" varchar(100) NOT NULL,
  "password" varchar NOT NULL,
  "dp" varchar NOT NULL,
  "access_level" int ,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
);
