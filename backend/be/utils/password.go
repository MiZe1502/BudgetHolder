package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func convertPasswordFromStringToByteArray(password string) []byte {
	return []byte(password)
}

func convertPasswordFromByteArrayToString(password []byte) string {
	return string(password)
}

// HashPasswordWithSalt prepares password to store it in db
func HashPasswordWithSalt(password string) (string, error) {
	pwd := convertPasswordFromStringToByteArray(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	return convertPasswordFromByteArrayToString(hash), err
}

// ComparePasswords compares password from user and from db
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := convertPasswordFromStringToByteArray(hashedPwd)
	bytePwd := convertPasswordFromStringToByteArray(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}

	return true
}
