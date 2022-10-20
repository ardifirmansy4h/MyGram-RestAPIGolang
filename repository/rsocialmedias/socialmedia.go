package rsocialmedias

import (
	"final-projek/app/config"
	"final-projek/models/msocialmedias"
)

type SocialMediaRepositoryIMPL struct{}

func (sr *SocialMediaRepositoryIMPL) GetAllSocialMedia() []msocialmedias.ResponseGetSM {
	var socialmedia []msocialmedias.SocialMedia
	config.DB.Preload("User").Find(&socialmedia)
	var responseData []msocialmedias.ResponseGetSM
	for _, sms := range socialmedia {
		input := msocialmedias.ResponseGetSM{}
		input.ID = sms.ID
		input.Name = sms.Name
		input.UserID = sms.UserID
		input.SocialMediaUrl = sms.SocialMediaUrl
		input.CreatedAt = sms.CreatedAt
		input.UpdatedAt = sms.UpdatedAt
		input.User.ID = sms.User.ID
		input.User.Username = sms.User.Username
		responseData = append(responseData, input)
	}
	return responseData
}

func (sr *SocialMediaRepositoryIMPL) GetByIDSocialMedia(id string) msocialmedias.SocialMedia {
	var socialmedia msocialmedias.SocialMedia

	config.DB.First(&socialmedia, "id = ?", id)
	return socialmedia
}

func (sr *SocialMediaRepositoryIMPL) CreateSocialMedia(input msocialmedias.InputSocialMedia) msocialmedias.ResponseCreateSocialMedia {

	var cSocial msocialmedias.SocialMedia = msocialmedias.SocialMedia{
		Name:           input.Name,
		SocialMediaUrl: input.SocialMediaUrl,
		UserID:         input.UserID,
	}

	res := config.DB.Create(&cSocial)
	social := msocialmedias.ResponseCreateSocialMedia{}

	res.Last(&social)
	return social

}

func (sr *SocialMediaRepositoryIMPL) UpdateSocialMedia(id string, input msocialmedias.InputSocialMedia) msocialmedias.ResponseUpdateSocialMedia {
	var uSocial msocialmedias.SocialMedia = sr.GetByIDSocialMedia(id)

	uSocial.Name = input.Name
	uSocial.SocialMediaUrl = input.SocialMediaUrl
	res := config.DB.Save(&uSocial)

	social := msocialmedias.ResponseUpdateSocialMedia{}
	res.Last(&social)
	return social

}

func (sr *SocialMediaRepositoryIMPL) DeleteSocialMedia(id string) bool {
	var dSocial msocialmedias.SocialMedia = sr.GetByIDSocialMedia(id)
	res := config.DB.Delete(&dSocial)

	if res.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
