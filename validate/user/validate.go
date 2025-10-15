package user

import (
	"net/http"

	"github.com/relaunch-cot/lib-relaunch-cot/models"
	"google.golang.org/grpc/status"
)

func ValidateUserSettings(userSettings models.UserSettings) error {
	if userSettings.DateOfBirth == "" {
		return status.Error(http.StatusBadRequest, "DateOfBirth is required")
	}
	if userSettings.Cpf == "" {
		return status.Error(http.StatusBadRequest, "Cpf required")
	}
	if userSettings.Phone == "" {
		return status.Error(http.StatusBadRequest, "Phone required")
	}
	return nil
}
