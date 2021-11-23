package server

import (
	"deliveryhero/handler"
	"deliveryhero/helper"
	"deliveryhero/model"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()

	// // serve doc
	// router.PathPrefix("/doc/").Handler(http.StripPrefix("/doc/", http.FileServer(http.Dir("./doc"))))
	router.PathPrefix("/lib/godoc/").Handler(http.StripPrefix("/lib/godoc/", http.FileServer(http.Dir("./lib/godoc"))))

	router.HandleFunc("/", Index).Methods(http.MethodGet)
	router.HandleFunc("/get/{key}", GetKeyHandler()).Methods(http.MethodGet)
	router.HandleFunc("/set", SetKeyHandler()).Methods(http.MethodPost)
	router.HandleFunc("/flushDb", FlushDbHandler()).Methods(http.MethodPost)

	router.Use(Logger(router))

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	_, err := w.Write([]byte("Server running."))
	if err != nil {
		panic(err)
	}
}

func GetKeyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		key := vars["key"]
		fmt.Println(vars)
		h := handler.NewKeyValueHttpHandler()
		result, err := h.HanldeGetKey(key)
		if err != nil {
			helper.ErrorResponseHttp(w, err)
		} else {
			helper.OkResponseHttp(w, result)
		}
	}
}

func SetKeyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var keyValue model.KeyValue
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			helper.ErrorResponseHttp(w, err)
			return
		}
		defer func() {
			if err := r.Body.Close(); err != nil {
				panic(err)
			}
		}()

		if err := json.Unmarshal(body, &keyValue); err != nil {
			helper.ErrorResponseHttp(w, err)
			return
		}

		h := handler.NewKeyValueHttpHandler()
		result, err := h.HanldeSetKey(keyValue)
		if err != nil {
			helper.ErrorResponseHttp(w, err)
		} else {
			helper.OkResponseHttp(w, result)
		}
	}
}

func FlushDbHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h := handler.NewKeyValueHttpHandler()
		result, err := h.HandleFlushDb()
		if err != nil {
			helper.ErrorResponseHttp(w, err)
		} else {
			helper.OkResponseHttp(w, result)
		}
	}
}
