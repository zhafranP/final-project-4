package service

import (
	"finalProject4/dto"
	"finalProject4/pkg/errs"
	"finalProject4/pkg/helpers"
	"finalProject4/repository/user_repository"
	"strconv"
)

type userService struct {
	userRepo user_repository.Repository
}

type UserService interface {
	CreateUser(u dto.NewUserRequest) (*dto.NewUserResponse, errs.Error)
	Login(u dto.LoginRequest) (*dto.LoginResponse, errs.Error)
	Topup(u dto.TopUpRequest) (*dto.TopUpResponse, errs.Error)
}

func NewUserService(userRepo user_repository.Repository) UserService {
	return &userService{userRepo: userRepo}
}

func (us *userService) CreateUser(u dto.NewUserRequest) (*dto.NewUserResponse, errs.Error) {
	validateErr := helpers.ValidateStruct(u)
	if validateErr != nil {
		return nil, validateErr
	}

	count, err := us.userRepo.CountEmail(u.Email)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errs.NewBadRequest("Email Has Already Been Used")
	}

	generatePW, errGenerate := helpers.GenerateHashedPassword([]byte(u.Password))
	if errGenerate != nil {
		return nil, errs.NewInternalServerError(errGenerate.Error())
	}

	u.Password = generatePW
	resp, err := us.userRepo.CreateUser(u)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (us *userService) Login(u dto.LoginRequest) (*dto.LoginResponse, errs.Error) {
	var resp dto.LoginResponse

	validateErr := helpers.ValidateStruct(u)
	if validateErr != nil {
		return nil, validateErr
	}

	user, err := us.userRepo.Login(u.Email)
	if err != nil {
		return nil, err
	}

	comparePassword := helpers.ComparePass([]byte(user.Password), []byte(u.Password))
	if comparePassword {
		token, err := helpers.GenerateToken(user.ID, u.Email)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}

		resp.Token = token
		return &resp, nil
	}

	return nil, errs.NewUnauthenticatedError("Invalid Email or Password")

}

func (us *userService) Topup(u dto.TopUpRequest) (*dto.TopUpResponse, errs.Error) {
	var resp dto.TopUpResponse

	currentBalance, err := us.userRepo.GetBalance(u.ID)
	if err != nil {
		return nil, err
	}

	if u.Balance < 0 {
		return nil, errs.NewBadRequest("invalid topup")
	}
	if currentBalance+u.Balance > 200000000 {
		return nil, errs.NewBadRequest("Your balance cannot exceed 200000000")
	}

	topup, err := us.userRepo.TopUp(u)
	if err != nil {
		return nil, err
	}

	resp.Message = "Your balance has been successfully updated to Rp." + strconv.Itoa(topup)

	return &resp, nil
}
