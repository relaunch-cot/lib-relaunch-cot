package models

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
	Type     string
	LikedAt  string
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
	Type      string
	replies   PostComments
	likes     PostLikes
	CreatedAt string
	UpdatedAt string
}
