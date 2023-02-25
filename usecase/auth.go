package usecase

import (
	"deepflower/repository"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	Rep    *repository.UserStorage
	Logger *zerolog.Logger
}

func NewAuthUsecase(rep *repository.UserStorage, logger *zerolog.Logger) AuthUsecase {
	return AuthUsecase{Rep: rep, Logger: logger}
}

func genUserName() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return fmt.Sprint(r.Uint32())

}

func genNewPassword() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return fmt.Sprint(r.Uint64())

}
func hashAndSalt(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func comparePasswords(HashedPassword string, plainPassword []byte) bool {
	byteHash := []byte(HashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		return false
	}
	return true
}

type NewUser struct {
	Username string
	Password string
}

// проверка - зарегистрирован ли пользователь в базе
// если зареган отправить уведомление
// сохранить данные пользователя в базу
// сгенерировать временный пароль
// отправить данные для входа пользователю
func (auth *AuthUsecase) RegistrationFromTg(tgId int, chatId int64, userName, firstName, lastName, languageCode string) (NewUser, error) {

	var NewErrUserNotFound repository.ErrStoreUserNotFound
	user, err := auth.Rep.GetUserByTgId(tgId)
	if err != nil {
		if errors.Is(err, NewErrUserNotFound) {
			newusername := genUserName()
			newpassword := genNewPassword()
			hash, err := hashAndSalt([]byte(newpassword))
			if err != nil {
				return NewUser{}, err
			}
			_, err = auth.Rep.CreateUser(tgId, chatId, userName, firstName, lastName, languageCode, hash, newusername)
			if err != nil {
				return NewUser{}, err
			}
			return NewUser{Username: newusername, Password: newpassword}, nil

		}
	}
	return NewUser{Username: user.Username, Password: ""}, fmt.Errorf("user already exist")

}
