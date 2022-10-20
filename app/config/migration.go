package config

import (
	"final-projek/models/mcomments"
	"final-projek/models/mphotos"
	"final-projek/models/msocialmedias"
	"final-projek/models/musers"
)

func DBMigrate() {
	DB.AutoMigrate(&musers.User{}, &msocialmedias.SocialMedia{}, &mphotos.Photo{}, &mcomments.Comment{})
}
