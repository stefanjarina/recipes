package dtos

type CreateRecipeDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateRecipeDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
