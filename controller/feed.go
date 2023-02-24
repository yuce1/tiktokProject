package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"tiktok-go/repository"
	service_favor "tiktok-go/service/favorite"
	service_video "tiktok-go/service/video"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// TODO: the video IsFollow field is invaild
// Feed same demo video list for every request
func Feed(c *gin.Context) {

	var (
		err       error
		timeStamp int64
	)

	// gin will give a default timeStr but test not, so there need a verify and attach a default timeStr
	timeStr := c.Query("latest_time")
	if timeStr != "" {
		timeStamp, err = strconv.ParseInt(timeStr, 10, 64)
	} else {
		timeStamp = time.Now().Unix()
	}

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

	// need add favorite info
	if !c.GetBool("Visitor") {
		addFavoriteInfo(c.GetInt64("UserID"), videos)
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

func addFavoriteInfo(userid int64, videos *[]repository.Video) {

	favors, _ := service_favor.ListByUserId(userid)

	for _, f := range *favors {
		for i, v := range *videos {
			if v.Id == f.VideoId {
				(*videos)[i].IsFavorite = true
				break
			}
		}
	}
}
