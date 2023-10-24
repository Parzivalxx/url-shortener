package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomShortURL(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

    shortURL := make([]byte, length)
    for i := range shortURL {
        shortURL[i] = charset[seededRand.Intn(len(charset))]
    }

    return string(shortURL)
}
