package sphotos

import (
	"final-projek/models/mphotos"
	"final-projek/repository"
	"final-projek/repository/rphotos"
)

type PhotoService struct {
	PhotoRepository repository.PhotoRepository
}

func NewPhotoRepository() PhotoService {
	return PhotoService{
		PhotoRepository: &rphotos.PhotoRepositoryIMPL{},
	}
}

func (ps *PhotoService) GetAllPhoto() []mphotos.ResponseGetPhoto {
	return ps.PhotoRepository.GetAllPhoto()
}

func (ps *PhotoService) GetByIDPhoto(id string) mphotos.Photo {
	return ps.PhotoRepository.GetByIDPhoto(id)
}

func (ps *PhotoService) CreatePhoto(input mphotos.InputPhoto) mphotos.ReponseInputPhoto {
	return ps.PhotoRepository.CreatePhoto(input)
}

func (ps *PhotoService) UpdatePhoto(id string, input mphotos.InputPhoto) mphotos.ReponseUpdatePhoto {
	return ps.PhotoRepository.UpdatePhoto(id, input)
}

func (ps *PhotoService) DeletePhoto(id string) bool {
	return ps.PhotoRepository.DeletePhoto(id)
}
