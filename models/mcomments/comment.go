package mcomments

import (
	"final-projek/models/mphotos"
	"final-projek/models/musers"
	"time"

	"github.com/go-playground/validator/v10"
)

type Comment struct {
	ID        uint           `json:"id"`
	Message   string         `json:"message"`
	PhotoID   uint           `json:"photo_id"`
	UserID    uint           `json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	User      *musers.User   `gorm:"foreignKey:UserID"`
	Photo     *mphotos.Photo `gorm:"foreignKey:PhotoID"`
}

type InputComment struct {
	Message string `json:"message" validate:"required"`
	PhotoID uint   `json:"photo_id"`
	UserID  uint   `json:"user_id"`
}

func (input *InputComment) ValidInputComment() error {
	validate := validator.New()
	err := validate.Struct(input)
	return err

}

type ResponseInputComment struct {
	ID        uint      `json:"id" `
	Message   string    `json:"message" `
	PhotoID   uint      `json:"photo_id" `
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type ResponseUpdateComment struct {
	ID        uint      `json:"id" `
	Message   string    `json:"message" `
	PhotoID   uint      `json:"photo_id" `
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ResponseGetComments struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	User      struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"user"`
	Photo struct {
		ID       uint   `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption,omitempty"`
		PhotoURL string `json:"photo_url"`
		UserID   uint   `json:"user_id"`
	} `json:"photo"`
}
