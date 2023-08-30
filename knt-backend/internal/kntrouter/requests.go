package kntrouter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/logger"
)

var validate = validator.New()

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not implemented")
}

func checkAndSendError(w http.ResponseWriter, err error, status int) {
	logger.Error(err.Error())
	w.WriteHeader(status)

	var errData ErrorModel
	errData.Err = err.Error()

	fmt.Fprint(w, generateJsonFromStruct(errData))
}

func generateJsonFromStruct[K any](data K) string {
	jsonString, _ := json.Marshal(data)
	return string(jsonString)
}

func makeAndValidateStruct[K any](reader io.ReadCloser) (K, error) {
	var data K
	json.NewDecoder(reader).Decode(&data)
	err := validate.Struct(data)
	return data, err
}
