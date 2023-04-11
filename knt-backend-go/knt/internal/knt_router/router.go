package knt_router

import (
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func AssignRoutes(r *mux.Router) {
	configRoutes := viper.Sub("routes")
	r.HandleFunc(configRoutes.GetString("ping"), ping)
}
