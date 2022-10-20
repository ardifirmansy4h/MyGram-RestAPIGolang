package susers

import (
	"final-projek/models/musers"
	"final-projek/repository"
	"final-projek/repository/rusers"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService() UserService {
	return UserService{
		userRepository: &rusers.UserRepositoryIMPL{},
	}
}

func (us *UserService) GetAllUser() []musers.User {
	return us.userRepository.GetAllUser()
}

func (us *UserService) GetByIDUser(id string) musers.User {
	return us.userRepository.GetByIDUser(id)
}

func (us *UserService) CreateUser(input musers.InputUser) musers.ResponseCreateUser {
	return us.userRepository.CreateUser(input)
}

func (us *UserService) UpdateUser(id string, input musers.InputUser) musers.ResponseUpdateUser {
	return us.userRepository.UpdateUser(id, input)
}

func (us *UserService) DeleteUser(id string) bool {
	return us.userRepository.DeleteUser(id)
}
