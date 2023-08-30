package kntrouter

import (
	"fmt"
	"knt/internal/kntdb"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func getUsersAdmin(w http.ResponseWriter, r *http.Request) {
	data, err := kntdb.GetAllUsers()
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}
	fmt.Fprint(w, generateJsonFromStruct(data))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	data, err := kntdb.GetAllMinimalUsers()
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}
	fmt.Fprint(w, generateJsonFromStruct(data))
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	body, err := makeAndValidateStruct[kntdb.User](r.Body)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	if body.Password != "" {
		body.Password = kntdb.ShaHashing(body.Password)
	}

	id, err := kntdb.CreateNewUser(body)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	var idStruct IdResponse
	idStruct.Id = id
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, generateJsonFromStruct(idStruct))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	user, err := kntdb.GetMinimalUser(userId)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	fmt.Fprint(w, generateJsonFromStruct(user))
}

func getAdminUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	user, err := kntdb.GetUser(userId)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	fmt.Fprint(w, generateJsonFromStruct(user))
}

func makePurchase(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	body, err := makeAndValidateStruct[kntdb.PurchaseRequest](r.Body)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	var format SpentFormat
	format.Cost, err = kntdb.MakeTransaction(userId, body)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	// Write success
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, generateJsonFromStruct(format))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	body, err := makeAndValidateStruct[kntdb.User](r.Body)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	_, err = kntdb.UpdateUser(body)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func updateUserBalance(w http.ResponseWriter, r *http.Request) {
	format, err := makeAndValidateStruct[WebhookFormat](r.Body)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	user, err := kntdb.GetUserByVunetId(format.VunetId)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err = kntdb.UpdateUserBalance(user, format.Balance, generateJsonFromStruct(format), format.Reference)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	perPage := r.URL.Query().Get("perPage")
	page := r.URL.Query().Get("page")

	if perPage == "" {
		perPage = "10"
	}
	if page == "" {
		page = "0"
	}

	itemCount, err := strconv.Atoi(perPage)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	data, err := kntdb.GetPopulatedTransactions(itemCount, pageNum)
	if err != nil {
		checkAndSendError(w, err, http.StatusUnprocessableEntity)
		return
	}

	fmt.Fprint(w, generateJsonFromStruct(data))
}
