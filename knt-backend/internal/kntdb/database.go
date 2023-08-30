package kntdb

import (
	"database/sql"
	"errors"
	"reflect"

	"github.com/google/logger"
	"github.com/spf13/viper"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", viper.GetString("database"))
	if err != nil {
		logger.Fatal(err)
	}
}

// Returns a single entry in a specific structure
func getFirstEntry[K any](queryString string, args ...any) (K, error) {
	var output K
	rows, err := DB.Query(queryString, args...)

	if err != nil {
		return output, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(structForScan(&output)...)
	}
	return output, nil
}

// generic query that returns the first row of a single column query
func getFirstSingleValue[K any](queryString string, args ...any) (K, error) {
	var output K
	rows, err := DB.Query(queryString, args...)
	if err != nil {
		return output, err
	}
	defer rows.Close()
	outExists := false
	if rows.Next() {
		outExists = true
		rows.Scan(&output)
	}
	if !outExists {
		return output, errors.New("no values returned")
	}
	return output, nil
}

// Generic query handles scaling through all the rows.
// You only need to define the object for it to expect.
// In most cases this object should not be an array
// Also be sure to provide the exact struct matching the query
// Failure to do so might cause undeffined problems
func genericQuery[K any](queryString string, args ...any) ([]K, error) {
	var slice []K
	rows, err := DB.Query(queryString, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var temp K
		rows.Scan(structForScan(&temp)...)
		slice = append(slice, temp)
	}
	return slice, nil
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

// This allows inserting new rows into the table
// As always when passing args to query string pass it as aditional parameters
// Do NOT concatinate them as a string as that will make it vulnerable to exploits
func commitTransaction(queryString string, args ...any) (int64, error) {
	transaction, err := DB.Prepare(queryString)
	if err != nil {
		return 0, err
	}
	defer transaction.Close()
	res, err := transaction.Exec(args...)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Add to transaction
// This function adds to an already started transaction
// Incase any error occurs this function will automatically rollback the transaction
func addToTransaction(transaction *sql.Tx, queryString string, args ...any) (int64, error) {
	res, err := transaction.Exec(queryString, args...)
	if err != nil {
		transaction.Rollback()
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		transaction.Rollback()
		return 0, err
	}
	return id, nil
}

func AddAdminLogs(route string, method string, body string, admin string) error {
	_, err := commitTransaction("insert into admin_log (admin, route, method, data, timestamp) VALUES (?, ?, ?, ?, datetime())", admin, route, method, body)
	return err
}
