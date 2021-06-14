package req

import "time"

type ReqHealthDay struct {
	Userid string `json:"userid,omitempty" validate:"required`
	Water int `json:"water,omitempty" validate:"required`
	Steps int `json:"steps,omitempty" validate:"required`
	Heartrate float64 `json:"heartrate,omitempty" validate:"required`
	Calories float64 `json:"calories,omitempty" validate:"required`
	Height float64 `json:"height,omitempty" validate:"required`
	Weight float64 `json:"weight,omitempty" validate:"required`
	ActiveEnergyBurned float64 `json:"active_energy_burned,omitempty" validate:"required`
	BasalEnergyBurned float64 `json:"basal_energy_burned,omitempty" validate:"required`
	BloodOxygen float64 `json:"blood_oxygen,omitempty" validate:"required`
	DistanceWalkingRunning float64 `json:"distance_walking_running" validate:"distance_walking_running"`
}

type ReqFakeHealthDay struct {
	Userid string `json:"userid,omitempty" validate:"required`
	Createat time.Time `json:"createat,omitempty" validate:"required`
	Water int `json:"water,omitempty" validate:"required`
	Steps int `json:"steps,omitempty" validate:"required`
	Heartrate float64 `json:"heartrate,omitempty" validate:"required`
	Calories float64 `json:"calories,omitempty" validate:"required`
	Height float64 `json:"height,omitempty" validate:"required`
	Weight float64 `json:"weight,omitempty" validate:"required`
	ActiveEnergyBurned float64 `json:"active_energy_burned,omitempty" validate:"required`
	BasalEnergyBurned float64 `json:"basal_energy_burned,omitempty" validate:"required`
	BloodOxygen float64 `json:"blood_oxygen,omitempty" validate:"required`
	DistanceWalkingRunning float64 `json:"distance_walking_running" validate:"distance_walking_running"`
}