package model

type User struct {
	UserId string `json:"userid,omitempty" db:"userid, omitempty"`
	FullName string `json:"fullName,omitempty" db:"username, omitempty"`
	Blood int `json:"blood,omitempty" db:"blood, omitempty"` // A = 1, B = 2, AB = 3, O = 4
	Gender int `json:"gender,omitempty" db:"gender, omitempty"` // men = 1, women = 2, other = 3
}
