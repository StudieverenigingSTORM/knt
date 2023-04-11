package kntdatabase

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func ValidateKey(key string, priviliges string, db *sql.DB) bool {
	return priviliges == matchKey(key, db)
}

// matchKey compares the key with hashed keys in the database
// Returns the privilage level denoted as user and admin
func matchKey(key string, db *sql.DB) string {
	rows, err := db.Query("select token, privilages from keys")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var hashedKey string
		var privilages string
		rows.Scan(&hashedKey, &privilages)
		if CheckPasswordHash(key, hashedKey) {
			return privilages
		}
	}
	return ""
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
