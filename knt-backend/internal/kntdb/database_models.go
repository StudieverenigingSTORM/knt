package kntdb

type MinimalProduct struct {
	Id    int    `json:"id"`
	Price int    `json:"price"`
	Name  string `json:"name"`
}

type Product struct {
	Id          int    `json:"id"`
	Price       int    `json:"price"`
	Name        string `json:"name"`
	Visibility  int    `json:"visibility"`
	TaxCategory int    `json:"taxcategory"`
}

type User struct {
	Id         int    `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	VunetId    string `json:"vunetId"`
	Password   string `json:"password"`
	Balance    int    `json:"balance"`
	Visibility int    `json:"visibility"`
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
	Id              int    `json:"id"`
	UserId          int    `json:"userId"`
	StartingBalance int    `json:"startingBalance"`
	DeltaBalance    int    `json:"deltaBalance"`
	FinalBalance    int    `json:"finalBalance"`
	ReceiptId       int    `json:"receiptId"`
	Reference       string `json:"reference"`
}

type Receipt struct {
	Id        int    `json:"id"`
	Data      string `json:"data"`
	Timestamp string `json:"timestamp"`
}

type TaxEntry struct {
	Id        int `json:"id"`
	ProductId int `json:"productId"`
	Amount    int `json:"amount"`
	TotalCost int `json:"totalCost"`
	Year      int `json:"year"`
}

// This struct is specifically designed to be returned by the paginated transaction
type TransactionDetails struct {
	Id              int    `json:"id"`
	VunetId         string `json:"vunetid"`
	StartingBalance int    `json:"startingBalance"`
	DeltaBalance    int    `json:"deltaBalance"`
	FinalBalance    int    `json:"finalBalance"`
	Reference       string `json:"reference"`
	Timestamp       string `json:"timestamp"`
	Data            string `json:"string"`
}
