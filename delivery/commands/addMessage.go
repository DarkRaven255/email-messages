package commands

type AddMessageCmd struct {
	Email       string `json:"email" validate:"required,email"`
	Title       string `json:"title" validate:"required"`
	Content     string `json:"content" validate:"required"`
	MagicNumber int    `json:"magic_number" validate:"required,numeric"`
}
