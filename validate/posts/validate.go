package posts

func ValidateCommentAndLikeFields(commentType, likeType string) bool {
	if commentType != "" && (commentType == "comment" || commentType == "replyToComment") {
		return true
	}
	if likeType != "" && (likeType == "likeToPost" || likeType == "likeToComment") {
		return true
	}

	return false
}
