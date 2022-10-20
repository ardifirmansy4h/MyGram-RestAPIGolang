package repository

import (
	"final-projek/models/mcomments"
	"final-projek/models/mphotos"
	"final-projek/models/msocialmedias"
	"final-projek/models/musers"
)

type UserRepository interface {
	GetAllUser() []musers.User
	GetByIDUser(id string) musers.User
	CreateUser(input musers.InputUser) musers.ResponseCreateUser
	UpdateUser(id string, input musers.InputUser) musers.ResponseUpdateUser
	DeleteUser(id string) bool
}

type AuthRepository interface {
	Register(input musers.InputRegister) musers.ResponseCreateUser
	Login(input musers.InputLogin) string
}

type SocialMediaRepository interface {
	GetAllSocialMedia() []msocialmedias.ResponseGetSM
	GetByIDSocialMedia(id string) msocialmedias.SocialMedia
	CreateSocialMedia(input msocialmedias.InputSocialMedia) msocialmedias.ResponseCreateSocialMedia
	UpdateSocialMedia(id string, input msocialmedias.InputSocialMedia) msocialmedias.ResponseUpdateSocialMedia
	DeleteSocialMedia(id string) bool
}

type PhotoRepository interface {
	GetAllPhoto() []mphotos.ResponseGetPhoto
	GetByIDPhoto(id string) mphotos.Photo
	CreatePhoto(input mphotos.InputPhoto) mphotos.ReponseInputPhoto
	UpdatePhoto(id string, input mphotos.InputPhoto) mphotos.ReponseUpdatePhoto
	DeletePhoto(id string) bool
}

type CommentRepository interface {
	GetAllComment() []mcomments.ResponseGetComments
	GetByIDComment(id string) mcomments.Comment
	CreateCommment(input mcomments.InputComment) mcomments.ResponseInputComment
	UpdateComment(id string, input mcomments.InputComment) mcomments.ResponseUpdateComment
	DeleteComment(id string) bool
}
