package repository

import (
	"sync"

	"gorm.io/gorm"
)

type Favorite struct {
	UserId    int64 `gorm:"primary_key;autoIncrement:false" json:"user_id,omitempty"`
	VideoId   int64 `gorm:"primary_key;autoIncrement:false" json:"video_id,omitempty"`
	CreatedAt int64 `json:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty"`
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
	return dao.db.Save(f).Error
}

func (dao *FavoriteDao) Delete(f *Favorite) error {
	return dao.db.Delete(f).Error
}

// func (dao *FavoriteDao) GetDesignedFavourite(f *Favorite) error {
// 	result := dao.db.Where("user_id = ? and video_id = ?", f.UserId, f.VideoId).First(&Favorite{})
// 	return result.Error
// }

func (dao *FavoriteDao) ListFavouriteByUserId(userId int64) (*[]Favorite, error) {
	var f []Favorite
	result := dao.db.Where("user_id = ?", userId).Find(&f)
	return &f, result.Error
}

func (dao *FavoriteDao) ListFavouriteByVideoId(videoId int64) (*[]Favorite, error) {
	var f []Favorite
	result := dao.db.Where("video_id = ?", videoId).Find(&f)
	return &f, result.Error
}
