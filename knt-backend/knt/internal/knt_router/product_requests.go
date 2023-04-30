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

func getAdminProducts(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return generateJsonResponse(kntdatabase.GetAllProducts(db))
}

func getProducts(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return generateJsonResponse(kntdatabase.GetMinimalProducts(db))
}

func getAdminProduct(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		productId, err := strconv.Atoi(chi.URLParam(r, "productId"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		product, err := kntdatabase.GetProduct(db, productId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		jsonString, _ := json.Marshal(product)
		fmt.Fprintf(w, string(jsonString))
	}
}

func createNewProduct(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body kntdatabase.Product
		err := decoder.Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		id, err := kntdatabase.CreateNewProduct(db, body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		var idStruct struct {
			Id int64 `json:"id"`
		}
		idStruct.Id = id
		w.WriteHeader(http.StatusCreated)
		jsonString, _ := json.Marshal(idStruct)
		fmt.Fprintf(w, string(jsonString))
	}
}

func updateProduct(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body kntdatabase.Product
		err := decoder.Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		_, err = kntdatabase.UpdateProduct(db, body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
