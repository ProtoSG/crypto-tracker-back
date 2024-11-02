package routes

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/infrastructure/controller"
	"github.com/ProtoSG/crypto-tracker-back/internal/services"
	"github.com/gorilla/mux"
)

func NewHistorialPriceRoutes(r *mux.Router, serviceContainer *services.ServiceContainer) {
	controller := controller.NewHistorialPriceController(serviceContainer)

	r.HandleFunc("/historialPrice", controller.Create).Methods("POST")
	r.HandleFunc("/historialPrice", controller.Read).Methods("GET")
	r.HandleFunc("/historialPrice/{id}", controller.ReadByID).Methods("GET")
	r.HandleFunc("/historialPrice/crypto/{id}", controller.ReadByIDCrypto).Methods("GET")
	r.HandleFunc("/historialPrice/{id}", controller.Update).Methods("PUT")
	r.HandleFunc("/historialPrice/crypto/{id}", controller.UpdateByIDCrypto).Methods("PUT")
	r.HandleFunc("/historialPrice/{id}", controller.Delete).Methods("DELETE")
}
