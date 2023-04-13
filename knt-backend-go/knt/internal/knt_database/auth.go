package kntdatabase

import (
	"database/sql"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// Validate an api key recieved from the network.
func ValidateKey(key string, privilages string, db *sql.DB) bool {
	return checkUserPrivilages(key, db) == privilages
}

func ValidatePin(pin string, userID int, db *sql.DB) bool {
	return checkPasswordHash(pin, strconv.Itoa(GetUserPass(db, userID)))
}

func checkUserPrivilages(key string, db *sql.DB) string {
	rows, err := db.Query("select token from keys")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var hashedKey string
		var privilages string
		rows.Scan(&hashedKey, &privilages)
		if checkPasswordHash(key, hashedKey) {
			return privilages
		}
	}
	return ""
}

// Checks if the password is correct is correct.
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
