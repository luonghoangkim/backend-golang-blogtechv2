package req

type ReqSignIn struct {
	Email    string `json:"email,omitempty" validate:"email"`
	Password string `json:"password,omitempty" validate:"pwd"`
}