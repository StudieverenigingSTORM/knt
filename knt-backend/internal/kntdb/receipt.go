package kntdb

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

//This file handles everything having to do with purchases and receipts

func MakeTransaction(userId int, purchase PurchaseRequest) (int, error) {
	//Begins the transaction
	//This is important because if ANY error were to occur we need to reset the database to its original state
	transaction, err := DB.Begin()
	if err != nil {
		return 0, err
	}
	// get user
	user, err := GetUser(userId)
	if err != nil {
		return 0, err
	}

	if user.Password == "" {
		return 0, errors.New("User has no password, cannot complete the transaction")
	}

	//pin validation
	if !ValidatePin(purchase.Password, user) {
		return 0, errors.New("incorrect pin")
	}
	//calculate cost
	cost, err := calculateCost(purchase.Data)
	if err != nil {
		return 0, errors.New("could not find all products")
	}
	//validate users balance
	if cost > user.Balance {
		return 0, errors.New("insufficient balance on user")
	}
	//generate receipt
	receiptId, err := generateReceipt(transaction, purchase.Data)
	if err != nil {
		return 0, err
	}
	//make a transaction
	err = generateTransaction(transaction, userId, user.Balance, cost, user.Balance-cost, receiptId, "")
	if err != nil {
		return 0, err
	}
	//subtract balance
	err = setBalance(transaction, userId, user.Balance-cost)
	if err != nil {
		return 0, err
	}

	//Add tax
	err = addTaxTotals(purchase.Data, transaction, cost*-1)
	if err != nil {
		transaction.Rollback()
		return 0, err
	}

	//Finilize the transaction
	err = transaction.Commit()
	if err != nil {
		return 0, err
	}
	return cost * -1, nil
}

// Calculates the total cost of the purchased products
func calculateCost(entries []PurchaseEntry) (int, error) {
	var cost int

	for _, element := range entries {
		value, err := getProductValue(element)
		if err != nil {
			return 0, err
		}
		cost += value
	}

	// cost *= -1

	return cost, nil
}

// returns the value of a specific entry
func getProductValue(entry PurchaseEntry) (int, error) {
	value, err := getFirstSingleValue[int]("select price from product where id = ? and visibility = 1", entry.ProductId)
	return value * entry.Amount, err
}

// Generated the receipt and stores it in the database
func generateReceipt(transaction *sql.Tx, entries []PurchaseEntry) (int64, error) {
	dataString, err := json.Marshal(entries)
	if err != nil {
		return 0, err
	}

	return addReceipt(transaction, string(dataString))
}

func addReceipt(transaction *sql.Tx, dataString string) (int64, error) {
	receiptId, err := addToTransaction(transaction, "INSERT INTO receipts (data, timestamp) VALUES (?, datetime())", dataString)
	if err != nil {
		return 0, err
	}

	return receiptId, nil
}

// Generates the transaction and stores it in the database
func generateTransaction(transaction *sql.Tx, userId int, startingBal int, deltaBal int, finalBal int, receiptId int64, ref string) error {
	_, err := addToTransaction(transaction,
		"INSERT INTO transactions (user_id, starting_balance, delta_balance, final_balance, receipt_id, ref) VALUES (?, ?, ?, ?, ?, ?)",
		userId, startingBal, deltaBal, finalBal, receiptId, ref)
	return err
}

// Sets the users balance to a specified ammount
func setBalance(transaction *sql.Tx, userId int, balance int) error {
	_, err := addToTransaction(transaction, "UPDATE user SET balance = ? WHERE id = ?",
		balance, userId)
	return err
}

// Function designed to calculate total ammounts for tax reasons
// The taxes are stored in their own tables in the form of tax{year}
func addTaxTotals(entries []PurchaseEntry, transaction *sql.Tx, cost int) error {
	year := time.Now().Year()

	//Go through all the entries and apply the operation on all of them
	for _, entry := range entries {
		//check if entry exists
		tax, err := getFirstEntry[TaxEntry]("SELECT * FROM tax where year = ? and product_id = ?",
			year, entry.ProductId)
		if err != nil {
			return err
		}
		if tax.Id == 0 {
			//The row for this product doesnt already exist in the table so create a new row for it
			_, err = addToTransaction(transaction, "INSERT INTO tax (product_id, amount, totalCost, year) VALUES (?, ?, ?, ?)",
				entry.ProductId, entry.Amount, cost, year)
			if err != nil {
				return err
			}
			continue
		}
		//The product exists update it
		_, err = addToTransaction(transaction, "UPDATE tax SET amount = ?, totalCost = ? WHERE year = ? and id = ?",
			tax.Amount+entry.Amount, tax.TotalCost+cost, year, tax.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func UpdateUserBalance(user User, balance int, body string, ref string) error {
	transaction, err := DB.Begin()
	if err != nil {
		return err
	}

	receiptId, err := addReceipt(transaction, body)
	if err != nil {
		return err
	}

	err = generateTransaction(transaction, user.Id, user.Balance, balance, user.Balance+balance, receiptId, ref)
	if err != nil {
		return err
	}

	err = setBalance(transaction, user.Id, user.Balance+balance)
	if err != nil {
		return err
	}

	err = transaction.Commit()
	if err != nil {
		return err
	}

	return nil
}
