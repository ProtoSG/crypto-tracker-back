package routes

import (
	cryptoRoutes "github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/infrastructure/routes"
	historialPriceRoutes "github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/infrastructure/routes"
	quoteRoutes "github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/infrastructure/routes"
	"github.com/ProtoSG/crypto-tracker-back/internal/services"
	"github.com/gorilla/mux"
)

func NewRoutes(serviceContainer *services.ServiceContainer) *mux.Router {
	r := mux.NewRouter()

	cryptoRoutes.NewCryptoRoutes(r, serviceContainer)
	quoteRoutes.NewQuoteRoutes(r, serviceContainer)
	historialPriceRoutes.NewHistorialPriceRoutes(r, serviceContainer)

	return r
}
