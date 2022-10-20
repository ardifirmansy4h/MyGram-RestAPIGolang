package rphotos

import (
	"final-projek/app/config"
	"final-projek/models/mphotos"
)

type PhotoRepositoryIMPL struct{}

func (pr *PhotoRepositoryIMPL) GetAllPhoto() []mphotos.ResponseGetPhoto {
	var photo []mphotos.Photo

	config.DB.Preload("User").Find(&photo)

	var responseData []mphotos.ResponseGetPhoto
	for _, photos := range photo {
		input := mphotos.ResponseGetPhoto{}
		input.ID = photos.ID
		input.Title = photos.Title
		input.UserID = photos.UserID
		input.PhotoUrl = photos.PhotoUrl
		input.Caption = photos.Caption
		input.CreatedAt = photos.CreatedAt
		input.UpdatedAt = photos.UpdatedAt
		input.User.Username = photos.User.Username
		input.User.Email = photos.User.Email
		responseData = append(responseData, input)
	}
	return responseData
}

func (pr *PhotoRepositoryIMPL) GetByIDPhoto(id string) mphotos.Photo {
	var photo mphotos.Photo
	config.DB.First(&photo, "id = ?", id)
	return photo
}

func (pr *PhotoRepositoryIMPL) CreatePhoto(input mphotos.InputPhoto) mphotos.ReponseInputPhoto {

	var photo mphotos.Photo = mphotos.Photo{
		Title:    input.Title,
		Caption:  input.Caption,
		PhotoUrl: input.PhotoUrl,
		UserID:   input.UserID,
	}
	res := config.DB.Create(&photo)
	photos := mphotos.ReponseInputPhoto{}
	res.Last(&photos)
	return photos
}

func (pr *PhotoRepositoryIMPL) UpdatePhoto(id string, input mphotos.InputPhoto) mphotos.ReponseUpdatePhoto {
	var photo mphotos.Photo = pr.GetByIDPhoto(id)

	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoUrl = input.PhotoUrl
	photo.UserID = input.UserID

	res := config.DB.Save(&photo)
	photos := mphotos.ReponseUpdatePhoto{}
	res.Last(&photos)
	return photos
}

func (pr *PhotoRepositoryIMPL) DeletePhoto(id string) bool {
	var photo mphotos.Photo = pr.GetByIDPhoto(id)

	res := config.DB.Delete(photo)

	if res.RowsAffected == 0 {
		return false
	} else {
		return true
	}

}
