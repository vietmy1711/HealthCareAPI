package req
type ReqUpdateInfo struct {
	Userid    string `json:"userid,omitempty" validate:"required"`
	FullName string `json:"username,omitempty" validate:"required"` // tags
	Age int `json:"age,omitempty" validate:"required"`
}
