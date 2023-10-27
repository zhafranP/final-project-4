package user_repository

import (
	"finalProject4/dto"
	"finalProject4/entity"
	"finalProject4/pkg/errs"
)

type Repository interface {
	CreateUser(u dto.NewUserRequest) (*dto.NewUserResponse, errs.Error)
	Login(email string) (*entity.User, errs.Error)
	TopUp(u dto.TopUpRequest) (int, errs.Error)
	CountEmail(email string) (int, errs.Error)
	GetBalance(id int) (int, errs.Error)
}
