package controller

import (
	"log"
	"net/http"
	"strconv"
	"tiktok-go/repository"
	service_favor "tiktok-go/service/favorite"
	service_user "tiktok-go/service/user"
	service_video "tiktok-go/service/video"

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
		log.Printf("[WARN] Request User (token: %s) doesn't exist", token)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	// are there need a video legal verify?

	video, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	favour := repository.Favorite{UserId: u.Id, VideoId: video}

	actioType := c.Query("action_type")

	switch actioType {
	case "1":
		err = service_favor.Do(&favour, u.Id)
	case "2":
		err = service_favor.Undo(&favour, u.Id)
	}

	if err != nil {
		log.Printf("[WARN] Favourite action on User[id: %d] faild, ERR: %s", u.Id, err)
		c.JSON(http.StatusOK, Response{StatusCode: 2, StatusMsg: "Action faild."})
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Success"})
}

// TODO: visitor request need be impl
// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {

	var (
		id  int64
		err error
	)

	token := c.Query("token")
	if id, err = strconv.ParseInt(c.Query("user_id"), 10, 64); err != nil {
		log.Printf("[WARN] Request UserID is invaild, User[id: %d]", id)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Invaild user id",
			},
			VideoList: nil,
		})
		return
	}

	u, exist := service_user.GetUserByToken(token)
	if !exist || u.Id != id {
		// log.Printf("[WARN] User[id: %d] doesn't exist or authentication faild. [!exist = %t, (u.ID == id) = %t]", id, !exist, u.Id != id)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 2,
				StatusMsg:  "User authenticate fail.",
			},
			VideoList: nil,
		})
		return
	}

	favors, err := service_favor.ListByUserId(u.Id)
	if err != nil {
		log.Printf("[WARN] User[id: %d] Fetch favourite list faild, ERR: %s.", id, err)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 3,
				StatusMsg:  "Fetch favorite list faild.",
			},
			VideoList: nil,
		})
		return
	}

	videoIds := make([]int64, len(*favors))
	for _, f := range *favors {
		videoIds = append(videoIds, f.VideoId)
	}

	videos, err := service_video.GetVideoBySet(videoIds)
	if err != nil {
		log.Printf("[WARN] User[id: %d] Fetch favourite list faild, ERR: %s.", id, err)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 3,
				StatusMsg:  "Fetch favorite list faild.",
			},
			VideoList: nil,
		})
		return
	}

	var respVideoList []Video
	for _, video := range *videos {
		respVideoList = append(respVideoList, *RepoVideoToCon(&video))
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: respVideoList,
	})
}
