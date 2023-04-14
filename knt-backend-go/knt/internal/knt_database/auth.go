package kntdatabase

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"log"
)

// Validate an api key recieved from the network.
func ValidateKey(key string, privilages string, db *sql.DB) bool {
	return CheckUserPrivileges(key, db) == privilages
}

func ValidatePin(pin string, userID int, db *sql.DB) bool {
	//TODO: Implement this function
	//return checkPasswordHash(pin, strconv.Itoa(GetUser(db, userID).Password))
	return true
}


//CheckUserPrivileges iterates through every logged api key and compares it to the current function
func CheckUserPrivileges(key string, db *sql.DB) string {
	hashedClientKey := shaHashing(key)
	rows, err := db.Query("select token, privileges from keys")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var hashedKey string
		var privilages string
		rows.Scan(&hashedKey, &privilages)
		if hashedKey == hashedClientKey {
			return privilages
		}
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
