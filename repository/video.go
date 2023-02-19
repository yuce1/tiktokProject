package repository

import (
	"sync"

	"gorm.io/gorm"
)

// Author可以作为外键
type Video struct {
	Id            int64  `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	Author        string `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	CreatedAt     int64  `json:"created_at,omitempty"`
	UpdatedAt     int64  `json:"updated_at,omitempty"`
	DeleteAt      int64  `json:"delete_at,omitempty"`
	HashCode      string `json:"hash_code,omitempty"`
	Title         string `json:"title,omitempty"`
}

type VideoDao struct {
	db *gorm.DB
}

var (
	videoDao  *VideoDao
	videoOnce sync.Once
)

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(func() {
		videoDao = &VideoDao{
			db: db,
		}
	})
	return videoDao
}

func (v *VideoDao) CreateVideo(video *Video) error {
	res := v.db.Create(video)
	return res.Error
}

func (v *VideoDao) GetVideoByID(id int64) (*Video, error) {
	var video Video
	res := v.db.Where("id = ?", id).First(&video)
	return &video, res.Error
}

func (v *VideoDao) GetVideosByAuthor(username string) (*[]Video, error) {
	var videos []Video
	res := v.db.Where("author = ?", username).Order("created_at DESC").Find(&videos)
	return &videos, res.Error
}

func (v *VideoDao) GetVideoList(timeStamp int64) (*[]Video, error) {
	var video []Video
	res := v.db.Where("created_at <= ?", timeStamp).Order("created_at DESC").Limit(30).Find(&video)
	return &video, res.Error
}

func (v *VideoDao) CheckVideoHash(hash_code string) (*Video, error) {
	var video Video
	res := v.db.Where("hash_code = ?", hash_code).First(&video)
	return &video, res.Error
}

func (v *VideoDao) UpdateCommentCount(videoid int64, op int) error {
	var video Video
	res := v.db.Where("id = ?", videoid).First(&video)
	if op == 1 {
		video.CommentCount += 1
	} else {
		video.CommentCount -= 1
	}

	v.db.Save(&video)
	return res.Error
}

func (v *VideoDao) ListByVideoIdSet(idSet []int64) (*[]Video, error) {
	var videos []Video
	result := v.db.Where("id in ?", idSet).Find(&videos)
	return &videos, result.Error
}
