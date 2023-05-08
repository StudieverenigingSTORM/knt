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

	productOld, err := GetProduct(product.Id)
	if err != nil {
		return 0, err
	}
	finalProduct := ModifyProduct(product, productOld)

	return commitTransaction(
		"update product set price = ?, name = ?, visibility = ?, taxcategory = ? where id = ?",
		finalProduct.Price, finalProduct.Name, finalProduct.Visibility, finalProduct.TaxCategory, finalProduct.Id,
	)
}

// Returns a user object made from the old user and changes in the request
func ModifyProduct(new Product, old Product) Product {
	if new.Name != "" {
		old.Name = new.Name
	}

	old.TaxCategory = new.TaxCategory
	old.Visibility = new.Visibility

	return old
}
