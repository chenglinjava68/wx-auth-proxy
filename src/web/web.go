package web

import (
	"conf"
	"log"
	"net/http"
	"web/api"

	"github.com/gorilla/mux"
)

func Start() {
	log.Println("web starts")
	r := mux.NewRouter()

	r.HandleFunc("/", api.ProxyHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(conf.Conf.Listen, r))
}
