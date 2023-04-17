package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	var first string
	fmt.Println("Enter key to be hashed (Leave blank for random 32 key): ")
	fmt.Scanln(&first)
	var token string
	if first != "" {
		token = first
	} else {
		token = GenerateSecureToken()
	}
	fmt.Println("Key to be hashed: " + token)
	hashedToken := shaHashing(token)
	fmt.Println("Hashed key: ", hashedToken)
}

func GenerateSecureToken() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func shaHashing(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
