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
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title,omitempty"`
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
		IsFavorite:    video.IsFavorite,
		Title:         video.Title,
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
	Id             int64  `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	FollowCount    int64  `json:"follow_count,omitempty"`
	FollowerCount  int64  `json:"follower_count,omitempty"`
	IsFollow       bool   `json:"is_follow,omitempty"`
	Avatar         string `json:"avatar,omitempty"`
	BackgrounImage string `json:"background_image,omitempty"`
	Signature      string `json:"signature,omitempty"`
	TotalFavorited int64  `json:"total_favorited,omitempty"`
	WorkCount      int64  `json:"work_count,omitempty"`
	FavoriteCount  int64  `json:"favorite_count,omitempty"`
}

type FriendUser struct {
	User
	Message string `json:"message,omitempty"` // 和该好友的最新聊天消息
	MsgType int64  `json:"msgType,omitempty"` // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

func RepoUserToCon(user *repository.User) *User {
	return &User{
		Id:             user.Id,
		Name:           user.Name,
		FollowCount:    user.FollowCount,
		FollowerCount:  user.FollowerCount,
		IsFollow:       user.IsFollow,
		Avatar:         user.Avatar,
		BackgrounImage: user.BackgrounImage,
		Signature:      user.Signature,
		TotalFavorited: user.TotalFavorited,
		WorkCount:      user.WorkCount,
		FavoriteCount:  user.FavoriteCount,
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
