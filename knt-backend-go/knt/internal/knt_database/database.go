package kntdatabase

import (
	"database/sql"
	"log"
)

func GetAllProducts(db *sql.DB) ProductList {
	var newProductList ProductList
	newProductList.Products = genericQuery(db, "select * from product", func(r *sql.Rows) Product {
		var p Product
		err := r.Scan(&p.Id, &p.Price, &p.Name)
		if err != nil {
			log.Fatal(err)
		}
		return p
	})
	return newProductList
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
