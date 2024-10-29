package routes

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/infrastructure/routes"
	"github.com/ProtoSG/crypto-tracker-back/internal/services"
	"github.com/gorilla/mux"
)

func NewRoutes(serviceContainer *services.ServiceContainer) *mux.Router {
	r := mux.NewRouter()

	routes.NewCryptoRoutes(r, serviceContainer)

	return r
}
