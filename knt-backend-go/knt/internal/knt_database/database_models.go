package kntdatabase

type Product struct {
	Id    int     `json:"id"`
	Price float64 `json:"price"`
	Name  string  `json:"name"`
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	VunetId   string `json:"vunetId"`
	Password  int    `json:"password"`
	Balance   int    `json:"balance"`
}

type MinimalUser struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Balance   int    `json:"balance"`
}
