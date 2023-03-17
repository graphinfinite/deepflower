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
func HashAndSalt(password []byte, cost int) (string, error) {

	hash, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(HashedPassword, plainPassword string) bool {
	byteHash := []byte(HashedPassword)
	bplainPassword := []byte(plainPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, bplainPassword)
	if err != nil {
		return false
	}
	return true
}
