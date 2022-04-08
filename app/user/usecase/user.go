package usecase

import (
	"umu/golang-api/app/domain"
	"umu/golang-api/common"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(user domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: user,
	}
}

func (u *userUsecase) GetUserByCriteria(search domain.SearchCriteria) (user *domain.User, err error) {
	user, err = u.userRepo.Get(search)
	return
}

func (u *userUsecase) Register(request *domain.UserRegisterRequest) (user *domain.User, err error) {
	search := domain.SearchCriteria{
		Username: &request.Username,
		Email:    &request.Email,
		Phone:    &request.Phone,
	}
	user, _ = u.GetUserByCriteria(search)
	if user != nil { // user already exist
		err = common.ErrUserAlreadyExist
		return nil, err
	}
	user = request.FillModel()
	user, err = u.userRepo.Put(user)
	return
}

func (u *userUsecase) Login(request *domain.UserLoginRequest) (user *domain.User, err error) {
	search := domain.SearchCriteria{
		Username: &request.Username,
	}
	user, _ = u.GetUserByCriteria(search)

	if user == nil {
		err = common.ErrUserNotFound
		return nil, err
	}

	err = common.ComparePassword([]byte(request.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	return
}
