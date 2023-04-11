package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	var first int
	fmt.Println("Enter key length (32 for the knt): ")
	fmt.Scanln(&first)

	fmt.Println(GenerateSecureToken(first))
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
