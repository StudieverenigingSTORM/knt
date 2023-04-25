package kntdatabase

import (
	"database/sql"
	"reflect"
)

func GetAllProducts(db *sql.DB) ([]Product, error) {
	return genericQuery[Product](queryBuilder(db, "select * from product"))
}

func GetProduct(db *sql.DB, productId int) ([]Product, error) {
	return genericQuery[Product](queryBuilder(db, "select * from product where id = ?", productId))
}

func GetMinimalProducts(db *sql.DB) ([]MinimalProduct, error) {
	return genericQuery[MinimalProduct](queryBuilder(db, "select * from product where visibility = 1"))
}

func GetAllUsers(db *sql.DB) ([]User, error) {
	return genericQuery[User](queryBuilder(db, "select * from user"))
}

func GetAllMinimalUsers(db *sql.DB) ([]MinimalUser, error) {
	return genericQuery[MinimalUser](queryBuilder(db, "select id, first_name, last_name, balance from user where visibility = 1"))
}

func GetMinimalUser(db *sql.DB, userId int) (MinimalUser, error) {
	return getFirstEntry[MinimalUser](queryBuilder(db, "select id, first_name, last_name, balance from user where id = ? and visibility = 1", userId))
}

func GetUser(db *sql.DB, userID int) (User, error) {
	return getFirstEntry[User](queryBuilder(db, "select * from user where id = ?", userID))
}

func CreateNewUser(db *sql.DB, user User) (int64, error) {
	return commitTransaction(db,
		"insert into user (first_name, last_name, vunetid, password, balance, visibility) VALUES (?, ?, ?, ?, ?, ?)",
		user.FirstName, user.LastName, user.VunetId, user.Password, user.Balance, user.Visibility)
}

// Returns a single entry in a specific structure
func getFirstEntry[K any](rows *sql.Rows, err error) (K, error) {
	var output K
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
func getFirstSingleValue[K any](rows *sql.Rows, err error) (K, error) {
	var output K
	if err != nil {
		return output, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&output)
	}
	return output, nil
}

// Generic query handles scaling through all the rows.
// You only need to define the object for it to expect.
// In most cases this object should not be an array
// Also be sure to provide the exact struct matching the query
// Failure to do so might cause undeffined problems
func genericQuery[K any](rows *sql.Rows, err error) ([]K, error) {
	var slice []K
	if err != nil {
		return slice, err
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

// Function to simplify building queries and reduce code reuse, this should be used whenever any query is made.
// Keep note do not append any data to the query string instead use the key ? and pass in aditional parameters
func queryBuilder(db *sql.DB, queryString string, args ...any) (*sql.Rows, error) {
	rows, err := db.Query(queryString, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// This allows inserting new rows into the table
// As always when passing args to query string pass it as aditional parameters
// Do NOT concatinate them as a string as that will make it vulnerable to exploits
func commitTransaction(db *sql.DB, queryString string, args ...any) (int64, error) {
	transaction, err := db.Prepare(queryString)
	if err != nil {
		return 0, err
	}

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
