package kntrouter

import (
	"fmt"
	"knt/internal/kntdb"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func getAdminProducts(w http.ResponseWriter, r *http.Request) {
	data, err := kntdb.GetAllProducts()
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}
	fmt.Fprint(w, generateJsonFromStruct(data))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	data, err := kntdb.GetMinimalProducts()
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}
	fmt.Fprint(w, generateJsonFromStruct(data))
}

func getAdminProduct(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(chi.URLParam(r, "productId"))
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	product, err := kntdb.GetProduct(productId)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	fmt.Fprint(w, generateJsonFromStruct(product))
}

func createNewProduct(w http.ResponseWriter, r *http.Request) {
	body, err := makeAndValidateStruct[kntdb.Product](r.Body)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	id, err := kntdb.CreateNewProduct(body)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	var idStruct IdResponse
	idStruct.Id = id
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, generateJsonFromStruct(idStruct))
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := makeAndValidateStruct[kntdb.Product](r.Body)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	_, err = kntdb.UpdateProduct(body)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func getTaxCategories(w http.ResponseWriter, r *http.Request) {
	data, err := kntdb.GetTaxCategories()
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}
	fmt.Fprint(w, generateJsonFromStruct(data))
}
