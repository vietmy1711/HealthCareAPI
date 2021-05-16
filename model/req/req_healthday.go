package req

type ReqHealthDay struct {
	Userid string `json:"userid,omitempty" validate:"required`
	Createat string `json:"createat,omitempty" validate:"required`
	Water int `json:"water,omitempty" validate:"required`
	Steps int `json:"steps,omitempty" validate:"required`
	Heartrate int `json:"heartrate,omitempty" validate:"required`
	Calogries int `json:"calogries,omitempty" validate:"required`
	Height float64 `json:"height,omitempty" validate:"required`
	Weight float64 `json:"weight,omitempty" validate:"required`
}