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
	Calories float64 `json:"calories" db:"calories"`
	Height float64 `json:"height" db:"height"`
	Weight float64 `json:"weight" db:"weight"`
	ActiveEnergyBurned float64 `json:"active_energy_burned" db:"active_energy_burned"`
	BasalEnergyBurned float64 `json:"basal_energy_burned" db:"basal_energy_burned"`
	BloodOxygen float64 `json:"blood_oxygen" db:"blood_oxygen"`
	DistanceWalkingRunning float64 `json:"distance_walking_running" db:"distance_walking_running"`
}

