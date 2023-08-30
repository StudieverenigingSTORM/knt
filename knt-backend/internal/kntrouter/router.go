package kntrouter

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AssignRoutes(r chi.Router) {
	assignGeneralMiddlewares(r)

	r.MethodFunc(http.MethodGet, "/ping", ping)

	assignUserRoutes(r)
	assignAdminRoutes(r)
}

// Assigns user routes with the approriate user middleware
func assignUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		assignUserMiddleware(r)
		r.MethodFunc(http.MethodGet, "/", getUsers)
		r.MethodFunc(http.MethodGet, "/{userId}", getUser)
		r.MethodFunc(http.MethodPost, "/{userId}/purchase", makePurchase)

		r.MethodFunc(http.MethodGet, "/products", getProducts)
	})
}

// Assigns admin routes with the approriate admin middleware
func assignAdminRoutes(r chi.Router) {
	r.Route("/admin", func(r chi.Router) {
		assignAdminMiddleware(r)
		r.MethodFunc(http.MethodGet, "/users", getUsersAdmin)
		r.MethodFunc(http.MethodGet, "/users/{userId}", getAdminUser)

		r.MethodFunc(http.MethodGet, "/products", getAdminProducts)
		r.MethodFunc(http.MethodGet, "/products/{productId}", getAdminProduct)

		r.MethodFunc(http.MethodGet, "/tax", getTaxCategories)

		r.MethodFunc(http.MethodGet, "/transactions", getTransactions)

		r.Group(func(r chi.Router) {
			r.Use(logAdminMiddleware)

			r.MethodFunc(http.MethodPost, "/tax", notImplemented)
			r.MethodFunc(http.MethodPut, "/tax", notImplemented)

			r.MethodFunc(http.MethodPost, "/products", createNewProduct)
			r.MethodFunc(http.MethodPut, "/products", updateProduct)

			r.MethodFunc(http.MethodPost, "/users", createNewUser)
			r.MethodFunc(http.MethodPut, "/users", updateUser)
			r.MethodFunc(http.MethodPost, "/users/balance", updateUserBalance)
		})
	})
}
