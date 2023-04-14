package kntrouter

import (
	"database/sql"
	"kntdatabase"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

// Middleware assigning calls
// These basically get called from the router on the appropriate routes
func assignGeneralMiddlewares(r chi.Router) {
	r.Use(setCors)
	r.Use(loggingMiddleware)
}

func assignAdminMiddleware(r chi.Router, db *sql.DB) {
	r.Use(generateAdminMiddleware(db))
}

func assignUserMiddleware(r chi.Router) {
	r.Use(userMiddleware)
}

// This sets the cors for all requests, this should be edited in the config
func setCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", viper.GetString("cors"))
		w.Header().Set("Access-Control-Allow-Credentials", viper.GetString("cors"))
		w.Header().Set("Access-Control-Allow-Methods", viper.GetString("cors"))
		w.Header().Set("Access-Control-Allow-Headers", viper.GetString("cors"))
		next.ServeHTTP(w, r)
	})
}

// Middleware to log all http calls
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// Middleware to auth admin
// We need to do this to provide database access to the middleware
func generateAdminMiddleware(db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//Allow preflight
			if r.Method == "OPTIONS" {
				w.WriteHeader(200)
				return
			}
			//Get the header and validate it
			key := r.Header.Get("X-API-Key")
			if key == "" {
				http.Error(w, http.StatusText(401), 401)
				return
			}
			privileges := kntdatabase.CheckUserPrivileges(key, db)
			//Allow admins
			if privileges == "admin" {
				next.ServeHTTP(w, r)
				return
			}
			//Write appropriate headers

			http.Error(w, http.StatusText(407), 407)

		})
	}
}

// Middleware to auth user
func userMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Trying to auth user... and not succeeding")
		next.ServeHTTP(w, r)
	})
}
