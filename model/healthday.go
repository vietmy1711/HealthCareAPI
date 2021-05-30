package model

import "time"

//"iduser" varchar,
//"createat" date PRIMARY KEY,
//"water" integer,
//"steps" integer,
//"heartrate" integer,
//"calogries" integer,
//"height" float,
//"weight" float

type HealthDay struct {
	Userid string `json:"userid" db:"userid"`
	Createat time.Time `json:"createat" db:"createat"`
	Water int `json:"water" db:"water"`
	Steps int `json:"steps" db:"steps"`
	Heartrate float64 `json:"heartrate" db:"heartrate"`
	Calogries float64 `json:"calogries" db:"calories"`
	Height float64 `json:"height" db:"height"`
	Weight float64 `json:"weight" db:"weight"`
	ActiveEnergyBunred float64 `json:"active_energy_bunred" db:"active_energy_bunred"`
	BasalEnergyBunred float64 `json:"basal_energy_bunred" db:"basal_energy_bunred"`
	BloodOxygen float64 `json:"blood_oxygen" db:"blood_oxygen"`
}

