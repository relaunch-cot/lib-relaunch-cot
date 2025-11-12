package user

import (
	"github.com/invopop/validation"
	"github.com/invopop/validation/is"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/user"
)

func ValidateCreateUserRequest(u *pb.CreateUserRequest) error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 100)),
	)

	if err != nil {
		return err
	}

	err = validation.ValidateStruct(u.Settings,
		validation.Field(&u.Settings.DateOfBirth, validation.Required),
		validation.Field(&u.Settings.Cpf, validation.Required),
		validation.Field(&u.Settings.Phone, validation.Required),
	)

	return err
}

func ValidateLoginUserRequest(u *pb.LoginUserRequest) error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required),
	)
}

func ValidateGetUserRequest(u *pb.GetUserProfileRequest) error {
	return validation.ValidateStruct(u,
		validation.Field(&u.UserId, validation.Required, is.UUID),
	)
}

func UpdateUserPasswordRequest(u *pb.UpdateUserPasswordRequest) error {
	return validation.ValidateStruct(u,
		validation.Field(&u.UserId, validation.Required, is.UUID),
		validation.Field(&u.NewPassword, validation.Required, validation.Length(8, 100)),
	)
}

func ValidateUpdateUserRequest(u *pb.UpdateUserRequest) error {
	err := validation.ValidateStruct(u,
		validation.Field(&u.UserId, validation.Required, is.UUID),
	)

	if err != nil {
		return err
	}

	err = validation.ValidateStruct(u,
		validation.Field(&u.NewUser.Email, is.Email),
	)
	return err
}

func ValidateDeleteUserRequest(u *pb.DeleteUserRequest) error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
	)
}

func ValidateGenerateReportRequest(u *pb.GenerateReportRequest) error {
	return validation.ValidateStruct(u,
		validation.Field(&u.JsonData, validation.Required, is.JSON),
	)
}

func ValidateSendPasswordRecoveryEmailRequest(u *pb.SendPasswordRecoveryEmailRequest) error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.RecoveryLink, validation.Required),
	)
}
