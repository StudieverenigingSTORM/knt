package kntrouter

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func AssignRoutes(r *mux.Router, db *sql.DB) {
	configRoutes := viper.Sub("routes")
	r.HandleFunc(configRoutes.GetString("ping"), ping)
	r.HandleFunc(configRoutes.GetString("getUserList"), getUsers(db))
}
