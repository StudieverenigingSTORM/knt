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

func getUsersAdmin(w http.ResponseWriter, r *http.Request) {
	data, err := generateJsonResponse(kntdb.GetAllUsers())
	if err != nil {
		logger.Error("Unable to get users: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	fmt.Fprint(w, string(data))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	data, err := generateJsonResponse(kntdb.GetAllMinimalUsers())
	if err != nil {
		logger.Error("Unable to get minimal users: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	fmt.Fprint(w, string(data))
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body kntdb.User
	err := decoder.Decode(&body)
	if err != nil {
		logger.Error("Unable to decode body: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	body.Password = kntdb.ShaHashing(body.Password)
	id, err := kntdb.CreateNewUser(body)
	if err != nil {
		logger.Error("Unable to create new user: ", err.Error())
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

func getUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		logger.Error("Unable to get user id: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	user, err := kntdb.GetMinimalUser(userId)
	if err != nil {
		logger.Error("Unable to get user: ", err.Error(), " id: ", userId)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	jsonString, _ := json.Marshal(user)
	fmt.Fprint(w, string(jsonString))
}

func getAdminUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		logger.Error("Unable to get user id: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	user, err := kntdb.GetUser(userId)
	if err != nil {
		logger.Error("Unable to get user: ", err.Error(), " id: ", userId)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	jsonString, _ := json.Marshal(user)
	fmt.Fprint(w, string(jsonString))
}

func makePurchase(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "userId"))
	if err != nil {
		logger.Error("Unable to get user id: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var body kntdb.PurchaseRequest
	err = decoder.Decode(&body)
	if err != nil {
		logger.Error("Unable to decode body: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	var format struct {
		Cost int `json:"moneySpent"`
	}
	format.Cost, err = kntdb.MakeTransaction(userId, body)
	if err != nil {
		logger.Error("Unable to make transaction: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// Write success
	w.WriteHeader(http.StatusCreated)
	s, _ := json.Marshal(format)
	fmt.Fprint(w, string(s))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body kntdb.User
	err := decoder.Decode(&body)
	if err != nil {
		logger.Error("Unable to decode body: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	_, err = kntdb.UpdateUser(body)
	if err != nil {
		logger.Error("Unable to decode body: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func updateUserBalance(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var format struct {
		Balance   int    `json:"balance"`
		VunetId   string `json:"vunetid"`
		Reference string `json:"reference"`
	}
	err := decoder.Decode(&format)
	if err != nil {
		logger.Error("Unable to decode body: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	user, err := kntdb.GetUserByVunetId(format.VunetId)
	if err != nil {
		logger.Error("Unable to get user: ", err.Error(), " vunetid: ", format.VunetId)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	if user.Id == 0 {
		logger.Error("User does not exist vunetid: ", format.VunetId)
		http.Error(w, "Cannot find user", http.StatusUnprocessableEntity)
		return
	}
	data, _ := json.Marshal(format)
	err = kntdb.UpdateUserBalance(user, format.Balance, string(data), format.Reference)
	if err != nil {
		logger.Error("Unable to update user balance: ", err.Error(), " vunetid: ", format.VunetId)
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	pp := r.URL.Query().Get("pp")
	p := r.URL.Query().Get("p")

	if pp == "" {
		pp = "10"
	}
	if p == "" {
		p = "0"
	}

	itemCount, err := strconv.Atoi(pp)
	if err != nil {
		logger.Error("Unable to get item count: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	pageNum, err := strconv.Atoi(p)
	if err != nil {
		logger.Error("Unable to get page: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	data, err := kntdb.GetPopulatedTransactions(itemCount, pageNum)
	if err != nil {
		logger.Error("Unable to return transactions: ", err.Error())
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	s, _ := json.Marshal(data)
	fmt.Fprint(w, string(s))
}
