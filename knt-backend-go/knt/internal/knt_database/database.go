package kntdatabase

import (
	"database/sql"
	"log"
)

func GetAllProducts(db *sql.DB) ProductList {
	var l ProductList
	l.Products = genericQuery(db, "select * from product", func(r *sql.Rows) Product {
		var p Product
		err := r.Scan(&p.Id, &p.Price, &p.Name)
		if err != nil {
			log.Fatal(err)
		}
		return p
	})
	return l
}

func GetAllUsers(db *sql.DB) UserList {
	var l UserList
	l.Users = genericQuery(db, "select * from user", func(r *sql.Rows) User {
		var p User
		err := r.Scan(&p.Id, &p.FirstName, &p.LastName, &p.VunetId, &p.Password, &p.Balance)
		if err != nil {
			log.Fatal(err)
		}
		return p
	})
	return l
}

//Generic query handles scaling through all the rows. Functions using this query have to provide
//Code that will handle each row. Appending and everything else is handles by the function
func genericQuery[K any](db *sql.DB, queryString string, f func(*sql.Rows) K) (output []K) {
	rows, err := db.Query(queryString)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var slice []K
	for rows.Next() {
		slice = append(slice, f(rows))
	}
	return slice
}
