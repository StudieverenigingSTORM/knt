package kntdb

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func ValidatePin(pin string, user User) bool {
	return user.Password == ShaHashing(pin)
}

// CheckUserPrivileges iterates through every logged api key and compares it to the current function
func CheckUserPrivileges(key string) string {
	hashedClientKey := ShaHashing(key)
	rows, err := DB.Query("select privileges from keys where token = ?", hashedClientKey)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	if rows.Next() {
		var privileges string
		rows.Scan(&privileges)
		return privileges
	}
	return ""
}

// Hash the password to compare it.
func ShaHashing(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
