package controller

import (
	"log"
	"net/http"

	service_user "tiktok-go/service/user"

	"github.com/gin-gonic/gin"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User UserInfoResp `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	password = service_user.GetSHA256HashCode([]byte(password))

	if exist := service_user.CheckUserExist(username); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		id, token, _ := service_user.RegisterUser(username, password)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   id,
			Token:    token,
		})
		log.Printf("[INFO] User: {ID: %d, NAME: %s, PASS: %s} is CREATE", id, username, password)
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	password = service_user.GetSHA256HashCode([]byte(password))

	if user, exist := service_user.GetUserWithVerify(username, password); exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    service_user.GenerateToken(user.Id),
		})

		log.Printf("[INFO] User: {ID: %d, NAME: %s, PASS: %s} is LOGIN", user.Id, user.Name, user.Passwd)

	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist or password errors"},
		})
	}
}

func UserInfo(c *gin.Context) {
	if c.GetBool("TokenProvide") {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist or token invalid"},
		})
	}

	id := c.GetInt64("UserID")
	user, _ := service_user.GetUserbyId(id)

	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     *RepoUserToInfo(user),
	})
}
