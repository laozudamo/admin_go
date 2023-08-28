package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func GetPwd(pwd string) (string, string) {
	salt := generateSalt()
	passwordWithSalt := []byte(pwd + salt)
	hash := sha256.Sum256(passwordWithSalt)
	hashedPassword := base64.URLEncoding.EncodeToString(hash[:])
	return hashedPassword, salt
}

func generateSalt() string {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(salt)
}

func CheckPwd(pwd string, salt string, hashedPwd string) bool {
	passwordWithSalt := []byte(pwd + salt)
	hash := sha256.Sum256(passwordWithSalt)
	hashedPassword := base64.URLEncoding.EncodeToString(hash[:])
	return hashedPassword == hashedPwd
}
