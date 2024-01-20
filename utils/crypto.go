package utils

import "golang.org/x/crypto/bcrypt"

const cost = 10

func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}

func ComparePasswords(encryptedPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(plainPassword))
	return err == nil
}
