package kntdatabase

import (
	"database/sql"
	"errors"
)

func GetAllUsers(db *sql.DB) ([]User, error) {
	return genericQuery[User](queryBuilder(db, "select * from user"))
}

func GetAllMinimalUsers(db *sql.DB) ([]MinimalUser, error) {
	return genericQuery[MinimalUser](queryBuilder(db, "select id, first_name, last_name, balance from user where visibility = 1"))
}

func GetMinimalUser(db *sql.DB, userId int) (MinimalUser, error) {
	return getFirstEntry[MinimalUser](queryBuilder(db, "select id, first_name, last_name, balance from user where id = ? and visibility = 1", userId))
}

func GetUser(db *sql.DB, userID int) (User, error) {
	return getFirstEntry[User](queryBuilder(db, "select * from user where id = ?", userID))
}

func GetUserByVunetId(db *sql.DB, VunetId string) (User, error) {
	return getFirstEntry[User](queryBuilder(db, "select * from user where vunetid = ?", VunetId))
}

func CreateNewUser(db *sql.DB, user User) (int64, error) {
	return commitTransaction(db,
		"insert into user (first_name, last_name, vunetid, password, balance, visibility) VALUES (?, ?, ?, ?, ?, ?)",
		user.FirstName, user.LastName, user.VunetId, user.Password, user.Balance, user.Visibility)
}

func UpdateUser(db *sql.DB, userNew User) (int64, error) {
	if userNew.Id == 0 {
		return 0, errors.New("Invalid user")
	}

	userOld, err := GetUser(db, userNew.Id)
	if err != nil {
		return 0, err
	}
	finalUser := ModifyUser(userNew, userOld)

	return commitTransaction(db,
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
