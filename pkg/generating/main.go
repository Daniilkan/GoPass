package generating

import (
	"math/rand"
	"time"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialBytes = "!@#$%^&*()_+-=|;':\",.<>/?`~"
	numBytes     = "0123456789"
)

func GeneratePassword(length int) (string, error) {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	for i := range b {
		symbol := rand.Intn(3) + 1
		if symbol == 1 {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		} else if symbol == 2 {
			b[i] = specialBytes[rand.Intn(len(specialBytes))]
		} else if symbol == 3 {
			b[i] = numBytes[rand.Intn(len(numBytes))]
		}
	}
	return string(b), nil
}
