package usecase

import (
	"context"
	"deepflower/internal/model"
)

type UserUC struct {
	Rep UserStorageInterface
}

func NewUserUC(r UserStorageInterface) *UserUC {
	return &UserUC{Rep: r}
}

func (u *UserUC) GetUserInfo(ctx context.Context, userId string) (user model.User, err error) {
	user, err = u.Rep.GetUserById(ctx, userId)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *UserUC) UpdateUser(ctx context.Context, m model.User) (user model.User, err error) {
	user, err = u.Rep.UpdateUser(ctx, m)
	if err != nil {
		return model.User{}, err
	}
	return user, nil

}
