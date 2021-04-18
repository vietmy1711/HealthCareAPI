package req

type ReqSignIn struct {
	userid  string `json:"email,omitempty" validate:"required"`
}