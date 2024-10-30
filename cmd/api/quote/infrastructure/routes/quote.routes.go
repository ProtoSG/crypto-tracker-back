package routes

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/infrastructure/controller"
	"github.com/ProtoSG/crypto-tracker-back/internal/services"
	"github.com/gorilla/mux"
)

func NewQuoteRoutes(r *mux.Router, serviceContainer *services.ServiceContainer) {
	controller := controller.NewQuoteController(serviceContainer)

	r.HandleFunc("/quote", controller.Create).Methods("POST")
	r.HandleFunc("/quote", controller.Read).Methods("GET")
	r.HandleFunc("/quote/{id}", controller.ReadByID).Methods("GET")
	r.HandleFunc("/quote/{id}", controller.Update).Methods("PUT")
	r.HandleFunc("/quote/{id}", controller.Delete).Methods("DELETE")
}
