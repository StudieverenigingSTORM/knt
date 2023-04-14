package kntdatabase

import (
	"database/sql"
	"log"
	"reflect"
)

func GetAllProducts(db *sql.DB) []Product {
	return genericQuery[Product](queryBuilder(db, "select * from product"))
}

func GetAllUsers(db *sql.DB) []User {
	return genericQuery[User](queryBuilder(db, "select * from user"))
}

func GetAllMinimalUsers(db *sql.DB) []MinimalUser {
	return genericQuery[MinimalUser](queryBuilder(db, "select id, first_name, last_name, balance from user"))
}

func GetUser(db *sql.DB, userID int) User {
	return getSingleEntry[User](queryBuilder(db, "select * from user where id = ?", userID))
}

func getSingleEntry[K any](rows *sql.Rows) K {
	defer rows.Close()
	var output K
	if rows.Next() {
		rows.Scan(structForScan(&output)...)
	}
	return output
}

// Generic query handles scaling through all the rows.
// You only need to define the object for it to expect.
// In most cases this object should not be an array
// Also be sure to provide the exact struct matching the query
// Failure to do so might cause undeffined problems
func genericQuery[K any](rows *sql.Rows) (output []K) {
	defer rows.Close()
	var slice []K
	for rows.Next() {
		var temp K
		rows.Scan(structForScan(&temp)...)
		slice = append(slice, temp)
	}
	return slice
}

// Okay I know this looks scary but I promise it makes sense
// What this function does it converts a struct into an interface that provides access to itself.
// This allows the databases scan function to populate the privided struct.
// Pro tip do not touch this function as its just boiler plate
func structForScan(u interface{}) []interface{} {
	val := reflect.ValueOf(u).Elem()
	v := make([]interface{}, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		v[i] = valueField.Addr().Interface()
	}
	return v
}

// Function to simplify building queries and reduce code reuse, this should be used whenever any query is made.
// Keep note do not append any data to the query string instead use the key ? and pass in aditional parameters
func queryBuilder(db *sql.DB, queryString string, args ...any) *sql.Rows {
	rows, err := db.Query(queryString, args...)
	if err != nil {
		log.Println(err)
	}
	return rows
}
