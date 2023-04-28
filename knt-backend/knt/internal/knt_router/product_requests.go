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
