package msocialmedias

import (
	"final-projek/models/musers"
	"time"

	"github.com/go-playground/validator/v10"
)

type SocialMedia struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           musers.User
}

type ResponseGetSM struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	} `json:"user"`
}

type InputSocialMedia struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
	UserID         uint   `json:"user_id"`
}

func (input *InputSocialMedia) ValidInputSM() error {
	validate := validator.New()
	err := validate.Struct(input)
	return err

}

type ResponseCreateSocialMedia struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type ResponseUpdateSocialMedia struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}
