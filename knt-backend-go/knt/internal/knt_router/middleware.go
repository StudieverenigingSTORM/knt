package kntrouter

import (
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

func assignAdminMiddleware(r chi.Router) {
	r.Use(adminMiddleware)
}

func assignUserMiddleware(r chi.Router) {
	r.Use(userMiddleware)
}

// This sets the cors for all requests, this should be edited in the config
func setCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", viper.GetString("cors"))
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
func adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Trying to auth admin... and not succeeding")
		next.ServeHTTP(w, r)
	})
}

// Middleware to auth user
func userMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Trying to auth user... and not succeeding")
		next.ServeHTTP(w, r)
	})
}
