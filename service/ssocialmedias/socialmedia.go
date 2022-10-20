package ssocialmedias

import (
	"final-projek/models/msocialmedias"
	"final-projek/repository"
	"final-projek/repository/rsocialmedias"
)

type SocialMediaService struct {
	socialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaService() SocialMediaService {
	return SocialMediaService{
		socialMediaRepository: &rsocialmedias.SocialMediaRepositoryIMPL{},
	}
}

func (ss *SocialMediaService) GetAllSocialMedia() []msocialmedias.ResponseGetSM {
	return ss.socialMediaRepository.GetAllSocialMedia()
}

func (ss *SocialMediaService) GetByIDSocialMedia(id string) msocialmedias.SocialMedia {
	return ss.socialMediaRepository.GetByIDSocialMedia(id)
}
func (ss *SocialMediaService) CreateSocialMedia(input msocialmedias.InputSocialMedia) msocialmedias.ResponseCreateSocialMedia {
	return ss.socialMediaRepository.CreateSocialMedia(input)
}

func (ss *SocialMediaService) UpdateSocialMedia(id string, input msocialmedias.InputSocialMedia) msocialmedias.ResponseUpdateSocialMedia {
	return ss.socialMediaRepository.UpdateSocialMedia(id, input)
}

func (ss *SocialMediaService) DeleteSocialMedia(id string) bool {
	return ss.socialMediaRepository.DeleteSocialMedia(id)
}
