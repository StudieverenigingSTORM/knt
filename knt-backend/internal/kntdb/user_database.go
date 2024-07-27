package kntdb

import (
	"errors"
)

func GetAllUsers() ([]User, error) {
	return genericQuery[User]("select * from user")
}

func GetAllMinimalUsers() ([]MinimalUser, error) {
	return genericQuery[MinimalUser]("select vunetid, first_name, last_name, balance from user where visibility = 1")
}

func GetMinimalUser(userId string) (MinimalUser, error) {
	user, err := getFirstEntry[MinimalUser]("select vunetid, first_name, last_name, balance from user where vunetid = ? and visibility = 1", userId)
	if user.VunetId == "" && err == nil {
		err = errors.New("User not found")
	}
	return user, err
}

func GetUser(userID string) (User, error) {
	user, err := getFirstEntry[User]("select * from user where vunetid = ?", userID)
	if user.Id == "" && err == nil {
		err = errors.New("User not found")
	}
	return user, err
}

func CreateNewUser(user User) (int64, error) {
	return commitTransaction(
		"insert into user (vunetid, first_name, last_name, password, visibility) VALUES (?, ?, ?, ?, ?)",
		user.Id, user.FirstName, user.LastName, user.Password, user.Visibility)
}

func UpdateUser(user User) (int64, error) {
	if user.Id == "" {
		return 0, errors.New("invalid user")
	}

	oldUser, err := GetUser(user.Id)
	if err != nil {
		return 0, errors.New("invalid user")
	}

	if user.Password != "" {
		user.Password = ShaHashing(user.Password)
	} else {
		user.Password = oldUser.Password
	}
	return commitTransaction(
		"update user set first_name = ?, last_name = ?, password = ?, visibility = ? where vunetid = ?",
		user.FirstName, user.LastName, user.Password, user.Visibility, user.Id)
}

func GetPopulatedTransactions(perPage int, page int) ([]TransactionDetails, error) {
	return genericQuery[TransactionDetails](
		"select T.id, U.vunetid, T.starting_balance, T.delta_balance,"+
			"T.final_balance, T.ref, R.timestamp, R.data "+
			"from transactions T "+
			"left join user U "+
			"on U.vunetid = T.user_id "+
			"left join receipts R "+
			"on T.receipt_id = R.id "+
			"order by T.id asc limit ? offset ?",
		perPage, page*perPage)
}
