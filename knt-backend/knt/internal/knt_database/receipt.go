package kntdatabase

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

//This file handles everything having to do with purchases and receipts

func MakeTransaction(userId int, purchase PurchaseRequest, db *sql.DB) error {
	// get user
	user, err := GetUser(db, userId)
	if err != nil {
		return err
	}
	//pin validation
	if !ValidatePin(purchase.Password, user, db) {
		return errors.New("Unauthorized")
	}
	//calculate cost
	cost, err := calculateCost(purchase.Data, db)
	if err != nil {
		return err
	}
	//validate users balance
	if cost > user.Balance {
		return errors.New("Insufficient balance on user")
	}
	//generate receipt
	receiptId, err := generateReceipt(db, purchase.Data)
	if err != nil {
		return err
	}
	//make a transaction
	err = generateTransaction(db, userId, user.Balance, cost, user.Balance-cost, receiptId)
	if err != nil {
		return err
	}
	//subtract balance
	err = setBalance(db, userId, user.Balance-cost)
	if err != nil {
		return err
	}

	//Add tax
	err = addTaxTotals(purchase.Data, db, cost)
	if err != nil {
		return err
	}

	return nil
}

// Calculates the total cost of the purchased products
func calculateCost(entries []PurchaseEntry, db *sql.DB) (int, error) {
	var cost int

	for _, element := range entries {
		value, err := getProductValue(element, db)
		if err != nil {
			return 0, err
		}
		cost += value
	}

	return cost, nil
}

// returns the value of a specific entry
func getProductValue(entry PurchaseEntry, db *sql.DB) (int, error) {
	value, err := getSingleValue[int](queryBuilder(db, "select price from product where id = ?", entry.ProductId))
	return value * entry.Amount, err
}

// Generated the receipt and stores it in the database
func generateReceipt(db *sql.DB, entries []PurchaseEntry) (int64, error) {
	dataString, err := json.Marshal(entries)
	if err != nil {
		return 0, err
	}

	receiptId, err := commitTransaction(db, "INSERT INTO receipts (data, timestamp) VALUES (?, datetime())", dataString)
	if err != nil {
		return 0, err
	}

	return receiptId, nil
}

// Generates the transaction and stores it in the database
func generateTransaction(db *sql.DB, userId int, startingBal int, deltaBal int, finalBal int, receiptId int64) error {
	_, err := commitTransaction(db, "INSERT INTO transactions (user_id, starting_balance, delta_balance, final_balance, receipt_id) VALUES (?, ?, ?, ?, ?)",
		userId, startingBal, deltaBal, finalBal, receiptId)
	return err
}

// Sets the users balance to a specified ammount
func setBalance(db *sql.DB, userId int, balance int) error {
	_, err := commitTransaction(db, "UPDATE user SET balance = ? WHERE id = ?", balance, userId)
	return err
}

func addTaxTotals(entries []PurchaseEntry, db *sql.DB, cost int) error {
	year := strconv.Itoa(time.Now().Year())

	//Ensure yearly tables existance
	err := ensureTaxTableExists(db, year)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		//check if entry exists
		tax, err := getSingleEntry[TaxEntry](queryBuilder(db, "SELECT * FROM tax"+year+" where product_id = ?", entry.ProductId))
		if err != nil {
			return err
		}
		if tax.Id == 0 {
			commitTransaction(db, "INSERT INTO tax"+year+" (product_id, amount, totalCost) VALUES (?, ?, ?)", entry.ProductId, entry.Amount, cost)
		} else {
			_, err := commitTransaction(db, "UPDATE tax"+year+" SET amount = ?, totalCost = ? WHERE id = ?", tax.Amount+entry.Amount, tax.TotalCost+cost, tax.Id)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func ensureTaxTableExists(db *sql.DB, year string) error {
	_, err := commitTransaction(db, "CREATE TABLE IF NOT EXISTS tax"+year+" (id INTEGER PRIMARY KEY AUTOINCREMENT, product_id INT, amount INT, totalCost INT)")
	return err
}
