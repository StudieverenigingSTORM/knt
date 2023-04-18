package kntdatabase

type Product struct {
	Id    int    `json:"id"`
	Price int    `json:"price"`
	Name  string `json:"name"`
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	VunetId   string `json:"vunetId"`
	Password  string `json:"password"`
	Balance   int    `json:"balance"`
}

type MinimalUser struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Balance   int    `json:"balance"`
}

type PurchaseRequest struct {
	Password string          `json:"password"`
	Data     []PurchaseEntry `json:"data"`
}

type PurchaseEntry struct {
	ProductId int `json:"productId"`
	Amount    int `json:"amount"`
}

type Transaction struct {
	Id              int
	UserId          int
	StartingBalance int
	DeltaBalance    int
	FinalBalance    int
	ReceiptId       int
}

type Receipt struct {
	Id   int
	Data string
}

type TaxEntry struct {
	Id        int
	ProductId int
	Amount    int
	TotalCost int
}
