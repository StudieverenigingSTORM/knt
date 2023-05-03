package kntrouter

import (
	"encoding/json"
	"fmt"
	"knt/internal/kntdb"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/logger"
)

func getAdminProducts(w http.ResponseWriter, r *http.Request) {
	data, err := generateJsonResponse(kntdb.GetAllProducts())
	if err != nil {
		logger.Error("Unable to get products: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	fmt.Fprint(w, string(data))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	data, err := generateJsonResponse(kntdb.GetMinimalProducts())
	if err != nil {
		logger.Error("Unable to get minimal products: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	fmt.Fprint(w, string(data))
}

func getAdminProduct(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(chi.URLParam(r, "productId"))
	if err != nil {
		logger.Error("Unable to get product id: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	product, err := kntdb.GetProduct(productId)
	if err != nil {
		logger.Error("Unable to get product: ", err.Error(), " id: ", productId)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	jsonString, _ := json.Marshal(product)
	fmt.Fprint(w, string(jsonString))
}

func createNewProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body kntdb.Product
	err := decoder.Decode(&body)
	if err != nil {
		logger.Error("Unable to decode body: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	id, err := kntdb.CreateNewProduct(body)
	if err != nil {
		logger.Error("Unable to create new product: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	var idStruct struct {
		Id int64 `json:"id"`
	}
	idStruct.Id = id
	w.WriteHeader(http.StatusCreated)
	jsonString, _ := json.Marshal(idStruct)
	fmt.Fprint(w, string(jsonString))
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body kntdb.Product
	err := decoder.Decode(&body)
	if err != nil {
		logger.Error("Unable to decode body: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	_, err = kntdb.UpdateProduct(body)
	if err != nil {
		logger.Error("Unable to update product: ", err.Error(), " id: ", body.Id)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
