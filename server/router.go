package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phati/circleci-demo/handler"
)

func InitRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/math/add/{num1}/{num2}", handler.AddHandler).Methods(http.MethodGet)
	router.HandleFunc("/math/sub/{num1}/{num2}", handler.SubHandler).Methods(http.MethodGet)
	return
}
