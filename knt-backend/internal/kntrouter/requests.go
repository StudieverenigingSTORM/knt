package kntrouter

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func generateJsonResponse[K any](data K, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	jsonString, _ := json.Marshal(data)
	return jsonString, nil
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented")
}
