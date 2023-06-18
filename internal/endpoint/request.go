package endpoint

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	PasswordAgain string `json:"passwordAgain"`
}
