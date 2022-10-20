package scomments

import (
	"final-projek/models/mcomments"
	"final-projek/repository"
	"final-projek/repository/rcomments"
)

type CommentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentRepository() CommentService {
	return CommentService{
		commentRepository: &rcomments.CommentRepositoryIMPL{},
	}
}

func (cs *CommentService) GetAllComment() []mcomments.ResponseGetComments {
	return cs.commentRepository.GetAllComment()
}

func (cs *CommentService) GetByIDComment(id string) mcomments.Comment {
	return cs.commentRepository.GetByIDComment(id)
}

func (cs *CommentService) CreateCommment(input mcomments.InputComment) mcomments.ResponseInputComment {
	return cs.commentRepository.CreateCommment(input)
}

func (cs *CommentService) UpdateComment(id string, input mcomments.InputComment) mcomments.ResponseUpdateComment {
	return cs.commentRepository.UpdateComment(id, input)
}

func (cs *CommentService) DeleteComment(id string) bool {
	return cs.commentRepository.DeleteComment(id)
}
