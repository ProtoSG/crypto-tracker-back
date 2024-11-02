package application

import (
	"database/sql"

	"github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/domain"
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/infrastructure"
)

type CryptoService struct {
	Create       CryptoCreate
	Read         CryptoRead
	ReadByID     CryptoReadByID
	ReadByName   CryptoReadByName
	Update       CryptoUpdate
	UpdateByName CryptoUpdateByName
	Delete       CryptoDelete
}

func NewCryptoService(db *sql.DB) *CryptoService {
	var repo domain.CryptoRepository = infrastructure.NewSqliteCryptoRepository(db)

	return &CryptoService{
		Create:       *NewCryptoCreate(repo),
		Read:         *NewCryptoRead(repo),
		ReadByID:     *NewCryptoReadByID(repo),
		ReadByName:   *NewCryptoReadByName(repo),
		Update:       *NewCryptoUpdate(repo),
		UpdateByName: *NewCryptoUpdateByName(repo),
		Delete:       *NewCryptoDelete(repo),
	}
}
