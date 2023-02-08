package repository

import (
	"sync"
	"time"

	"gorm.io/gorm"
)

type Favorite struct {
	Id        int64 `gorm:"primary_key" json:"id,omitempty"`
	UserId    int64 `json:"user_id,omitempty"`
	VideoId   int64 `json:"video_id,omitempty"`
	CreatedAt int64 `json:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty"`
	DeleteAt  int64 `json:"delete_at,omitempty"`
}

type FavoriteDao struct {
	db *gorm.DB
}

var (
	favoriteDao *FavoriteDao
	once        sync.Once
)

func NewFavouriteInstance() *FavoriteDao {
	once.Do(func() {
		favoriteDao = &FavoriteDao{
			db: db,
		}
	})
	return favoriteDao
}

type FavouriteOption func(f *Favorite)

func SetUserID(id int64) FavouriteOption {
	return func(f *Favorite) {
		f.UserId = id
	}
}

func SetVideoID(id int64) FavouriteOption {
	return func(f *Favorite) {
		f.VideoId = id
	}
}

func NewFavourite(option ...FavouriteOption) *Favorite {
	f := &Favorite{
		UserId:  0,
		VideoId: 0,
	}
	for _, opt := range option {
		opt(f)
	}
	return f
}

// the function below, paramter need to be Favourite? or just give some key value ,such as UserID, VideoID?
// Use a struct as paramter will add code complx or prefermence influence?

func (dao *FavoriteDao) Create(f *Favorite) error {
	return dao.db.Create(f).Error
}

func (dao *FavoriteDao) Delete(f *Favorite) error {

	t := time.Now().Unix()

	err := dao.db.Where("user_id = ? and video_id = ?", f.UserId, f.VideoId).Delete(&Favorite{}).Error

	if err == nil {
		// consist for record and var
		f.DeleteAt = t
	}

	return err
}

func (dao *FavoriteDao) SoftDelete(f *Favorite) error {

	t := time.Now().Unix()

	err := dao.db.Model(&Favorite{}).Where(
		"user_id = ? and video_id = ?",
		f.UserId, f.VideoId,
	).Update("delete_at", t).Error

	if err == nil {
		f.DeleteAt = t
	}

	return err
}

func (dao *FavoriteDao) CheckFavouriteExist(f *Favorite) bool {
	return dao.db.Where("user_id = ? and video_id = ?", f.UserId, f.VideoId).First(&Favorite{}).Error == nil
}

func (dao *FavoriteDao) ListFavouriteByUserId(userId int64) (*[]Favorite, error) {
	var f *[]Favorite
	result := dao.db.Where("user_id = ?", userId).Find(f)
	return f, result.Error
}

func (dao *FavoriteDao) ListFavouriteByVideoId(videoId int64) (*[]Favorite, error) {
	var f *[]Favorite
	result := dao.db.Where("video_id = ?", videoId).Find(f)
	return f, result.Error
}
