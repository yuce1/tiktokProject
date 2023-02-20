package repository

import (
	"sync"

	"gorm.io/gorm"
)

type User struct {
	Id             int64  `gorm:"primary_key" json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Passwd         string `json:"passwd,omitempty"`
	CreatedAt      int64  `json:"created_at,omitempty"`
	UpdatedAt      int64  `json:"updated_at,omitempty"`
	DeleteAt       int64  `json:"delete_at,omitempty"`
	WorkCount      int64  `json:"work_count"`
	FavoriteCount  int64  `json:"favorite_count"`
	FollowCount    int64  `json:"follow_count"`
	FollowerCount  int64  `json:"follower_count"`
	Avatar         string `json:"avatar,omitempty"`
	BackgrounImage string `json:"background_image,omitempty"`
	Signature      string `json:"signature,omitempty"`
	TotalFavorited int64  `json:"total_favorited,omitempty"`
	// I want to remove this field in db, this field should consider with other single user
	// But there is some code is using this field
	IsFollow bool `json:"is_follow"`
}

func NewUser() *User {
	return &User{
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}
}

type UserDao struct {
	db *gorm.DB
}

var (
	userDao  *UserDao
	userOnce sync.Once
)

func NewUserDaoInstance() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{
			db: db,
		}
	})
	return userDao
}

func (t *UserDao) CreateUser(user *User) error {
	result := t.db.Create(user)
	return result.Error
}

func (t *UserDao) VerifyUser(username string, password string) (*User, bool) {
	var user User
	result := t.db.Where("name = ? AND passwd = ?", username, password).First(&user)
	return &user, result.RowsAffected == 1
}

func (t *UserDao) CheckUserExist(username string) bool {
	var user User
	result := t.db.Where("name = ?", username).First(&user)
	return result.RowsAffected == 1
}

func (t *UserDao) GetUserByName(username string) (*User, bool) {
	var user User
	result := t.db.Where("name = ?", username).First(&user)
	return &user, result.RowsAffected == 1
}

func (t *UserDao) GetUserById(userid int64) (*User, bool) {
	var user User
	result := t.db.Where("id = ?", userid).First(&user)
	return &user, result.RowsAffected == 1
}

func (t *UserDao) UpdateWorkCount(userid int64, action int) error {
	// action must be 1 or -1
	result := t.db.Model(&User{Id: userid}).UpdateColumn("work_count", gorm.Expr("work_count + ?", action))
	return result.Error
}

func (t *UserDao) UpdateFavoriteCount(userid int64, action int) error {
	// action must be 1 or -1
	result := t.db.Model(&User{Id: userid}).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", action))
	return result.Error
}

func (t *UserDao) UpdateFollowCount(userid int64, action int) error {
	// action must be 1 or -1
	result := t.db.Model(&User{Id: userid}).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", action))
	return result.Error
}

func (t *UserDao) UpdateFollowerCount(userid int64, action int) error {
	// action must be 1 or -1
	result := t.db.Model(&User{Id: userid}).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", action))
	return result.Error
}
