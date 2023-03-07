package usecase

import (
	"deepflower/internal/repository"
	"errors"
	"fmt"
	"time"

	h "deepflower/internal/helpers"
	m "deepflower/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecase struct {
	Rep            UserStorageInterface
	hashSalt       int
	signingKey     string
	expireDuration time.Duration
}

type UserStorageInterface interface {
	CreateUser(u m.User) (int, error)
	GetUserByTgId(tgId int) (m.User, error)
	GetUserByUsername(username string) (m.User, error)
}

func NewAuthUsecase(r UserStorageInterface, hashSalt int, signingKey string, expireDuration time.Duration) AuthUsecase {
	return AuthUsecase{Rep: r, hashSalt: hashSalt, signingKey: signingKey, expireDuration: expireDuration}
}

// generate new username and password. save new user data.
// return user model with Username, Password.
// if user with tgId already exist -> ErrAuthUserAlreadyExist (and update chatId, userName, firstName, lastName, languageCode)
func (auth *AuthUsecase) RegistrationFromTg(tguser m.UserTelegram) (m.User, error) {
	var ErrUserNotFound *repository.ErrStoreUserNotFound
	_, err := auth.Rep.GetUserByTgId(tguser.TgId)
	switch {
	case errors.As(err, &ErrUserNotFound):
		newusername := h.GenUserName()
		newpassword := h.GenNewPassword()
		hash, err := h.HashAndSalt([]byte(newpassword), auth.hashSalt)
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

func (auth *AuthUsecase) Login(username, password string) (token string, err error) {
	user, err := auth.Rep.GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	if ok := h.ComparePasswords(user.HashedPassword, password); !ok {
		return "", fmt.Errorf("password invalide")
	}
	jwttoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(auth.expireDuration).Unix(),
	})
	tokenString, err := jwttoken.SignedString([]byte(auth.signingKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (auth *AuthUsecase) ValidateJwtToken(tokenString string) (bool, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(auth.signingKey), nil
	})

	if err != nil {
		return false, nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	//print(claims.GetSubject())
	if ok && token.Valid {
		return true, claims, nil
	}
	return false, nil, nil
}
