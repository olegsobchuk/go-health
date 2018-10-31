package secret

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash make crypt string
func Hash(str string) (string, error) {
	bytes := []byte(str)
	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Check checks if encrypted hash match with needed string
func Check(hash, str string) bool {
	byteHash := []byte(hash)
	byteStr := []byte(str)
	err := bcrypt.CompareHashAndPassword(byteHash, byteStr)
	if err != nil {
		return false
	}
	return true
}
