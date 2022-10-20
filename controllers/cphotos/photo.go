package cphotos

import (
	"final-projek/controllers"
	"final-projek/models/mphotos"
	"final-projek/service/sphotos"
	"net/http"

	"github.com/labstack/echo/v4"
)

var servicePhoto sphotos.PhotoService = sphotos.NewPhotoRepository()

func GetAllPhoto(c echo.Context) error {
	photo := servicePhoto.GetAllPhoto()

	res := controllers.NewResponse(c, http.StatusOK, "succes", "all photo", photo)
	return res
}

func GetByIDPhoto(c echo.Context) error {
	photoID := c.Param("id")

	photo := servicePhoto.GetByIDPhoto(photoID)

	if photo.ID == 0 {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "Id Tidak Ada", "")
	} else {
		return controllers.NewResponse(c, http.StatusAccepted, "success", "Id ditemukan", "")
	}
}

func CreatePhoto(c echo.Context) error {
	input := new(mphotos.InputPhoto)
	if err := c.Bind(input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "Created Photo", "")
	}
	err := input.ValidInputPhoto()
	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "field tidak boleh kosong", "")
	}

	photo := servicePhoto.CreatePhoto(*input)

	return controllers.NewResponse(c, http.StatusCreated, "success", "Created Photo", photo)
}

func UpdatePhoto(c echo.Context) error {
	photoID := c.Param("id")
	input := new(mphotos.InputPhoto)
	if err := c.Bind(input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "Update Photo", "")
	}
	photo := servicePhoto.UpdatePhoto(photoID, *input)

	return controllers.NewResponse(c, http.StatusCreated, "success", "your photo has been succesfully updated", photo)
}

func DeletePhoto(c echo.Context) error {
	photoID := c.Param("id")

	dPhoto := servicePhoto.DeletePhoto(photoID)
	if !dPhoto {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "your photo has been failed deleted", "")
	} else {
		return controllers.NewResponse(c, http.StatusOK, "success", "your photo has been succesfully deleted", "")
	}
}
