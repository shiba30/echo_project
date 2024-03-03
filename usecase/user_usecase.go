// プリケーション固有のビジネスロジック
// CRUD操作の制御を実装

package usecase

import (
	"twitter_echo_backend/model"
	"twitter_echo_backend/repository"
	"twitter_echo_backend/validator"
)

type UserUsecase interface {
	CreateUser(user model.User) (model.User, error)
	GetAllUsers() ([]model.User, error)
	GetUserById(userId uint) (model.User, error)
	UpdateUser(user model.User, userId uint) (model.User, error)
	DeleteUser(userId uint) error
}

type userUsecase struct {
	ur repository.UserRepository
	uv validator.UserValidator
}

// コンストラクタの作成、main.goで呼び出し
func NewUserUsecase(ur repository.UserRepository, uv validator.UserValidator) UserUsecase {
	return &userUsecase{ur, uv}
}

// サインアップ機能
func (uu *userUsecase) CreateUser(user model.User) (model.User, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.User{}, err
	}
	if err := uu.ur.CreateUser(&user); err != nil {
		return model.User{}, err
	}
	resUser := model.User{
		DisplayName: user.DisplayName,
		Password:    user.Password,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
	return resUser, nil
}

func (uu *userUsecase) GetAllUsers() ([]model.User, error) {
	users := []model.User{}
	if err := uu.ur.GetAllUsers(&users); err != nil {
		return nil, err
	}
	resUsers := []model.User{}
	for _, v := range users {
		t := model.User{
			ID:        v.ID,
			Email:     v.Email,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resUsers = append(resUsers, t)
	}
	return resUsers, nil
}

func (uu *userUsecase) GetUserById(userId uint) (model.User, error) {
	user := model.User{}
	if err := uu.ur.GetUserById(&user, userId); err != nil {
		return model.User{}, err
	}
	resUser := model.User{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return resUser, nil
}

func (uu *userUsecase) UpdateUser(user model.User, userId uint) (model.User, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.User{}, err
	}
	if err := uu.ur.UpdateUser(&user, userId); err != nil {
		return model.User{}, err
	}
	resUser := model.User{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return resUser, nil
}

func (uu *userUsecase) DeleteUser(userId uint) error {
	if err := uu.ur.DeleteUser(userId); err != nil {
		return err
	}
	return nil
}
