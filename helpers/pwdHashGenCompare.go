package helpers

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenUserName() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return fmt.Sprint(r.Uint32())

}

func GenNewPassword() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return fmt.Sprint(r.Uint64())

}
func HashAndSalt(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(HashedPassword string, plainPassword []byte) bool {
	byteHash := []byte(HashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		return false
	}
	return true
}
