package kntdb

import (
	"errors"
)

func GetAllUsers() ([]User, error) {
	return genericQuery[User]("select * from user")
}

func GetAllMinimalUsers() ([]MinimalUser, error) {
	return genericQuery[MinimalUser]("select id, first_name, last_name, balance from user where visibility = 1")
}

func GetMinimalUser(userId int) (MinimalUser, error) {
	return getFirstEntry[MinimalUser]("select id, first_name, last_name, balance from user where id = ? and visibility = 1", userId)
}

func GetUser(userID int) (User, error) {
	user, err := getFirstEntry[User]("select * from user where id = ?", userID)
	if user.Id == 0 {
		return user, errors.New("User not found")
	}
	return user, err
}

func GetUserByVunetId(VunetId string) (User, error) {
	return getFirstEntry[User]("select * from user where vunetid = ?", VunetId)
}

func CreateNewUser(user User) (int64, error) {
	return commitTransaction(
		"insert into user (first_name, last_name, vunetid, password, visibility) VALUES (?, ?, ?, ?, ?)",
		user.FirstName, user.LastName, user.VunetId, user.Password, user.Visibility)
}

func UpdateUser(userNew User) (int64, error) {
	if userNew.Id == 0 {
		return 0, errors.New("invalid user")
	}

	userOld, err := GetUser(userNew.Id)
	if err != nil {
		return 0, err
	}
	finalUser := ModifyUser(userNew, userOld)

	return commitTransaction(
		"update user set first_name = ?, last_name = ?, vunetid = ?, password = ?, visibility = ? where id = ?",
		finalUser.FirstName, finalUser.LastName, finalUser.VunetId, finalUser.Password, finalUser.Visibility, finalUser.Id)
}

// Returns a user object made from the old user and changes in the request
func ModifyUser(new User, old User) User {
	if new.FirstName != "" {
		old.FirstName = new.FirstName
	}
	if new.LastName != "" {
		old.LastName = new.LastName
	}
	if new.VunetId != "" {
		old.VunetId = new.VunetId
	}
	old.Visibility = new.Visibility
	if new.Password != "" {
		old.Password = ShaHashing(new.Password)
	}
	return old
}

func GetPopulatedTransactions(pp int, p int) ([]TransactionDetails, error) {
	basicTrans, err := getBasicTransactions(pp, p)
	if err != nil {
		return nil, err
	}

	var result []TransactionDetails
	for _, t := range basicTrans {
		u, err := GetUser(t.UserId)
		if err != nil {
			return nil, err
		}
		r, err := getReceipt(t.ReceiptId)
		if err != nil {
			return nil, err
		}

		result = append(result, TransactionDetails{
			Id:              t.Id,
			VunetId:         u.VunetId,
			StartingBalance: t.StartingBalance,
			DeltaBalance:    t.DeltaBalance,
			FinalBalance:    t.FinalBalance,
			Reference:       t.Reference,
			Timestamp:       r.Timestamp,
			Data:            r.Data,
		})
	}
	return result, nil
}
