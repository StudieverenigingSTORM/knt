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
	user, err := getFirstEntry[MinimalUser]("select id, first_name, last_name, balance from user where id = ? and visibility = 1", userId)
	if user.Id == 0 && err == nil {
		err = errors.New("User not found")
	}
	return user, err
}

func GetUser(userID int) (User, error) {
	user, err := getFirstEntry[User]("select * from user where id = ?", userID)
	if user.Id == 0 && err == nil {
		err = errors.New("User not found")
	}
	return user, err
}

func GetUserByVunetId(VunetId string) (User, error) {
	user, err := getFirstEntry[User]("select * from user where vunetid = ?", VunetId)
	if user.Id == 0 && err == nil {
		err = errors.New("User not found")
	}
	return user, err
}

func CreateNewUser(user User) (int64, error) {
	return commitTransaction(
		"insert into user (first_name, last_name, vunetid, password, visibility) VALUES (?, ?, ?, ?, ?)",
		user.FirstName, user.LastName, user.VunetId, user.Password, user.Visibility)
}

func UpdateUser(user User) (int64, error) {
	if user.Id == 0 {
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
		"update user set first_name = ?, last_name = ?, vunetid = ?, password = ?, visibility = ? where id = ?",
		user.FirstName, user.LastName, user.VunetId, user.Password, user.Visibility, user.Id)
}

func GetPopulatedTransactions(perPage int, page int) ([]TransactionDetails, error) {
	return genericQuery[TransactionDetails](
		"select T.id, U.vunetid, T.starting_balance, T.delta_balance,"+
			"T.final_balance, T.ref, R.timestamp, R.data "+
			"from transactions T "+
			"left join user U "+
			"on U.id = T.user_id "+
			"left join receipts R "+
			"on T.receipt_id = R.id "+
			"order by T.id asc limit ? offset ?",
		perPage, page*perPage)
}
