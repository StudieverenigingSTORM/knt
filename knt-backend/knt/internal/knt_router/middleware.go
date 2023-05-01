package kntrouter

import (
	"bytes"
	"database/sql"
	"io"
	"kntdatabase"
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

func assignAdminMiddleware(r chi.Router, db *sql.DB) {
	r.Use(generateAdminMiddleware(db))

}

func assignUserMiddleware(r chi.Router, db *sql.DB) {
	r.Use(generateUserMiddleware(db))
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
func generateAdminMiddleware(db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//Get the header and validate it
			key := r.Header.Get("X-API-Key")
			if key == "" {
				logger.Error("API key missing")
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			privileges := kntdatabase.CheckUserPrivileges(key, db)
			//Allow admins
			if privileges == "admin" {
				next.ServeHTTP(w, r)
				return
			}

			//Write appropriate headers
			logger.Error("Unauthorized")
			http.Error(w, http.StatusText(http.StatusProxyAuthRequired), http.StatusProxyAuthRequired)
		})
	}
}

func logAdminMiddleware(db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//Get the header and validate it
			key := r.Header.Get("X-Admin-Id")
			if key == "" {
				logger.Error("No admin key provided")
				http.Error(w, "No admin key provided", http.StatusUnauthorized)
				return
			}

			data, _ := io.ReadAll(r.Body)
			//after reading the data we want to put it back in the buffer for other middlewares/requests to read it
			r.Body.Close()
			r.Body = io.NopCloser(bytes.NewBuffer(data))

			err := kntdatabase.AddAdminLogs(db, r.URL.Path, r.Method, string(data), key)
			if err != nil {
				logger.Error("Unable to write admin log: ", err.Error())
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			//Write appropriate headers
			next.ServeHTTP(w, r)
		})
	}
}

// Middleware to auth user
func generateUserMiddleware(db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//Get the header and validate it
			key := r.Header.Get("X-API-Key")
			if key == "" {
				logger.Error("API key missing")
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			privileges := kntdatabase.CheckUserPrivileges(key, db)
			//Allow admins
			if privileges == "admin" || privileges == "user" {
				next.ServeHTTP(w, r)
				return
			}
			//Write appropriate headers
			logger.Error("Unauthorized")
			http.Error(w, http.StatusText(http.StatusProxyAuthRequired), http.StatusProxyAuthRequired)
		})
	}
}
