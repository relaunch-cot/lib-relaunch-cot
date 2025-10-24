package notification

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

func IsValidNotificationType(notificationType string) bool {
	for _, v := range validTypes {
		if v == notificationType {
			return true
		}
	}

	return false
}
