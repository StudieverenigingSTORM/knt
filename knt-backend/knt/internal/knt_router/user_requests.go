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

func getUsersAdmin(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return generateJsonResponse(kntdatabase.GetAllUsers(db))

}

func getUsers(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return generateJsonResponse(kntdatabase.GetAllMinimalUsers(db))
}

func createNewUser(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body kntdatabase.User
		err := decoder.Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		body.Password = kntdatabase.ShaHashing(body.Password)
		id, err := kntdatabase.CreateNewUser(db, body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		var idStruct struct {
			Id int64 `json:"id"`
		}
		idStruct.Id = id
		jsonString, _ := json.Marshal(idStruct)
		fmt.Fprintf(w, string(jsonString))
	}
}

func getUser(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.Atoi(chi.URLParam(r, "userId"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		user, err := kntdatabase.GetMinimalUser(db, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		jsonString, _ := json.Marshal(user)
		fmt.Fprintf(w, string(jsonString))
	}
}

func getAdminUser(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.Atoi(chi.URLParam(r, "userId"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		user, err := kntdatabase.GetUser(db, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		jsonString, _ := json.Marshal(user)
		fmt.Fprintf(w, string(jsonString))
	}
}

func makePurchase(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.Atoi(chi.URLParam(r, "userId"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var body kntdatabase.PurchaseRequest
		err = decoder.Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		var format struct {
			Cost int `json:"moneySpent"`
		}
		format.Cost, err = kntdatabase.MakeTransaction(userId, body, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		// Write success
		w.WriteHeader(http.StatusCreated)
		s, _ := json.Marshal(format)
		fmt.Fprintf(w, string(s))
	}
}

func updateUser(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body kntdatabase.User
		err := decoder.Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		_, err = kntdatabase.UpdateUser(db, body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func updateUserBalance(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var format struct {
			Balance   int    `json:"balance"`
			VunetId   string `json:"vunetid"`
			Reference string `json:"reference"`
		}
		err := decoder.Decode(&format)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		user, err := kntdatabase.GetUserByVunetId(db, format.VunetId)
		if err != nil || user.Id == 0 {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		data, _ := json.Marshal(format)
		err = kntdatabase.UpdateUserBalance(user, format.Balance, db, string(data), format.Reference)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
