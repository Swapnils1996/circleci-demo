package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/phati/circleci-demo/api"
	"github.com/phati/circleci-demo/service"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	num1, err := strconv.Atoi(strings.Trim(mux.Vars(r)["num1"], " "))
	if err != nil {
		api.RespondWithJSON(w, http.StatusBadRequest, api.Response{
			Error: "invalid request",
		})
	}

	num2, err := strconv.Atoi(strings.Trim(mux.Vars(r)["num2"], " "))
	if err != nil {
		api.RespondWithJSON(w, http.StatusBadRequest, api.Response{
			Error: "invalid request",
		})
	}

	result := service.Add(num1, num2)

	api.RespondWithJSON(w, http.StatusOK, api.Response{
		Message: "success",
		Data:    result,
	})
}
func SubHandler(w http.ResponseWriter, r *http.Request) {
	num1, err := strconv.Atoi(strings.Trim(mux.Vars(r)["num1"], " "))
	if err != nil {
		api.RespondWithJSON(w, http.StatusBadRequest, api.Response{
			Error: "invalid request",
		})
	}

	num2, err := strconv.Atoi(strings.Trim(mux.Vars(r)["num2"], " "))
	if err != nil {
		api.RespondWithJSON(w, http.StatusBadRequest, api.Response{
			Error: "invalid request",
		})
	}

	result := service.Sub(num1, num2)

	api.RespondWithJSON(w, http.StatusOK, api.Response{
		Message: "success",
		Data:    result,
	})
}
