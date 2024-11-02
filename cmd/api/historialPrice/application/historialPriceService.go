package application

import (
	"database/sql"

	cryptoDomain "github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/domain"
	cryptoInfrastructure "github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/infrastructure"
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/domain"
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/infrastructure"
)

type HistorialPriceService struct {
	Create           HistorialPriceCreate
	Read             HistorialPriceRead
	ReadByID         HistorialPriceReadByID
	ReadByIDCrypto   HistorialPriceReadByIDCrypto
	Update           HistorialPriceUpdate
	UpdateByIDCrypto HistorialPriceUpdateByIDCrypto
	Delete           HistorialPriceDelete
}

func NewHistorialPriceService(db *sql.DB) *HistorialPriceService {
	var repo domain.HistorialPriceRepository = infrastructure.NewSqliteHistorialPriceRepository(db)
	var cryptoRepo cryptoDomain.CryptoRepository = cryptoInfrastructure.NewSqliteCryptoRepository(db)

	return &HistorialPriceService{
		Create:           *NewHistorialPriceCreate(repo),
		Read:             *NewHistorialPriceRead(repo),
		ReadByID:         *NewHistorialPriceReadByID(repo),
		ReadByIDCrypto:   *NewHistorialPriceReadByIDCrypto(repo),
		Update:           *NewHistorialPriceUpdate(repo),
		UpdateByIDCrypto: *NewHistorialPriceUpdateByIDCrypto(repo, cryptoRepo),
		Delete:           *NewHistorialPriceDelete(repo),
	}
}
