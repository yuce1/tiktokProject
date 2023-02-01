package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	service_video "tiktok-go/service/video"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {

	// gin will give a default timestamp
	timeStamp, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "Invaild timeStamp.",
			},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
		return
	}

	videos, err := service_video.GetStreamFeed(timeStamp)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: 2,
				StatusMsg:  fmt.Sprintf("Fetch video list faild. Error: %v", err),
			},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
		return
	}

	// maybe preallocate enough memory will better
	var respVideoList []Video
	for _, video := range *videos {
		respVideoList = append(respVideoList, *RepoVideoToCon(&video))
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: respVideoList,
		NextTime:  time.Now().Unix(),
	})
}
