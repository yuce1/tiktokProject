package repository

import (
	"sync"

	"gorm.io/gorm"
)

type ChatRecord struct {
	Id         int64  `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	ChatKey    string `json:"chat_key"`
	FromUserId int64  `json:"from_user_id"`
	ToUserId   int64  `json:"to_user_id"`
	Content    string `json:"content"`
	CreatedAt  int64  `json:"created_at,omitempty"`
}

type ChatRecordDao struct {
	db *gorm.DB
}

var (
	chatRecordDao  *ChatRecordDao
	chatRecordOnce sync.Once
)

func NewChatRecordInstance() *ChatRecordDao {
	chatRecordOnce.Do(func() {
		chatRecordDao = &ChatRecordDao{
			db: db,
		}
	})
	return chatRecordDao
}

func (c *ChatRecordDao) Create(r *ChatRecord) error {
	return c.db.Save(r).Error
}

func (c *ChatRecordDao) ListByKey(chatKey string) (*[]ChatRecord, error) {
	var cr []ChatRecord
	result := c.db.Where("chat_key = ?", chatKey).Order("created_at ASC").Find(&cr)
	return &cr, result.Error
}
