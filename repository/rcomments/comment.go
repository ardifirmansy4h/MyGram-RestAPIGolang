package rcomments

import (
	"final-projek/app/config"
	"final-projek/models/mcomments"

	"github.com/jinzhu/copier"
)

type CommentRepositoryIMPL struct{}

func (cr *CommentRepositoryIMPL) GetAllComment() []mcomments.ResponseGetComments {
	var comment []mcomments.Comment

	config.DB.Preload("User").Preload("Photo").Find(&comment)

	var response []mcomments.ResponseGetComments

	for _, comments := range comment {
		var ResponseData mcomments.ResponseGetComments
		copier.Copy(&ResponseData, &comments)
		response = append(response, ResponseData)
	}

	return response

}

func (cr *CommentRepositoryIMPL) GetByIDComment(id string) mcomments.Comment {
	var comment mcomments.Comment
	config.DB.First(&comment, "id = ?", id)

	return comment
}

func (cr *CommentRepositoryIMPL) CreateCommment(input mcomments.InputComment) mcomments.ResponseInputComment {
	var comment mcomments.Comment = mcomments.Comment{
		Message: input.Message,
		UserID:  input.UserID,
		PhotoID: input.PhotoID,
	}
	res := config.DB.Create(&comment)

	comments := mcomments.ResponseInputComment{}
	res.Last(&comments)
	return comments
}

func (cr *CommentRepositoryIMPL) UpdateComment(id string, input mcomments.InputComment) mcomments.ResponseUpdateComment {
	var comment mcomments.Comment = cr.GetByIDComment(id)

	comment.Message = input.Message

	res := config.DB.Save(&comment)
	comments := mcomments.ResponseUpdateComment{}
	res.Last(&comments)
	return comments

}

func (cr *CommentRepositoryIMPL) DeleteComment(id string) bool {
	var comment mcomments.Comment = cr.GetByIDComment(id)
	res := config.DB.Delete(comment)

	if res.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
