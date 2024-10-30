package services

import (
	"database/sql"

	cryptoApplication "github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/application"
	historialPriceApplication "github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/application"
	quoteApplication "github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/application"
)

type ServiceContainer struct {
	Crypto         *cryptoApplication.CryptoService
	Quote          *quoteApplication.QuoteService
	HistorialPrice *historialPriceApplication.HistorialPriceService
}

func NewServiceContainer(db *sql.DB) *ServiceContainer {
	return &ServiceContainer{
		Crypto:         cryptoApplication.NewCryptoService(db),
		Quote:          quoteApplication.NewQuoteService(db),
		HistorialPrice: historialPriceApplication.NewHistorialPriceService(db),
	}
}
