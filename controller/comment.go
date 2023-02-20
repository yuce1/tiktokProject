package controller

import (
	"net/http"

	"tiktok-go/repository"
	service_comment "tiktok-go/service/comment"

	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	userid := c.GetInt64("UserID")
	actionType := c.Query("action_type")
	videoidstr := c.Query("video_id")
	videoid, _ := strconv.ParseInt(videoidstr, 10, 64)
	user, _ := repository.NewUserDaoInstance().GetUserById(userid)

	// Add Comment
	if actionType == "1" {
		timeStr := time.Now().Format("2006-01-02 15:04:05")
		text := c.Query("comment_text")

		comment := &repository.Comment{
			UserId:     userid,
			VideoId:    videoid,
			Content:    text,
			CreateDate: timeStr,
		}

		service_comment.PublishComment(comment)

		c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0, StatusMsg: "publish successfully"},
			Comment: Comment{
				Id:         comment.Id,
				User:       *RepoUserToCon(user),
				Content:    text,
				CreateDate: timeStr,
			},
		})
		return
	} else { // Delete Comment
		commentidstr := c.Query("comment_id")
		commentid, _ := strconv.ParseInt(commentidstr, 10, 64)
		service_comment.DeleteComment(commentid)
		c.JSON(http.StatusOK, Response{StatusCode: 0})
		return
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoidstr := c.Query("video_id")
	videoid, _ := strconv.ParseInt(videoidstr, 10, 64)
	comments, err := service_comment.GetCommentByVideoId(videoid)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "get comment failed",
			},
			CommentList: nil,
		})
		return
	}

	var respCommentList []Comment
	for _, comment := range *comments {
		respCommentList = append(respCommentList, *RepoCommentToCon(&comment))
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: respCommentList,
	})
}
