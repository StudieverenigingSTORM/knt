package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var first int
	fmt.Println("Enter key length (32 for the knt): ")
	fmt.Scanln(&first)

	token := GenerateSecureToken(first)

	fmt.Println(token)
	hashedToken, _ := HashPassword(token)
	fmt.Println(hashedToken)
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
