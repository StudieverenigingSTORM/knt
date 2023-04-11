package kntrouter

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"kntdatabase"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getProducts(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		products := kntdatabase.GetAllProducts(db)
		jsonString, _ := json.Marshal(products)
		fmt.Fprintf(w, string(jsonString))
	}
}

func getUsers(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users := kntdatabase.GetAllUsers(db)
		jsonString, _ := json.Marshal(users)
		fmt.Fprintf(w, string(jsonString))
	}
}
