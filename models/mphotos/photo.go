package mphotos

import (
	"final-projek/models/musers"
	"time"

	"github.com/go-playground/validator/v10"
)

type Photo struct {
	ID        uint         `json:"id"`
	Title     string       `json:"title"`
	Caption   string       `json:"caption"`
	PhotoUrl  string       `json:"photo_url "`
	UserID    uint         `json:"user_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      *musers.User `json:"user" gorm:"foreignKey:UserID"`
}

type ResponseGetPhoto struct {
	ID        uint      `json:"id,omitempty"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption,omitempty" `
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	User      struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"user"`
}

type InputPhoto struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" validate:"required"`
	UserID   uint   `json:"user_id"`
}

func (input *InputPhoto) ValidInputPhoto() error {
	validate := validator.New()
	err := validate.Struct(input)
	return err

}

type ReponseInputPhoto struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url "`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ReponseUpdatePhoto struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url "`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
