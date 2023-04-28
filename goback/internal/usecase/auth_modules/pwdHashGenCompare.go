package auth_modules

import (
	"math/rand"
	"time"

	"github.com/goombaio/namegenerator"
	"golang.org/x/crypto/bcrypt"
)

func GenUserName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	return nameGenerator.Generate()
}

func GenNewPassword() string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials
	length := 16
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	str := string(buf)
	return str
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
	return err == nil
}
