package usecase

import (
	"deepflower/internal/repository"
	"errors"
	"fmt"

	h "deepflower/internal/helpers"
	m "deepflower/internal/model"
)

type AuthUsecase struct {
	Rep *repository.UserStorage
}

func NewAuthUsecase(rep *repository.UserStorage) AuthUsecase {
	return AuthUsecase{Rep: rep}
}

// generate new username and password. save new user data.
// return user model with Username, Password.
// if user with tgId already exist -> ErrAuthUserAlreadyExist (and update chatId, userName, firstName, lastName, languageCode)
func (auth *AuthUsecase) RegistrationFromTg(tgId int, chatId int64, userName, firstName, lastName, languageCode string) (m.User, error) {
	var NewErrUserNotFound repository.ErrStoreUserNotFound
	user, err := auth.Rep.GetUserByTgId(tgId)
	if err != nil {
		if errors.Is(err, NewErrUserNotFound) {
			newusername := h.GenUserName()
			newpassword := h.GenNewPassword()
			hash, err := h.HashAndSalt([]byte(newpassword))
			if err != nil {
				return m.User{}, err
			}
			_, err = auth.Rep.CreateUser(tgId, chatId, userName, firstName, lastName, languageCode, hash, newusername)
			if err != nil {
				return m.User{}, err
			}
			return m.User{Username: newusername, Password: newpassword}, nil
		}
	}
	//TODO: update chatId, userName, firstName, lastName, languageCode for user
	return m.User{Username: user.Username, Password: ""}, NewErrAuthUserAlreadyExist("", fmt.Errorf("user with tgid: %d already exist", tgId))

}
