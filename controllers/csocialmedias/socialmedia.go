package csocialmedias

import (
	"final-projek/controllers"
	"final-projek/models/msocialmedias"
	"final-projek/service/ssocialmedias"
	"net/http"

	"github.com/labstack/echo/v4"
)

var socialMedia ssocialmedias.SocialMediaService = ssocialmedias.NewSocialMediaService()

func GetAllSocialMedia(c echo.Context) error {
	socialmedias := socialMedia.GetAllSocialMedia()

	res := controllers.NewResponse(c, http.StatusOK, "success", "succes get all", socialmedias)

	return res

}

func GetByIDSocialMedia(c echo.Context) error {
	socialID := c.Param("id")

	social := socialMedia.GetByIDSocialMedia(socialID)

	if social.ID == 0 {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "ID Tidak ada", "")
	} else {
		return controllers.NewResponse(c, http.StatusOK, "success", "success get id", "")
	}

}

func CreateSocialMedia(c echo.Context) error {
	input := new(msocialmedias.InputSocialMedia)
	if err := c.Bind(input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "Created Social Media", "")
	}
	err := input.ValidInputSM()
	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "field tidak boleh kosong", "")
	}
	socialmedia := socialMedia.CreateSocialMedia(*input)

	return controllers.NewResponse(c, http.StatusCreated, "success", "Created Social Media", socialmedia)

}

func UpdateSocialMedia(c echo.Context) error {
	socialID := c.Param("id")
	input := new(msocialmedias.InputSocialMedia)
	if err := c.Bind(input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "Update Social Media", "")
	}
	social := socialMedia.UpdateSocialMedia(socialID, *input)

	return controllers.NewResponse(c, http.StatusCreated, "success", "your social media has been succesfully updated", social)
}

func DeleteSocialMedia(c echo.Context) error {
	socialID := c.Param("id")

	dSocial := socialMedia.DeleteSocialMedia(socialID)
	if !dSocial {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "your social media has been failed deleted", "")
	} else {
		return controllers.NewResponse(c, http.StatusOK, "success", "your social media has been succesfully deleted", "")
	}
}
