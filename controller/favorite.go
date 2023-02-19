package controller

import (
	"log"
	"net/http"
	"strconv"
	"tiktok-go/repository"
	service_favor "tiktok-go/service/favorite"
	service_video "tiktok-go/service/video"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	var err error

	id := c.GetInt64("UserID")

	// are there need a video legal verify?

	video, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	favour := repository.Favorite{UserId: id, VideoId: video}

	actioType := c.Query("action_type")

	switch actioType {
	case "1":
		err = service_favor.Do(&favour, id)
	case "2":
		err = service_favor.Undo(&favour, id)
	}

	if err != nil {
		log.Printf("[WARN] Favourite action on User[id: %d] faild, ERR: %s", id, err)
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

	id = c.GetInt64("UserID")

	favors, err := service_favor.ListByUserId(id)
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
