package musers

import "github.com/go-playground/validator/v10"

type InputLogin struct {
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type InputRegister struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Age      int    `json:"age" validate:"required,min=8"`
}

func (input *InputRegister) ValidRegister() error {
	validate := validator.New()
	err := validate.Struct(input)
	return err

}
func (input *InputLogin) ValidLogin() error {
	validate := validator.New()
	err := validate.Struct(input)
	return err

}
