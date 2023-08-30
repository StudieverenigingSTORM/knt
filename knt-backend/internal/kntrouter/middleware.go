package kntrouter

import (
	"bytes"
	"errors"
	"io"
	"knt/internal/kntdb"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/logger"
	"github.com/spf13/viper"
)

// Middleware assigning calls
// These basically get called from the router on the appropriate routes
// Be careful when assigning the middlewares
// Their order matters, the middlewares that get declared first also get executed first
func assignGeneralMiddlewares(r chi.Router) {
	r.Use(setCors)
	r.Use(preflightMiddleware)
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
		w.Header().Set("Access-Control-Allow-Origin", viper.GetString("corsAllowOrigin"))
		w.Header().Set("Access-Control-Allow-Credentials", viper.GetString("corsAllowCredentials"))
		w.Header().Set("Access-Control-Allow-Methods", viper.GetString("corsAllowMethods"))
		w.Header().Set("Access-Control-Allow-Headers", viper.GetString("corsAllowHeaders"))

		next.ServeHTTP(w, r)
	})
}

// Middleware to log all http calls
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Access: ", r.URL.Path, " ", r.Method)
		next.ServeHTTP(w, r)
	})
}

// Middleware to allow preflight in browser
func preflightMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Middleware to auth admin
// We need to do this to provide database access to the middleware
func adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Get the header and validate it
		key := r.Header.Get("X-API-Key")
		if key == "" {
			checkAndSendError(w, errors.New("API key missing"), http.StatusUnauthorized)
			return
		}
		privileges := kntdb.CheckUserPrivileges(key)
		//Allow admins
		if privileges == "admin" {
			next.ServeHTTP(w, r)
			return
		}

		checkAndSendError(w, errors.New("Unauthorized"), http.StatusProxyAuthRequired)
	})
}

func logAdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Get the header and validate it
		key := r.Header.Get("X-Admin-Id")
		if key == "" {
			checkAndSendError(w, errors.New("no admin key provided"), http.StatusUnauthorized)
			return
		}

		data, _ := io.ReadAll(r.Body)
		//after reading the data we want to put it back in the buffer for other middlewares/requests to read it
		r.Body.Close()
		r.Body = io.NopCloser(bytes.NewBuffer(data))

		err := kntdb.AddAdminLogs(r.URL.Path, r.Method, string(data), key)
		if err != nil {
			checkAndSendError(w, err, http.StatusUnauthorized)
			return
		}
		//Write appropriate headers
		next.ServeHTTP(w, r)
	})
}

// Middleware to auth user
func userMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Get the header and validate it
		key := r.Header.Get("X-API-Key")
		if key == "" {
			checkAndSendError(w, errors.New("API key missing"), http.StatusUnauthorized)
			return
		}
		privileges := kntdb.CheckUserPrivileges(key)
		//Allow admins
		if privileges == "admin" || privileges == "user" {
			next.ServeHTTP(w, r)
			return
		}
		//Write appropriate headers
		checkAndSendError(w, errors.New("Unauthorized"), http.StatusProxyAuthRequired)
	})
}
