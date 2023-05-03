package kntrouter

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

func AssignRoutes(r chi.Router) {

	configRoutes := viper.Sub("routes")
	assignGeneralMiddlewares(r)

	r.MethodFunc(http.MethodGet, configRoutes.GetString("ping"), ping)

	assignUserRoutes(r, configRoutes)
	assignAdminRoutes(r, configRoutes)
}

// Assigns user routes with the approriate user middleware
func assignUserRoutes(r chi.Router, configRoutes *viper.Viper) {
	r.Route(configRoutes.GetString("basicEndpoint"), func(r chi.Router) {
		assignUserMiddleware(r)
		r.MethodFunc(http.MethodGet, configRoutes.GetString("getUsers"), getUsers)
		r.MethodFunc(http.MethodGet, configRoutes.GetString("getUser"), getUser)
		r.MethodFunc(http.MethodPost, configRoutes.GetString("makePurchase"), makePurchase)

		r.MethodFunc(http.MethodGet, configRoutes.GetString("getProducts"), getProducts)
	})
}

// Assigns admin routes with the approriate admin middleware
func assignAdminRoutes(r chi.Router, configRoutes *viper.Viper) {
	r.Route(configRoutes.GetString("adminEndpoint"), func(r chi.Router) {
		assignAdminMiddleware(r)
		r.MethodFunc(http.MethodGet, configRoutes.GetString("getUsersAdmin"), getUsersAdmin)
		r.MethodFunc(http.MethodGet, configRoutes.GetString("getUserAdmin"), getAdminUser)

		r.MethodFunc(http.MethodGet, configRoutes.GetString("getFullProducts"), getAdminProducts)
		r.MethodFunc(http.MethodGet, configRoutes.GetString("getFullProduct"), getAdminProduct)

		r.MethodFunc(http.MethodGet, configRoutes.GetString("taxcategories"), notImplemented)

		r.MethodFunc(http.MethodGet, configRoutes.GetString("transactions"), getTransactions)

		r.Group(func(r chi.Router) {
			r.Use(logAdminMiddleware)

			r.MethodFunc(http.MethodPost, configRoutes.GetString("taxcategories"), notImplemented)
			r.MethodFunc(http.MethodPut, configRoutes.GetString("taxcategories"), notImplemented)

			r.MethodFunc(http.MethodPost, configRoutes.GetString("createNewProduct"), createNewProduct)
			r.MethodFunc(http.MethodPut, configRoutes.GetString("updateProduct"), updateProduct)

			r.MethodFunc(http.MethodPost, configRoutes.GetString("createNewUser"), createNewUser)
			r.MethodFunc(http.MethodPut, configRoutes.GetString("updateUser"), updateUser)
			r.MethodFunc(http.MethodPost, configRoutes.GetString("updateUserMoney"), updateUserBalance)
		})
	})
}
