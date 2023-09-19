package req

type ReqSignUp struct {
	Email    string `json:"email,omitempty" validate:"email"`
	Password string `json:"password,omitempty" validate:"pwd"`
	FullName string `json:"fullName,omitempty" validate:"required"`
}