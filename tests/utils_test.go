package tests

import (
	"strings"
	"testing"
	"url-shortener/utils"
)

func TestGenerateRandomShortURL(t *testing.T) {
    expectedLength := 6

    shortURL := utils.GenerateRandomShortURL(expectedLength)

    if len(shortURL) != expectedLength {
        t.Errorf("Expected short URL length of %d, but got %d", expectedLength, len(shortURL))
    }

    validCharset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    for _, char := range shortURL {
        if !strings.Contains(validCharset, string(char)) {
            t.Errorf("Generated short URL contains invalid character: %s", string(char))
        }
    }
}
