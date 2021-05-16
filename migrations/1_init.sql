-- +migrate Up
CREATE TABLE "account"
(
    "iduser"   varchar PRIMARY KEY,
    "username" varchar,
    "blood"    integer,
    "gender"   integer
);

CREATE TABLE "healthday"
(
    "iduser"    varchar,
    "createat"  date PRIMARY KEY,
    "water"     integer,
    "steps"     integer,
    "heartrate" integer,
    "calogries" integer,
    "height"    float,
    "weight"    float
);

ALTER TABLE "healthday"
    ADD CONSTRAINT "acountheath" FOREIGN KEY ("iduser") REFERENCES "account" ("iduser");

-- +migrate Down
DROP TABLE "healthday";
DROP TABLE "user";