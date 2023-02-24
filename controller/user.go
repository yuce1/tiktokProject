package controller

import (
	"log"
	"net/http"
	"strconv"

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
	User User `json:"user"`
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

	var (
		id  int64
		err error
	)

	// if c.GetBool("Visitor") { // the visitor can see all the info in person page
	// } else {
	// 	id = c.GetInt64("UserID")
	// }

	if id, err = strconv.ParseInt(c.Query("user_id"), 10, 64); err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{
				StatusCode: http.StatusOK,
				StatusMsg:  "Invaild user id",
			},
		})
	}

	user, _ := service_user.GetUserbyId(id)

	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     *RepoUserToCon(user),
	})
}
