package pswd

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// GenerateHashPwd 生成hash密码
func GenerateHashPwd(plainText string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashPass), nil
}

// CheckPassword 检查密码是否正确
func CheckPassword(plainText string, hashPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(plainText))
}
