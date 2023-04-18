package kntrouter

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"kntdatabase"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func getUsers(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return generateJsonResponse(kntdatabase.GetAllMinimalUsers(db))
}

func getUser(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.Atoi(chi.URLParam(r, "userId"))
		checkErr(w, err)
		user, err := kntdatabase.GetMinimalUser(db, userId)
		checkErr(w, err)
		jsonString, _ := json.Marshal(user)
		fmt.Fprintf(w, string(jsonString))
	}
}

func makePurchase(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.Atoi(chi.URLParam(r, "userId"))
		checkErr(w, err)

		decoder := json.NewDecoder(r.Body)
		var body kntdatabase.PurchaseRequest
		err = decoder.Decode(&body)
		checkErr(w, err)

		err = kntdatabase.MakeTransaction(userId, body, db)
		checkErr(w, err)

		// Write success
		w.WriteHeader(http.StatusCreated)
	}
}

func generateJsonResponse[K any](data K, err error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		checkErr(w, err)
		jsonString, _ := json.Marshal(data)
		fmt.Fprintf(w, string(jsonString))
	}
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented")
}

func checkErr(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
}
