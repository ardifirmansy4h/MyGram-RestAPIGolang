package rusers

import (
	"final-projek/app/config"
	"final-projek/helper/token"
	"final-projek/models/musers"

	"golang.org/x/crypto/bcrypt"
)

type UserAuthIMPL struct{}

func (ur *UserAuthIMPL) Register(input musers.InputRegister) musers.ResponseCreateUser {
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	var newUser musers.User = musers.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(password),
		Age:      input.Age,
	}

	res := config.DB.Create(&newUser)
	user := musers.ResponseCreateUser{}
	res.Last(&user)
	return user

}

func (ur *UserAuthIMPL) Login(input musers.InputLogin) string {
	var user musers.User = musers.User{}

	config.DB.First(&user, "email", input.Email)
	if user.ID == 0 {
		return ""
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return ""
	}
	tokens := token.CreateToken(user.ID)

	return tokens
}
