package req
type ReqGetUser struct {
	Userid string `json:"userid,omitempty" validate:"required`
}
