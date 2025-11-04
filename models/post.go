package models

import "time"

type Post struct {
	PostId       string
	AuthorId     string
	AuthorName   string
	Title        string
	Content      string
	Type         string
	UrlImagePost string
	CreatedAt    string
	UpdatedAt    string
}

type PostLikes struct {
	LikesCount int64
	Likes      []Like
}

type Like struct {
	UserId   string
	UserName string
	LikedAt  *time.Time
}

type PostComments struct {
	CommentsCount int64
	Comments      []Comment
}

type Comment struct {
	CommentId string
	UserId    string
	UserName  string
	Content   string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
