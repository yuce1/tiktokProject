package controller

import (
	"tiktok-go/repository"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

func RepoVideoToCon(video *repository.Video) *Video {
	author, _ := repository.NewUserDaoInstance().GetUserByName(video.Author)
	return &Video{
		Id:            video.Id,
		Author:        *RepoUserToCon(author),
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		// TODO: favourite list hasnt been develop yet
		IsFavorite: video.IsFavorite,
	}
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

func RepoCommentToCon(comment *repository.Comment) *Comment {
	author, _ := repository.NewUserDaoInstance().GetUserById(comment.UserId)
	return &Comment{
		Id:         comment.Id,
		User:       *RepoUserToCon(author),
		Content:    comment.Content,
		CreateDate: comment.CreateDate,
	}
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

func RepoUserToCon(user *repository.User) *User {
	return &User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	FromUserId int64  `json:"from_user_id"`
	ToUserId   int64  `json:"to_user_id"`
	Content    string `json:"content,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
}

func RepoChatToMsg(cr *repository.ChatRecord) *Message {
	return &Message{
		Id:         cr.Id,
		FromUserId: cr.FromUserId,
		ToUserId:   cr.ToUserId,
		Content:    cr.Content,
		CreateTime: cr.CreatedAt,
	}
}

// a struct for query self info, due to add work_count & favorite_count fields
type UserInfoResp struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	WorkCount     int64  `json:"work_count"`
	FavoriteCount int64  `json:"favorite_count"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

// a func turn repo.user to UserInfoResp
func RepoUserToInfo(u *repository.User) *UserInfoResp {
	return &UserInfoResp{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		WorkCount:     u.WorkCount,
		FavoriteCount: u.FavoriteCount,
		IsFollow:      true,
	}
}

// maybe we can change the message struct to this Info struct, as same as the  struct in repo.
// and in marshal, we can choose dont marshal the work_count and favorite_count field, to match the response Message struct for response.
