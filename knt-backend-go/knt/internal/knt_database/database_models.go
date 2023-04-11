package kntdatabase

type ProductList struct {
	Products []Product `json:"Products"`
}

type Product struct {
	Id    int     `json:"id"`
	Price float64 `json:"price"`
	Name  string  `json:"name"`
}
