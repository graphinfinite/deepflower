package usecase

import "deepflower/internal/model"

type UserUC struct {
	Rep UserStorageInterface
}

func NewUserUC(r UserStorageInterface) UserUC {
	return UserUC{Rep: r}
}

func (u *UserUC) GetUserInfo(userId string) (user model.User, err error) {
	user, err = u.Rep.GetUserById(userId)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *UserUC) UpdateUser(m model.User) (user model.User, err error) {
	user, err = u.Rep.UpdateUser(m)
	if err != nil {
		return model.User{}, err
	}
	return user, nil

}
