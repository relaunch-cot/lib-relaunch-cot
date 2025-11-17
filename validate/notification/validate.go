package notification

import (
	"github.com/invopop/validation"
	"github.com/invopop/validation/is"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/notification"
)

var validTypes []string = []string{
	"PROJECT_INVITE",
	"PROJECT_REQUEST",
	"PROJECT_ACCEPTED",
	"PROJECT_REJECTED",
	"PROJECT_STARTED",
	"PROJECT_COMPLETED",
	"PROJECT_UPDATED",

	// Chat
	"CHAT_MESSAGE",

	// Pagamento
	"PAYMENT_PENDING",
	"PAYMENT_RECEIVED",
	"PAYMENT_RELEASED",
	"PAYMENT_DISPUTE",

	// Arquivos
	"FILE_UPLOADED",
	"FILE_APPROVED",
	"FILE_REJECTED",

	// Prazos
	"DEADLINE_APPROACHING",
	"DEADLINE_TODAY",
	"DEADLINE_EXPIRED",

	// Sistema
	"ACCOUNT_VERIFIED",
	"PROFILE_VIEWED",
	"REVIEW_RECEIVED",
	"SYSTEM_ANNOUNCEMENT",
	"TERMS_UPDATED",
}

func ValidateSendNotificationRequest(n *pb.SendNotificationRequest) error {
	return validation.ValidateStruct(n,
		validation.Field(&n.ReceiverId, validation.Required, is.UUID),
		validation.Field(&n.SenderId, validation.Required, is.UUID),
		validation.Field(&n.Title, validation.Required),
		validation.Field(&n.Content, validation.Required),
		validation.Field(&n.Type, validation.Required, validation.In(validTypes...)),
	)
}

func ValidateGetNotificationRequest(n *pb.GetNotificationRequest) error {
	return validation.ValidateStruct(n,
		validation.Field(&n.NotificationId, validation.Required, is.UUID),
	)
}

func ValidateGetAllNotificationsFromUserRequest(n *pb.GetAllNotificationsFromUserRequest) error {
	return validation.ValidateStruct(n,
		validation.Field(&n.UserId, validation.Required, is.UUID),
	)
}

func ValidateDeleteNotificationRequest(n *pb.DeleteNotificationRequest) error {
	return validation.ValidateStruct(n,
		validation.Field(&n.NotificationId, validation.Required, is.UUID),
	)
}

func ValidateDeleteAllNotificationsFromUserRequest(n *pb.DeleteAllNotificationsFromUserRequest) error {
	return validation.ValidateStruct(n,
		validation.Field(&n.UserId, validation.Required, is.UUID),
	)
}
