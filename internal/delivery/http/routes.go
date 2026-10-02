package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handlers struct {
}

func InitRoutes(h Handlers) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}).Methods(http.MethodGet)

	return r
}