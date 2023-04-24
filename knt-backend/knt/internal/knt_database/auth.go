package kntdatabase

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"log"
)

func ValidatePin(pin string, user User, db *sql.DB) bool {
	return user.Password == shaHashing(pin)
}

// CheckUserPrivileges iterates through every logged api key and compares it to the current function
func CheckUserPrivileges(key string, db *sql.DB) string {
	hashedClientKey := shaHashing(key)
	rows, err := db.Query("select privileges from keys where token = ?", hashedClientKey)
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
func shaHashing(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
