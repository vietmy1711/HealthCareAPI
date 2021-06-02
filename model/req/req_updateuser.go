package req
type ReqUpdateUser struct {
	Userid string `json:"userid,omitempty" validate:"required`
	Token string `json:"token,omitempty" validate:"required`
}
