package usecase

import (
	"context"
	"deepflower/internal/repository"
	"errors"
	"fmt"
	"time"

	m "deepflower/internal/model"
	h "deepflower/internal/usecase/auth_modules"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecase struct {
	Rep            UserStorageInterface
	cost           int
	signingKey     string
	expireDuration time.Duration
}

func NewAuthUsecase(r UserStorageInterface, cost int, signingKey string, expireDuration time.Duration) *AuthUsecase {
	return &AuthUsecase{Rep: r, cost: cost, signingKey: signingKey, expireDuration: expireDuration}
}

// generate new username and password. save new user data.
// return user model with Username, Password.
// if user with tgId already exist -> ErrAuthUserAlreadyExist (and update chatId, userName, firstName, lastName, languageCode)
func (auth *AuthUsecase) RegistrationFromTg(ctx context.Context, tguser m.UserTelegram) (m.User, error) {
	var ErrUserNotFound *repository.ErrStoreUserNotFound
	userold, err := auth.Rep.GetUserByTgId(ctx, tguser.TgId)
	switch {
	case errors.As(err, &ErrUserNotFound):
		newusername := h.GenUserName()
		newpassword := h.GenNewPassword()
		hash, err := h.HashAndSalt([]byte(newpassword), auth.cost)
		if err != nil {
			return m.User{}, err
		}
		_, err = auth.Rep.CreateUser(ctx, m.User{UserTelegram: tguser, HashedPassword: hash, Username: newusername})
		if err != nil {
			return m.User{}, err
		}
		return m.User{Username: newusername, Password: newpassword}, nil
	case err != nil:
		return m.User{}, err
	default:
		return m.User{Username: userold.Username}, NewErrAuthUserAlreadyExist("", fmt.Errorf("user with tgid: %d already exist", tguser.TgId))
	}
}

type CustomClaims struct {
	X string `json:"foo"`
	jwt.RegisteredClaims
}

func (auth *AuthUsecase) Login(ctx context.Context, username, password string) (token string, err error) {
	user, err := auth.Rep.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	if ok := h.ComparePasswords(user.HashedPassword, password); !ok {
		return "", fmt.Errorf("password invalide")
	}

	claims := CustomClaims{
		"x",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(auth.expireDuration * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   fmt.Sprint(user.ID),
			ID:        "1",
			Audience:  []string{"aud"},
		},
	}
	jwttoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := jwttoken.SignedString([]byte(auth.signingKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (auth *AuthUsecase) ValidateJwtToken(ctx context.Context, tokenString string) (bool, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(auth.signingKey), nil
	})

	if err != nil {
		return false, nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return true, claims, nil
	}
	return false, nil, nil
}
