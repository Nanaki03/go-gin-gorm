package crypto

import (
    "golang.org/x/crypto/bcrypt"
)

// 暗号化用
func PasswordEncrypt(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hash), err
}

// パスワードの比較用
func CompareHashAndPassword(hash, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
