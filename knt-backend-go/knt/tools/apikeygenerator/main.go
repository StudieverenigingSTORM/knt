package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	var first int
	fmt.Println("Enter key length (32 for the knt): ")
	fmt.Scanln(&first)

	token := GenerateSecureToken(first)

	fmt.Println(token)
	hashedToken := shaHashing(token)
	fmt.Println(hashedToken)
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
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
