package web

import (
	"encoding/json"
	"net/http"

	"github.com/evandrobarbosadosreis/go-rest-development/application"
	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service application.CustomerService
}

func (handler *CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	if result, err := handler.service.GetAllCustomers(r.URL.Query().Get("status")); err != nil {
		writeResponse(w, err.Code, err)
	} else {
		writeResponse(w, http.StatusOK, result)
	}
}

func (handler *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	if result, err := handler.service.GetCustomer(mux.Vars(r)["id"]); err != nil {
		writeResponse(w, err.Code, err)
	} else {
		writeResponse(w, http.StatusOK, result)
	}
}

func writeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
