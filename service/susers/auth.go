package susers

import (
	"final-projek/models/musers"
	"final-projek/repository"
	"final-projek/repository/rusers"
)

type AuthService struct {
	AuthRepository repository.AuthRepository
}

func NewUserAuth() AuthService {
	return AuthService{
		AuthRepository: &rusers.UserAuthIMPL{},
	}
}
func (as *AuthService) Register(input musers.InputRegister) musers.ResponseCreateUser {
	return as.AuthRepository.Register(input)
}

func (as *AuthService) Login(input musers.InputLogin) string {
	return as.AuthRepository.Login(input)
}
