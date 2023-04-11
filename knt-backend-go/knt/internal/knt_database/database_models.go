package kntdatabase

type ProductList struct {
	Products []Product `json:"products"`
}

type Product struct {
	Id    int     `json:"id"`
	Price float64 `json:"price"`
	Name  string  `json:"name"`
}

type UserList struct {
	Users []User `json:"users"`
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	VunetId   string `json:"vunetId"`
	Password  string `json:"password"`
	Balance   int    `json:"balance"`
}
