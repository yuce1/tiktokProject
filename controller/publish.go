package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync/atomic"
	"strconv"
	"tiktok-go/repository"
	service_user "tiktok-go/service/user"
	service_video "tiktok-go/service/video"
	utils "tiktok-go/utils"

	"github.com/gin-gonic/gin"
)

var videoIdSequence = int64(1)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	// TODO: file path tidy

	token := c.PostForm("token")

	if _, exist := service_user.GetUserByToken(token); !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	atomic.AddInt64(&userIdSequence, 1)
	videoIdstr := strconv.FormatInt(userIdSequence, 10)
	filename := videoIdstr+ ".mp4"
	user, _ := service_user.GetUserByToken(token)
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// gen cover
	var (
		coverFile string
		coverUrl  string
	)

	coverFile = videoIdstr+ ".jpg"
	finalCover := fmt.Sprintf("%d_%s", user.Id, coverFile)
	coverErr := utils.GenVideoCover(saveFile, filepath.Join("./public/", finalCover))

	host := os.Getenv("Host")
	videoUrl := host + fmt.Sprintf("/static/%s", finalName)
	if coverErr != nil {
		coverUrl = DemoVideos[0].CoverUrl
	} else {
		coverUrl = host + fmt.Sprintf("/static/%s", finalCover)
	}

	// add database
	video := &repository.Video{
		Author:   user.Name,
		PlayUrl:  videoUrl,
		CoverUrl: coverUrl,
	}

	// need remove the file which insert faild?
	// if err = service_video.PublishVideo(video); err != nil {
	// upload file can overwrite, not need to processing require
	// }
	service_video.PublishVideo(video)

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {

	// What the describe mean? is not resonable
	// I will follow this: list all videos that this user published

	var (
		id     int64
		err    error
		videos *[]repository.Video
	)

	token := c.Query("token")
	if id, err = strconv.ParseInt(c.Query("user_id"), 10, 64); err != nil {
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
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 2,
				StatusMsg:  "User authenticate fail.",
			},
			VideoList: nil,
		})
		return
	}

	if videos, err = service_video.GetVideosByUsername(u.Name); err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 3,
				StatusMsg:  "Get Video list faild.",
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
