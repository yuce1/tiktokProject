package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"tiktok-go/repository"
	service_user "tiktok-go/service/user"
	service_video "tiktok-go/service/video"
	utils "tiktok-go/utils"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	// TODO: file path tidy

	var (
		u *repository.User
		// exist bool
	)

	if c.GetBool("TokenProvide") {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Token not provide"})
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

	//chech if the video was published
	byteContainer := make([]byte, 1000000)
	fileContent, _ := data.Open()
	fileContent.Read(byteContainer)
	hash_code := service_user.GetSHA256HashCode([]byte(byteContainer))
	_, err = service_video.CheckVideo(hash_code)
	if err != nil {
		if err.Error() != "record not found" {
			c.JSON(http.StatusOK, Response{
				StatusCode: 2,
				StatusMsg:  err.Error(),
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: 3,
			StatusMsg:  "video had been publish before",
		})
		return
	}

	id := c.GetInt64("UserID")

	filename := hash_code + ".mp4"
	user, _ := service_user.GetUserbyId(id)
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

	coverFile = hash_code + ".jpg"
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
		HashCode: hash_code,
		Title:    c.PostForm("title"),
	}

	// need remove the file which insert faild?
	// if err = service_video.PublishVideo(video); err != nil {
	// upload file can overwrite, not need to processing require
	// }
	service_video.PublishVideo(video, u.Id)

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// TODO: visitor request need be impl
// PublishList all users have same publish video list
func PublishList(c *gin.Context) {

	// What the describe mean? is not resonable
	// I will follow this: list all videos that this user published

	var (
		err    error
		videos *[]repository.Video
	)

	id := c.GetInt64("UserID")

	u, exist := service_user.GetUserbyId(id)
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
