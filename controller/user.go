package controller

import (
	"net/http"
	"sync/atomic"

	service_user "tiktok-go/service/user"

	"github.com/gin-gonic/gin"
)

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if exist := service_user.CheckUserExist(username); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		service_user.RegisterUser(username, password)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    service_user.GenerateToken(username, password),
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if user, exist := service_user.GetUserWithVerify(username, password); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    service_user.GenerateToken(username, password),
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	if user, exist := service_user.GetUserByToken(token); exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     *RepoUserToCon(user),
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist or token invalid"},
		})
	}
}
