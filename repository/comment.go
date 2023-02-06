package repository

import (
	"sync"

	"gorm.io/gorm"
)

type Comment struct {
	Id            int64  `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	UserId        int64 `json:"uid"`
	VideoId       int64 `json:"vid"`
	Content string `json:"content,omitempty"`
	CreatedAt int64 `json:"created_at,omitempty"`
	CreateDate string `json:"created_date,omitempty"`
	UpCount int64  `json:"upcount"`
	DownCount  int64  `json:"downcount"`
}

type CommentDao struct {
	db *gorm.DB
}

var (
	commentDao  *CommentDao
	commentOnce sync.Once
)

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(func() {
		commentDao = &CommentDao{
			db: db,
		}
	})
	return commentDao
}

func (t *CommentDao) CreateComment(comment *Comment) error {
	result := t.db.Create(comment)
	return result.Error
}

func (t *CommentDao) GetCommentList(vid int64) (*[]Comment, error) {
	var comments []Comment
	res := t.db.Where("video_id = ?", vid).Order("created_at DESC").Find(&comments)
	return &comments, res.Error
}