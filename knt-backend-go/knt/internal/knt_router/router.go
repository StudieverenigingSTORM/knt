package kntrouter

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

func AssignRoutes(r chi.Router, db *sql.DB) {
	configRoutes := viper.Sub("routes")
	r.HandleFunc(configRoutes.GetString("ping"), ping)
	r.HandleFunc(configRoutes.GetString("getProducts"), getProducts(db))
	r.HandleFunc(configRoutes.GetString("getUsers"), getUsers(db))
}
