package usecase

import "deepflower/internal/model"

type UserUC struct {
	Rep UserStorageInterface
}

func NewUserUC(r UserStorageInterface) UserUC {
	return UserUC{Rep: r}
}

func (u *UserUC) GetUserInfo(userId string) (user model.User, err error) {
	return model.User{}, nil

}

func (u *UserUC) UpdateUser(userId string) (user model.User, err error) {
	return model.User{}, nil

}
