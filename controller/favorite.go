package controller

import (
	"net/http"
	"strconv"
	"tiktok-go/repository"
	service_favor "tiktok-go/service/favorite"
	service_user "tiktok-go/service/user"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	var (
		u     *repository.User
		exist bool
		err   error
	)

	token := c.Query("token")

	if u, exist = service_user.GetUserByToken(token); !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

	// are there need a video legal verify?

	video, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	favour := repository.Favorite{UserId: u.Id, VideoId: video}

	actioType := c.Query("action_type")

	switch actioType {
	case "1":
		err = service_favor.Do(&favour)
	case "2":
		err = service_favor.Undo(&favour)
	}

	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 2, StatusMsg: "Action faild."})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
