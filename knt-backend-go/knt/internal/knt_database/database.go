package kntdatabase

import (
	"database/sql"
	"log"
)

func GetAllProducts(db *sql.DB) []Product {
	var l []Product
	l = genericQuery(queryBuilder(db, "select * from product"), func(r *sql.Rows) Product {
		var p Product
		err := r.Scan(&p.Id, &p.Price, &p.Name)
		if err != nil {
			log.Fatal(err)
		}
		return p
	})
	return l
}

func GetAllUsers(db *sql.DB) []User {
	var l []User
	l = genericQuery(queryBuilder(db, "select * from user"), func(r *sql.Rows) User {
		var p User
		err := r.Scan(&p.Id, &p.FirstName, &p.LastName, &p.VunetId, &p.Password, &p.Balance)
		if err != nil {
			log.Fatal(err)
		}
		return p
	})
	return l
}

func GetUserPass(db *sql.DB, userID int) int {
	return getSingleVar[int](queryBuilder(db, "select password from user where id = ?", userID))
}

func getSingleVar[K any](rows *sql.Rows) K {
	defer rows.Close()
	var output K
	if rows.Next() {
		rows.Scan(&output)
	}
	return output
}

// Generic query handles scaling through all the rows. Functions using this query have to provide
// Code that will handle each row. Appending and everything else is handles by the function
func genericQuery[K any](rows *sql.Rows, f func(*sql.Rows) K) (output []K) {
	defer rows.Close()
	var slice []K
	for rows.Next() {
		slice = append(slice, f(rows))
	}
	return slice
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
