package application

import (
	"database/sql"

	cryptoDomain "github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/domain"
	cryptoInfrastructure "github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/infrastructure"
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/domain"
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/infrastructure"
)

type QuoteService struct {
	Create           *QuoteCreate
	Read             *QuoteRead
	ReadByID         *QuoteReadByID
	ReadByIDCrypto   *QuoteReadByIDCrypto
	Update           *QuoteUpdate
	UpdateByIDCrypto *QuoteUpdateByIDCrypto
	Delete           *QuoteDelete
}

func NewQuoteService(db *sql.DB) *QuoteService {
	var repo domain.QuoteRepository = infrastructure.NewSqliteQuoteRepository(db)
	var cryptoRepo cryptoDomain.CryptoRepository = cryptoInfrastructure.NewSqliteCryptoRepository(db)

	return &QuoteService{
		Create:           NewQuoteCreate(repo),
		Read:             NewQuoteRead(repo),
		ReadByID:         NewQuoteReadByID(repo),
		ReadByIDCrypto:   NewQuoteReadByIDCrypto(repo),
		Update:           NewQuoteUpdate(repo),
		UpdateByIDCrypto: NewQuoteUpdateByIDCrypto(repo, cryptoRepo),
		Delete:           NewQuoteDelete(repo),
	}
}
