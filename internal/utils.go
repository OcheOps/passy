package internal

import (
    "crypto/rand"
    "math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"

func GeneratePassword(length int) string {
    password := make([]byte, length)
    for i := range password {
        num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
        password[i] = charset[num.Int64()]
    }
    return string(password)
}
