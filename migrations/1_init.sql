-- +migrate Up
CREATE TABLE "healthday" (
                             "userid" varchar,
                             "createat" TIMESTAMP,
                             "water" integer,
                             "steps" integer,
                             "heartrate" integer,
                             "calories" integer,
                             "height" float,
                             "weight" float,
                             "active_energy_bunred" float,
                             "basal_energy_bunred" float,
                             "blood_oxygen" float,
                             "distance_walking_running" float,
                             PRIMARY KEY ("userid", "createat")
);

ALTER TABLE "healthday" ADD CONSTRAINT "accountheath" FOREIGN KEY ("userid") REFERENCES "account" ("userid");
-- +migrate Down
DROP TABLE "healthday";
