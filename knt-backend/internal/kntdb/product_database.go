package kntdb

import (
	"errors"
)

func GetAllProducts() ([]Product, error) {
	return genericQuery[Product]("select * from product")
}

func GetProduct(productId int) (Product, error) {
	product, err := getFirstEntry[Product]("select * from product where id = ?", productId)
	if product.Id == 0 {
		return product, errors.New("Product not found")
	}
	return product, err
}

func GetMinimalProducts() ([]MinimalProduct, error) {
	return genericQuery[MinimalProduct]("select id, price, name from product where visibility = 1")
}

func CreateNewProduct(product Product) (int64, error) {
	return commitTransaction(
		"insert into product (price, name, visibility, taxcategory) VALUES (?, ?, ?, ?)",
		product.Price, product.Name, product.Visibility, product.TaxCategory)
}

func UpdateProduct(product Product) (int64, error) {
	if product.Id == 0 {
		return 0, errors.New("invalid user")
	}
	return commitTransaction(
		"update product set name = ?, visibility = ?, taxcategory = ? where id = ?",
		product.Name, product.Visibility, product.TaxCategory, product.Id,
	)
}

func GetTaxCategories() ([]TaxCategory, error) {
	return genericQuery[TaxCategory]("select * from tax_categories")
}
