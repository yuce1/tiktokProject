package repository

import (
	"sync"

	"gorm.io/gorm"
)

type User struct {
	Id        int64  `gorm:"primary_key" json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Passwd    string `json:"passwd,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
	DeleteAt  int64  `json:"delete_at,omitempty"`

	FollowCount   int64 `json:"follow_count"`
	FollowerCount int64 `json:"follower_count"`
	IsFollow      bool  `json:"is_follow"`
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
