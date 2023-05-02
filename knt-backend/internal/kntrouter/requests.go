package kntrouter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/logger"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func generateJsonResponse[K any](data K, err error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			logger.Error("Unable to generate json response: ", err.Error())
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		jsonString, _ := json.Marshal(data)
		fmt.Fprint(w, string(jsonString))
	}
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented")
}
