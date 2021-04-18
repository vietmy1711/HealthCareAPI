package req

type ReqSignUp struct {
	Userid    string `json:"userid,omitempty" validate:"required"`
	FullName string `json:"fullName,omitempty" validate:"required"` // tags
	Gender int `json:"gender,omitempty" validate:"required"`
	Blood int `json:"blood,omitempty" validate:"required"`
}
