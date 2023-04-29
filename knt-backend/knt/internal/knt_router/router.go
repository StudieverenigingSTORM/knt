package kntrouter

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

func AssignRoutes(r chi.Router, db *sql.DB) {

	configRoutes := viper.Sub("routes")
	assignGeneralMiddlewares(r)

	r.MethodFunc(http.MethodGet, configRoutes.GetString("ping"), ping)

	assignUserRoutes(r, db, configRoutes)
	assignAdminRoutes(r, db, configRoutes)
}

// Assigns user routes with the approriate user middleware
func assignUserRoutes(r chi.Router, db *sql.DB, configRoutes *viper.Viper) {
	r.Route(configRoutes.GetString("basicEndpoint"), func(r chi.Router) {
		assignUserMiddleware(r, db)
		r.MethodFunc(http.MethodGet, configRoutes.GetString("getUsers"), getUsers(db))
		r.MethodFunc(http.MethodGet, configRoutes.GetString("getUser"), getUser(db))
		r.MethodFunc(http.MethodPost, configRoutes.GetString("makePurchase"), makePurchase(db))

		r.MethodFunc(http.MethodGet, configRoutes.GetString("getProducts"), getProducts(db))
	})
}

// Assigns admin routes with the approriate admin middleware
func assignAdminRoutes(r chi.Router, db *sql.DB, configRoutes *viper.Viper) {
	r.Route(configRoutes.GetString("adminEndpoint"), func(r chi.Router) {
		assignAdminMiddleware(r, db)
		r.MethodFunc(http.MethodGet, configRoutes.GetString("getUsersAdmin"), getUsersAdmin(db))
		r.MethodFunc(http.MethodPost, configRoutes.GetString("createNewUser"), createNewUser(db))
		r.MethodFunc(http.MethodPut, configRoutes.GetString("updateUser"), updateUser(db))
		r.MethodFunc(http.MethodPost, configRoutes.GetString("updateUserMoney"), updateUserBalance(db))
		r.MethodFunc(http.MethodGet, configRoutes.GetString("getUserAdmin"), getAdminUser(db))

		r.MethodFunc(http.MethodPost, configRoutes.GetString("createNewProduct"), createNewProduct(db))
		r.MethodFunc(http.MethodPut, configRoutes.GetString("updateProduct"), updateProduct(db))
		r.MethodFunc(http.MethodGet, configRoutes.GetString("getFullProducts"), getAdminProducts(db))
		r.MethodFunc(http.MethodGet, configRoutes.GetString("getFullProduct"), getAdminProduct(db))

		r.MethodFunc(http.MethodGet, configRoutes.GetString("taxcategories"), notImplemented)
		r.MethodFunc(http.MethodPost, configRoutes.GetString("taxcategories"), notImplemented)
		r.MethodFunc(http.MethodPut, configRoutes.GetString("taxcategories"), notImplemented)

		r.MethodFunc(http.MethodGet, configRoutes.GetString("transactions"), notImplemented)
	})
}
