package chat

import (
	"fmt"
	"tiktok-go/repository"
)

func SaveMsg(r *repository.ChatRecord) error {
	return repository.NewChatRecordInstance().Create(r)
}

func GetMsgList(chatKey string) (*[]repository.ChatRecord, error) {
	return repository.NewChatRecordInstance().ListByKey(chatKey)
}

// simple follow for the exist genChatKey function
func GenChatKey(userA int64, userB int64) string {
	if userA > userB {
		return fmt.Sprintf("%d_%d", userB, userA)
	}
	return fmt.Sprintf("%d_%d", userA, userB)
}
