package dtos

type LoginRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
