package controller

import (
	"net/http"
	"strconv"

	"tiktok-go/repository"
	service_relation "tiktok-go/service/relation"

	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	userid := c.GetInt64("UserID")
	touseridstr := c.Query("to_user_id")
	touserid, _ := strconv.ParseInt(touseridstr, 10, 64)
	actionType := c.Query("action_type")
	if actionType == "1" { //Add friend
		if repository.NewRelationDaoInstance().CheckRelation(userid, touserid) {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "您已经关注过该用户"})
			return
		} else {
			relation := &repository.Relation{
				FromId: userid,
				ToId:   touserid,
			}
			service_relation.Follow(relation)

			c.JSON(http.StatusOK, Response{StatusCode: 0})
			return
		}
	} else {
		touseridstr := c.Query("to_user_id")
		touserid, _ := strconv.ParseInt(touseridstr, 10, 64)
		if !repository.NewRelationDaoInstance().CheckRelation(userid, touserid) {
			c.JSON(http.StatusOK, Response{StatusCode: 3, StatusMsg: "您已经取关该用户"})
			return
		} else {
			service_relation.UnFollow(userid, touserid)
			c.JSON(http.StatusOK, Response{StatusCode: 0})
			return
		}
	}
}

// TODO: visitor request need be impl
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

// TODO: visitor request need be impl
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
	useridstr := c.Query("user_id")
	userid, _ := strconv.ParseInt(useridstr, 10, 64)
	follow_relations, err := service_relation.GetFollowListById(userid)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "FollowList failed",
			},
			UserList: []User{DemoUser},
		})
	}

	// TODO: Friend list return type was updated, need rebuild this logic
	var respFriendList []User
	for _, relation := range *follow_relations {
		user, _ := repository.NewUserDaoInstance().GetUserById(relation.ToId)
		if repository.NewRelationDaoInstance().CheckRelation(userid, relation.FromId) {
			respFriendList = append(respFriendList, *RepoUserToCon(user))
		}
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: respFriendList,
	})
}
