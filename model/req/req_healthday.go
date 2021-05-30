package req

type ReqHealthDay struct {
	Userid string `json:"userid,omitempty" validate:"required`
	Water int `json:"water,omitempty" validate:"required`
	Steps int `json:"steps,omitempty" validate:"required`
	Heartrate float64 `json:"heartrate,omitempty" validate:"required`
	Calogries float64 `json:"calogries,omitempty" validate:"required`
	Height float64 `json:"height,omitempty" validate:"required`
	Weight float64 `json:"weight,omitempty" validate:"required`
	ActiveEnergyBunred float64 `json:"active_energy_bunred,omitempty" validate:"required`
	BasalEnergyBunred float64 `json:"basal_energy_bunred,omitempty" validate:"required`
	BloodOxygen float64 `json:"blood_oxygen,omitempty" validate:"required`
}