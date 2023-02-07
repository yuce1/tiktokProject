package controller

import (
	"net/http"

	"tiktok-go/repository"
	service_comment "tiktok-go/service/comment"
	service_user "tiktok-go/service/user"

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
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := service_user.GetUserByToken(token); exist {
		// Add Comment
		if actionType == "1" {
			timeStr := time.Now().Format("2006-01-02 15:04:05")
			text := c.Query("comment_text")
			videoidstr := c.Query("video_id")
			videoid, _ := strconv.ParseInt(videoidstr, 10, 64)

			comment := &repository.Comment{
				UserId:     user.Id,
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

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
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
