package rusers

import (
	"final-projek/app/config"
	"final-projek/models/musers"

	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryIMPL struct{}

func (ur *UserRepositoryIMPL) GetAllUser() []musers.User {
	var users []musers.User

	config.DB.Find(&users)
	return users
}

func (ur *UserRepositoryIMPL) GetByIDUser(id string) musers.User {
	var user musers.User
	config.DB.First(&user, "id = ?", id)
	return user
}

func (ur *UserRepositoryIMPL) CreateUser(input musers.InputUser) musers.ResponseCreateUser {

	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	var cUser musers.User = musers.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(password),
		Age:      input.Age,
	}

	res := config.DB.Create(&cUser)
	user := musers.ResponseCreateUser{}
	res.Last(&user)
	return user
}

func (ur *UserRepositoryIMPL) UpdateUser(id string, input musers.InputUser) musers.ResponseUpdateUser {
	var uUser musers.User = ur.GetByIDUser(id)

	uUser.Username = input.Username
	uUser.Email = input.Email

	res := config.DB.Save(&uUser)

	user := musers.ResponseUpdateUser{}
	res.Last(&user)
	return user

}

func (ur *UserRepositoryIMPL) DeleteUser(id string) bool {
	var dUser musers.User = ur.GetByIDUser(id)

	res := config.DB.Delete(&dUser)

	if res.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
