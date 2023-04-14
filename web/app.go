package web

import (
	"log"
	"net/http"

	"github.com/evandrobarbosadosreis/go-rest-development/application"
	"github.com/evandrobarbosadosreis/go-rest-development/infrastructure"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	handler := CustomerHandler{
		service: application.NewDefaultCustomerService(infrastructure.NewMySqlRepository()),
	}

	router.HandleFunc("/customers", handler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", handler.GetCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", router))
}
