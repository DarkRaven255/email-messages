package commands

type SendMessagesCmd struct {
	MagicNumber int `json:"magic_number" validate:"required,numeric"`
}
