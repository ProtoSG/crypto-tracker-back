package services

import (
	"database/sql"

	"github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/application"
)

type ServiceContainer struct {
	Crypto *application.CryptoService
}

func NewServiceContainer(db *sql.DB) *ServiceContainer {
	return &ServiceContainer{
		Crypto: application.NewCryptoService(db),
	}
}
