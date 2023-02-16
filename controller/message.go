package controller

import (
	"net/http"
	"strconv"
	"tiktok-go/repository"
	service_chat "tiktok-go/service/chat"
	service_user "tiktok-go/service/user"

	"github.com/gin-gonic/gin"
)

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")

	user, exist := service_user.GetUserByToken(token)
	if !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	userIdB, _ := strconv.Atoi(toUserId)
	chatKey := service_chat.GenChatKey(user.Id, int64(userIdB))
	curMessage := repository.ChatRecord{
		ChatKey: chatKey,
		Content: content,
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
	token := c.Query("token")
	toUserId := c.Query("to_user_id")

	user, exist := service_user.GetUserByToken(token)
	if !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	userIdB, _ := strconv.Atoi(toUserId)
	chatKey := service_chat.GenChatKey(user.Id, int64(userIdB))

	chatRecord, err := service_chat.GetMsgList(chatKey)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 2, StatusMsg: "Fetch chat list faild."})
		return
	}
	resp := make([]Message, len(*chatRecord))
	for _, obj := range *chatRecord {
		resp = append(resp, *RepoChatToMsg(&obj))
	}

	c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: resp})
}
