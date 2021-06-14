package model

type WaningModel struct {
	AboutWater int `json:"aboutWater" db:"aboutKcal"`
	AboutStep int `json:"aboutStep" db:"aboutKcal"`
	AboutKcal float64 `json:"aboutKcal" db:"aboutKcal"`
	BRM float64 `json:"BRM" db:"BRM"`
	TDEE float64 `json:"TDEE" db:"TDEE"`
	StatusWater int `json:"StatusWater" db:"StatusWater"`
	StatusStep int `json:"StatusStep" db:"StatusStep"`
	StatusKcal int `json:"StatusKcal" db:"StatusKcal"`
}
