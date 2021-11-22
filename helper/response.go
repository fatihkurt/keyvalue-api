package helper

import (
	"encoding/json"
	"net/http"
)

type ResponseHttp struct {
	Result interface{} `json:"result"`
	Error  *string     `json:"error"`
}

func OkResponseHttp(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	response := ResponseHttp{Error: nil, Result: result}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

func ErrorResponseHttp(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(GetStatusCodeWithDefault(err))
	errMsg := err.Error()

	response := ResponseHttp{Error: &errMsg, Result: nil}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
