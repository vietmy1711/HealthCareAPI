-- +migrate Up
create table "account" (
    "userid" text primary key,
    "fullname" text,
    "email" text,
    "password" text,
    "created_at" date NOT NULL
);
-- +migrate Down
DROP TABLE account;