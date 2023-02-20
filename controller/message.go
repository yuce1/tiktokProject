package controller

import (
	"log"
	"net/http"
	"strconv"
	"tiktok-go/repository"
	service_chat "tiktok-go/service/chat"

	"github.com/gin-gonic/gin"
)

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {

	toUserId := c.Query("to_user_id")
	content := c.Query("content")

	id := c.GetInt64("UserID")

	userIdB, _ := strconv.ParseInt(toUserId, 10, 64)
	chatKey := service_chat.GenChatKey(id, userIdB)
	curMessage := repository.ChatRecord{
		ChatKey:    chatKey,
		FromUserId: id,
		ToUserId:   userIdB,
		Content:    content,
	}
	err := service_chat.SaveMsg(&curMessage)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 2, StatusMsg: "Save chat record faild."})
		return
	}

	c.JSON(http.StatusOK, Response{StatusCode: 0})
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {

	toUserId := c.Query("to_user_id")

	id := c.GetInt64("UserID")

	userIdB, _ := strconv.ParseInt(toUserId, 10, 64)
	chatKey := service_chat.GenChatKey(id, userIdB)

	chatRecord, err := service_chat.GetMsgList(chatKey)
	if err != nil {
		log.Printf("[WARN] Fetch chat list faild. %s", err)
		c.JSON(http.StatusOK, Response{StatusCode: 2, StatusMsg: "Fetch chat list faild."})
		return
	}
	resp := make([]Message, len(*chatRecord))
	for i, obj := range *chatRecord {
		resp[i] = *RepoChatToMsg(&obj)
	}

	c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: resp})
}
