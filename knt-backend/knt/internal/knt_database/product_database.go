package kntdatabase

import "database/sql"

func GetAllProducts(db *sql.DB) ([]Product, error) {
	return genericQuery[Product](queryBuilder(db, "select * from product"))
}

func GetProduct(db *sql.DB, productId int) ([]Product, error) {
	return genericQuery[Product](queryBuilder(db, "select * from product where id = ?", productId))
}

func GetMinimalProducts(db *sql.DB) ([]MinimalProduct, error) {
	return genericQuery[MinimalProduct](queryBuilder(db, "select * from product where visibility = 1"))
}
