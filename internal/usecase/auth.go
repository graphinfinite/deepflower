package usecase

import (
	"deepflower/internal/repository"
	"errors"
	"fmt"

	h "deepflower/internal/helpers"
	m "deepflower/internal/model"
)

type AuthUsecase struct {
	Rep UserStorageInterface
}

type UserStorageInterface interface {
	CreateUser(u m.User) (int, error)
	GetUserByTgId(tgId int) (m.User, error)
}

func NewAuthUsecase(r UserStorageInterface) AuthUsecase {
	return AuthUsecase{Rep: r}
}

// generate new username and password. save new user data.
// return user model with Username, Password.
// if user with tgId already exist -> ErrAuthUserAlreadyExist (and update chatId, userName, firstName, lastName, languageCode)
func (auth *AuthUsecase) RegistrationFromTg(tguser m.UserTelegram) (m.User, error) {
	var ErrUserNotFound repository.ErrStoreUserNotFound
	user, err := auth.Rep.GetUserByTgId(tguser.TgId)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			newusername := h.GenUserName()
			newpassword := h.GenNewPassword()
			hash, err := h.HashAndSalt([]byte(newpassword))
			if err != nil {
				return m.User{}, err
			}
			_, err = auth.Rep.CreateUser(m.User{UserTelegram: tguser, HashedPassword: hash, Username: newusername})
			if err != nil {
				return m.User{}, err
			}
			return m.User{Username: newusername, Password: newpassword}, nil
		}
	}
	//TODO: update chatId, userName, firstName, lastName, languageCode for user
	return m.User{Username: user.Username, Password: ""}, NewErrAuthUserAlreadyExist("", fmt.Errorf("user with tgid: %d already exist", tguser.TgId))

}
