package utils

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
)

const (
	salt          = "fkjhdfkhkfnfsdjkfnksj"
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	secretCodeLen = 12
)

func GeneratePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func GenerateRandomSecretCode() string {
	b := make([]byte, secretCodeLen)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
