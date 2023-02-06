package user

import (
	"tiktok-go/repository"
)

func PublishComment(comment *repository.Comment) error {
	return repository.NewCommentDaoInstance().CreateComment(comment)
}

func GetCommentByVideoId(videoid int64) (*[]repository.Comment, error) {
	return repository.NewCommentDaoInstance().GetCommentList(videoid)
}