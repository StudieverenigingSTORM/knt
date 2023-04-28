package kntrouter

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func generateJsonResponse[K any](data K, err error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		jsonString, _ := json.Marshal(data)
		fmt.Fprintf(w, string(jsonString))
	}
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented")
}
