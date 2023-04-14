package kntrouter

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

func AssignRoutes(r chi.Router, db *sql.DB) {

	configRoutes := viper.Sub("routes")
	assignGeneralMiddlewares(r)

	r.HandleFunc(configRoutes.GetString("ping"), ping)

	assignUserRoutes(r, db, configRoutes)
	assignAdminRoutes(r, db, configRoutes)
}

// Assigns user routes with the approriate user middleware
func assignUserRoutes(r chi.Router, db *sql.DB, configRoutes *viper.Viper) {
	r.Route(configRoutes.GetString("basicEndpoint"), func(r chi.Router) {
		assignUserMiddleware(r)
		r.HandleFunc(configRoutes.GetString("getUsers"), getUsers(db))
		r.HandleFunc(configRoutes.GetString("getUser"), notImplemented)
		r.HandleFunc(configRoutes.GetString("makePurchase"), notImplemented)

		r.HandleFunc(configRoutes.GetString("getProducts"), getProducts(db))
	})
}

// Assigns admin routes with the approriate admin middleware
func assignAdminRoutes(r chi.Router, db *sql.DB, configRoutes *viper.Viper) {
	r.Route(configRoutes.GetString("adminEndpoint"), func(r chi.Router) {
		assignAdminMiddleware(r, db)
		r.HandleFunc(configRoutes.GetString("getUsersAdmin"), getUsersAdmin(db))
		r.HandleFunc(configRoutes.GetString("createNewUser"), notImplemented)
		r.HandleFunc(configRoutes.GetString("deleteUser"), notImplemented)
		r.HandleFunc(configRoutes.GetString("updateUser"), notImplemented)
		r.HandleFunc(configRoutes.GetString("updateUserMoney"), notImplemented)
		r.HandleFunc(configRoutes.GetString("getUserAdmin"), notImplemented)

		r.HandleFunc(configRoutes.GetString("createNewProduct"), notImplemented)
		r.HandleFunc(configRoutes.GetString("deleteProduct"), notImplemented)
		r.HandleFunc(configRoutes.GetString("updateProduct"), notImplemented)

	})
}
