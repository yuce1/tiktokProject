package controller

import (
	"net/http"
	"strconv"

	"tiktok-go/repository"
	service_relation "tiktok-go/service/relation"
	service_user "tiktok-go/service/user"

	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	if user, exist := service_user.GetUserByToken(token); exist {
		if actionType == "1" { //Add friend
			touseridstr := c.Query("to_user_id")
			touserid, _ := strconv.ParseInt(touseridstr, 10, 64)
			relation := &repository.Relation{
				FromId: user.Id,
				ToId:   touserid,
			}

			service_relation.AddFriend(relation)

			c.JSON(http.StatusOK, Response{StatusCode: 0})
			return
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	useridstr := c.Query("user_id")
	userid, _ := strconv.ParseInt(useridstr, 10, 64)
	relations, err := service_relation.GetRelationListById(userid)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "None FollowList find",
			},
			UserList: []User{DemoUser},
		})
	}

	var respFollowList []User
	for _, relation := range *relations {
		user, _ := repository.NewUserDaoInstance().GetUserById(relation.ToId)
		respFollowList = append(respFollowList, *RepoUserToCon(user))
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: respFollowList,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}
