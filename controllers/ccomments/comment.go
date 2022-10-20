package ccomments

import (
	"final-projek/controllers"
	"final-projek/models/mcomments"
	"final-projek/service/scomments"
	"net/http"

	"github.com/labstack/echo/v4"
)

var serviceComment scomments.CommentService = scomments.NewCommentRepository()

func GetAllComment(c echo.Context) error {
	comment := serviceComment.GetAllComment()
	res := controllers.NewResponse(c, http.StatusOK, "succes", "all comment", comment)
	return res
}
func GetByIDComment(c echo.Context) error {
	commentID := c.Param("id")

	comment := serviceComment.GetByIDComment(commentID)

	if comment.ID == 0 {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "Id Tidak Ada", "")
	} else {
		return controllers.NewResponse(c, http.StatusAccepted, "success", "Id ditemukan", "")
	}

}
func CreateCommment(c echo.Context) error {
	input := new(mcomments.InputComment)
	if err := c.Bind(input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "Created Comments", "")
	}
	err := input.ValidInputComment()
	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "field tidak boleh kosong", "")
	}
	comment := serviceComment.CreateCommment(*input)

	return controllers.NewResponse(c, http.StatusCreated, "success", "Created comments", comment)
}

func UpdateComment(c echo.Context) error {
	commentID := c.Param("id")
	input := new(mcomments.InputComment)
	if err := c.Bind(input); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "Update Comment", "")
	}
	comment := serviceComment.UpdateComment(commentID, *input)

	return controllers.NewResponse(c, http.StatusCreated, "success", "your comment has been succesfully updated", comment)
}

func DeleteComment(c echo.Context) error {
	CommentID := c.Param("id")
	dComment := serviceComment.DeleteComment(CommentID)
	if !dComment {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "your comment has been failed deleted", "")
	} else {
		return controllers.NewResponse(c, http.StatusOK, "success", "your comment has been succesfully deleted", "")
	}
}
