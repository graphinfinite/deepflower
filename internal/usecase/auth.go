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
	fmt.Printf("\nREGISTRATIO USECASE %d \n", tguser.TgId)
	_, err := auth.Rep.GetUserByTgId(tguser.TgId)

	fmt.Printf("%T, %s", err, err.Error())
	switch {
	case errors.Is(err, ErrUserNotFound):
		fmt.Print("ErrUserNotFound")
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
	case err != nil:
		return m.User{}, err
	default:
		return m.User{}, NewErrAuthUserAlreadyExist("", fmt.Errorf("user with tgid: %d already exist", tguser.TgId))
	}
}
