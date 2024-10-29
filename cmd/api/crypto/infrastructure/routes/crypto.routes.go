package routes

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/infrastructure/controller"
	"github.com/ProtoSG/crypto-tracker-back/internal/services"
	"github.com/gorilla/mux"
)

func NewCryptoRoutes(r *mux.Router, serviceContainer *services.ServiceContainer) {
	controller := controller.NewCryptoController(serviceContainer)

	r.HandleFunc("/crypto", controller.Create).Methods("POST")
	r.HandleFunc("/crypto", controller.Read).Methods("GET")
	r.HandleFunc("/crypto/{id}", controller.ReadByID).Methods("GET")
	r.HandleFunc("/crypto/{id}", controller.Update).Methods("PUT")
	r.HandleFunc("/crypto/{id}", controller.Delete).Methods("DELETE")
}
