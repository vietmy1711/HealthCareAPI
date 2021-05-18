-- +migrate Up
CREATE TABLE "account" (
                           "userid" varchar PRIMARY KEY,
                           "username" varchar,
                           "blood" integer,
                           "gender" integer
);

CREATE TABLE "healthday" (
                             "userid" varchar,
                             "createat" date PRIMARY KEY,
                             "water" integer,
                             "steps" integer,
                             "heartrate" integer,
                             "calogries" integer,
                             "height" float,
                             "weight" float
);

ALTER TABLE "healthday" ADD CONSTRAINT "acountheath" FOREIGN KEY ("userid") REFERENCES "account" ("userid");
-- +migrate Down
DROP TABLE "healthday";
DROP TABLE "account";
