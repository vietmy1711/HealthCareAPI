package req
type ReqWater struct {
	Userid string `json:"userid,omitempty" validate:"required`
	Water int `json:"water,omitempty" validate:"required`
}
