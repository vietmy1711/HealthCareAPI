-- +migrate Up
create table "user"(
    "user_id" text primary key,
    "full_name" text,
    "gender" int,
    "blood" int
);
-- +migrate Down
DROP TABLE "user";