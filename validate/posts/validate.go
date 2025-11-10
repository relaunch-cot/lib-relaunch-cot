package posts

import (
	"github.com/invopop/validation"
	"github.com/invopop/validation/is"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/post"
)

func ValidateCreatePostRequest(p *pb.CreatePostRequest) error {
	return validation.ValidateStruct(p,
		validation.Field(&p.UserId, validation.Required, is.UUID),
		validation.Field(&p.Title, validation.Required, validation.Length(1, 200)),
		validation.Field(&p.Content, validation.Required, validation.Length(1, 5000)),
		validation.Field(&p.Type, validation.Required), // TODO: determninar quais tipos são válidos
	)
}

func ValidateGetPostRequest(p *pb.GetPostRequest) error {
	return validation.ValidateStruct(p,
		validation.Field(&p.PostId, validation.Required, is.UUID),
	)
}

func ValidateGetAllPostsFromUserRequest(p *pb.GetAllPostsFromUserRequest) error {
	return validation.ValidateStruct(p,
		validation.Field(&p.UserId, validation.Required, is.UUID),
	)
}

func ValidateUpdatePostRequest(p *pb.UpdatePostRequest) error {
	return validation.ValidateStruct(p,
		validation.Field(&p.PostId, validation.Required, is.UUID),
		validation.Field(&p.Title, validation.Length(1, 200)),
		validation.Field(&p.Content, validation.Length(1, 5000)),
	)
}

func ValidateDeletePostRequest(p *pb.DeletePostRequest) error {
	return validation.ValidateStruct(p,
		validation.Field(&p.PostId, validation.Required, is.UUID),
	)
}

func ValidateUpdateLikesFromPostOrCommentRequest(p *pb.UpdateLikesFromPostOrCommentRequest) error {
	return validation.ValidateStruct(p,
		validation.Field(&p.UserId, validation.Required, is.UUID),
		validation.Field(&p.PostId, is.UUID),
		validation.Field(&p.CommentId, is.UUID),
		validation.Field(&p.Type, validation.Required),
	)
}

func ValidateGetAllLikesFromPostRequest(p *pb.GetAllLikesFromPostRequest) error {
	return validation.ValidateStruct(p,
		validation.Field(&p.PostId, validation.Required, is.UUID),
		validation.Field(&p.UserId, validation.Required, is.UUID),
	)
}

func ValidateCreateCommentOrReplyRequest(p *pb.CreateCommentOrReplyRequest) error {
	return validation.ValidateStruct(p,
		validation.Field(&p.UserId, validation.Required, is.UUID),
		validation.Field(&p.PostId, validation.Required, is.UUID),
		validation.Field(&p.ParentCommentId, is.UUID),
		validation.Field(&p.Content, validation.Required, validation.Length(1, 1000)),
		validation.Field(&p.Type, validation.Required, validation.In("comment", "reply")),
	)
}

func DeleteCommentOrReplyRequest(p *pb.DeleteCommentOrReplyRequest) error {
	return validation.ValidateStruct(p,
		validation.Field(&p.UserId, validation.Required, is.UUID),
		validation.Field(&p.CommentId, is.UUID),
		validation.Field(&p.ReplyId, is.UUID),
		validation.Field(&p.Type, validation.Required, validation.In("comment", "reply")),
	)
}

func GetAllCommentsFromPostRequest(p *pb.GetAllCommentsFromPostRequest) error {
	return validation.ValidateStruct(p,
		validation.Field(&p.UserId, validation.Required, is.UUID),
		validation.Field(&p.PostId, validation.Required, is.UUID),
	)
}
