package chat

import (
	"github.com/invopop/validation"
	"github.com/invopop/validation/is"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/chat"
)

func ValidateCreateNewChatRequest(c *pb.CreateNewChatRequest) error {
	return validation.ValidateStruct(c,
		validation.Field(&c.UserIds, validation.Required, validation.Length(2, 2), validation.By(validateUserIds)),
	)
}

func ValidateSendMessageRequest(c *pb.SendMessageRequest) error {
	return validation.ValidateStruct(c,
		validation.Field(&c.ChatId, validation.Required, is.UUID),
		validation.Field(&c.SenderId, validation.Required, is.UUID),
		validation.Field(&c.MessageContent, validation.Required),
	)
}

func ValidateGetAllMessagesFromChatRequest(c *pb.GetAllMessagesFromChatRequest) error {
	return validation.ValidateStruct(c,
		validation.Field(&c.ChatId, validation.Required, is.UUID),
	)
}

func ValidateGetAllChatsFromUserRequest(c *pb.GetAllChatsFromUserRequest) error {
	return validation.ValidateStruct(c,
		validation.Field(&c.UserId, validation.Required, is.UUID),
	)
}

func ValidateGetChatFromUsersRequest(c *pb.GetChatFromUsersRequest) error {
	return validation.ValidateStruct(c,
		validation.Field(&c.UserIds, validation.Required, validation.Length(2, 2), validation.By(validateUserIds)),
		validation.Field(&c.UserId, validation.Required, is.UUID),
	)
}

func ValidateGetChatByIdRequest(c *pb.GetChatByIdRequest) error {
	return validation.ValidateStruct(c,
		validation.Field(&c.ChatId, validation.Required, is.UUID),
		validation.Field(&c.UserId, validation.Required, is.UUID),
	)
}

func validateUserIds(value interface{}) error {
	userIds := value.([]string)
	for _, id := range userIds {
		if err := validation.Validate(id, validation.Required, is.UUID); err != nil {
			return err
		}
	}
	return nil
}
