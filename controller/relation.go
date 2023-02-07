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
			if repository.NewRelationDaoInstance().CheckRelation(user.Id, touserid) {
				c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "您已经关注过该用户"})
				return
			} else {
				relation := &repository.Relation{
					FromId: user.Id,
					ToId:   touserid,
				}
				service_relation.Follow(relation)

				c.JSON(http.StatusOK, Response{StatusCode: 0})
				return
			}
		} else {
			touseridstr := c.Query("to_user_id")
			touserid, _ := strconv.ParseInt(touseridstr, 10, 64)
			if !repository.NewRelationDaoInstance().CheckRelation(user.Id, touserid) {
				c.JSON(http.StatusOK, Response{StatusCode: 3, StatusMsg: "您已经取关该用户"})
				return
			} else {
				service_relation.UnFollow(user.Id, touserid)
				c.JSON(http.StatusOK, Response{StatusCode: 0})
				return
			}
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 2, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	useridstr := c.Query("user_id")
	userid, _ := strconv.ParseInt(useridstr, 10, 64)
	relations, err := service_relation.GetFollowListById(userid)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "FollowList failed",
			},
			UserList: []User{DemoUser},
		})
	}

	var respFollowList []User
	for _, relation := range *relations {
		user, _ := repository.NewUserDaoInstance().GetUserById(relation.ToId)
		user.IsFollow = true
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
	useridstr := c.Query("user_id")
	userid, _ := strconv.ParseInt(useridstr, 10, 64)
	relations, err := service_relation.GetFollowerListById(userid)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "FollowerList failed",
			},
			UserList: []User{DemoUser},
		})
	}

	var respFollowerList []User
	for _, relation := range *relations {
		user, _ := repository.NewUserDaoInstance().GetUserById(relation.FromId)
		user.IsFollow = repository.NewRelationDaoInstance().CheckRelation(userid, relation.FromId)
		respFollowerList = append(respFollowerList, *RepoUserToCon(user))
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: respFollowerList,
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
