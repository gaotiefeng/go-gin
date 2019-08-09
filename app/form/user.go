package form

type UserAddForm struct {
	Mobile	string	`validate:"required"`
	Password	string `validate:"required"`
}
