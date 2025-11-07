package posts

import (
	"net/http"

	"google.golang.org/grpc/status"
)

func ValidateCommentAndLikeFields(commentType, likeType string) error {
	if commentType != "" && commentType != "comment" && commentType != "replyToComment" {
		return status.Error(http.StatusBadRequest, "the comment type is invalid. must be 'comment' or 'replyToComment'")
	}
	if likeType != "" && likeType != "likeToPost" && likeType != "likeToComment" {
		return status.Error(http.StatusBadRequest, "the like type is invalid. must be 'likeToPost' or 'likeToComment'")
	}

	return nil
}
