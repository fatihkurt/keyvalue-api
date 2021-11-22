package server

import (
	"deliveryhero/handler"
	"deliveryhero/helper"
	"deliveryhero/model"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	client := redis.NewClient(&redis.Options{
		Addr:     helper.GetEnv("REDIS_URL", "localhost:6379"),
		Password: helper.GetEnv("REDIS_PASSWORD", ""),
		DB:       0,
	})

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/get", GetKeyHandler(client)).Methods(http.MethodGet)
	router.HandleFunc("/api/set", SetKeyHandler(client)).Methods(http.MethodPost)

	router.Use(Logger(router))

	return router
}

func GetKeyHandler(client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := handler.NewKeyValueHttpHandler(client)
		vars := mux.Vars(r)
		key := vars["key"]
		result, err := h.HanldeGetKey(key)
		if err != nil {
			helper.ErrorResponseHttp(w, err)
		} else {
			helper.OkResponseHttp(w, result)
		}
	}
}

func SetKeyHandler(client *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var keyValue model.KeyValue
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}
		if err := json.Unmarshal(body, &keyValue); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			if err := json.NewEncoder(w).Encode(err); err != nil {
				helper.ErrorResponseHttp(w, err)
			}
		}

		h := handler.NewKeyValueHttpHandler(client)
		result, err := h.HanldeSetKey(keyValue)
		if err != nil {
			helper.ErrorResponseHttp(w, err)
		} else {
			helper.OkResponseHttp(w, result)
		}
	}
}
