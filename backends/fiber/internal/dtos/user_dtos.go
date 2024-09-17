package dtos

type UserToCreateDto struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

type UserToUpdateDto struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
}
