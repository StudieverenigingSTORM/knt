package kntdatabase

import (
	"database/sql"
	"log"
)

type product struct {
	id    int
	price float64
	name  string
}

func GetAllProducts(db *sql.DB) []product {
	return genericQuery(db, "select * from product", func(r *sql.Rows) product {
		var p product
		err := r.Scan(&p.id, &p.price, &p.name)
		if err != nil {
			log.Fatal(err)
		}
		return p
	})

}

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
