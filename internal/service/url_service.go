package service

import (
	"math/rand"
	"time"

	"url-shortener/internal/repository"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func CreateShortURL(originalURL string) (string, error) {
	shortCode := generateCode(6) 

	err := repository.SaveURL(shortCode, originalURL)
	if err != nil {
		return "", err
	}

	return shortCode, nil
}

func GetOriginalURL(code string) (string, error) {
	return repository.Get(code)
}
