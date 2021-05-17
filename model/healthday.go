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
	Heartrate int `json:"heartrate" db:"heartrate"`
	Calogries int `json:"calogries" db:"calogries"`
	Height float64 `json:"height" db:"height"`
	Weight float64 `json:"weight" db:"weight"`
}

