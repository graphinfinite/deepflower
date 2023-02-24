package usecase

import (
	"deepflower/repository"

	"github.com/rs/zerolog"
)

type AuthUsecase struct {
	Rep    *repository.UserStorage
	Logger *zerolog.Logger
}

func NewAuthUsecase(rep *repository.UserStorage, logger *zerolog.Logger) AuthUsecase {
	return AuthUsecase{Rep: rep, Logger: logger}
}

func (auth *AuthUsecase) RegistrationFromTg(tgId int, chatId int64, userName, firstName, lastName, languageCode string) (string, error) {
	return "", nil
}
