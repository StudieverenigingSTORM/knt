package kntdb

type MinimalProduct struct {
	Id    int    `json:"id"`
	Price int    `json:"price"`
	Name  string `json:"name"`
}

type Product struct {
	Id          int    `json:"id"`
	Price       int    `json:"price"`
	Name        string `json:"name" validate:"required,min=1,max=32"`
	Visibility  int    `json:"visibility" validate:"max=1"`
	TaxCategory int    `json:"taxcategory" validate:"required,min=1"`
}

type User struct {
	Id         int    `json:"id"`
	FirstName  string `json:"firstName" validate:"required,min=1,max=32"`
	LastName   string `json:"lastName" validate:"required,min=1,max=32"`
	VunetId    string `json:"vunetId" validate:"required"`
	Password   string `json:"password"`
	Balance    int    `json:"balance"`
	Visibility int    `json:"visibility" validate:"max=1"`
}

type MinimalUser struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Balance   int    `json:"balance"`
}

type PurchaseRequest struct {
	Password string          `json:"password" validate:"required,lte=8,gte=4"`
	Data     []PurchaseEntry `json:"data" validate:"required,dive"`
}

type PurchaseEntry struct {
	ProductId int `json:"productId" validate:"required"`
	Amount    int `json:"amount" validate:"required,lte=40000,gte=1"`
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

type TaxCategory struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Percentage int    `json:"percentage"`
}
