package web

import (
	"wx-auth-proxy/conf"
	"wx-auth-proxy/web/api"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	log.Println("web starts")
	r := mux.NewRouter()

	r.HandleFunc("/", api.ProxyHandler).Methods(http.MethodGet)
	r.HandleFunc("/user", api.UserHandler).Methods(http.MethodGet)
	r.HandleFunc("/user/info", api.UserInfoHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(conf.Conf.Listen, r))
}
