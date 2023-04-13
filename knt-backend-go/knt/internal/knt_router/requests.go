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
	return generateJsonResponse(kntdatabase.GetAllProducts(db))
}

func getUsersAdmin(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return generateJsonResponse(kntdatabase.GetAllUsers(db))

}

func generateJsonResponse[K any](data K) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonString, _ := json.Marshal(data)
		fmt.Fprintf(w, string(jsonString))
	}
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented")
}
