package controller

import (
	"tiktok-go/repository"
	"time"
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
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

func RepoChatToMsg(cr *repository.ChatRecord) *Message {
	return &Message{
		Id:         cr.Id,
		Content:    cr.Content,
		CreateTime: time.Unix(cr.CreatedAt, 0).Format("2006-01-02 15:04:05"),
	}
}
