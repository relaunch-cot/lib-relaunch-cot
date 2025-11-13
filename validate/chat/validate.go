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

func ValidateSendMessageRequest(m *pb.SendMessageRequest) error {
	return validation.ValidateStruct(m,
		validation.Field(&m.ChatId, validation.Required, is.UUID),
		validation.Field(&m.SenderId, validation.Required, is.UUID),
		validation.Field(&m.MessageContent, validation.Required),
	)
}

func ValidateGetAllMessagesFromChatRequest(g *pb.GetAllMessagesFromChatRequest) error {
	return validation.ValidateStruct(g,
		validation.Field(&g.ChatId, validation.Required, is.UUID),
	)
}

func ValidateGetAllChatsFromUserRequest(g *pb.GetAllChatsFromUserRequest) error {
	return validation.ValidateStruct(g,
		validation.Field(&g.UserId, validation.Required, is.UUID),
	)
}

func ValidateGetChatFromUsersRequest(g *pb.GetChatFromUsersRequest) error {
	return validation.ValidateStruct(g,
		validation.Field(&g.UserIds, validation.Required, validation.Length(2, 2), validation.By(validateUserIds)),
	)
}

func ValidateGetChatByIdRequest(g *pb.GetChatByIdRequest) error {
	return validation.ValidateStruct(g,
		validation.Field(&g.ChatId, validation.Required, is.UUID),
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
