package application

import (
	"database/sql"

	"github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/domain"
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/infrastructure"
)

type QuoteService struct {
	Create   *QuoteCreate
	Read     *QuoteRead
	ReadByID *QuoteReadByID
	Update   *QuoteUpdate
	Delete   *QuoteDelete
}

func NewQuoteService(db *sql.DB) *QuoteService {
	var repo domain.QuoteRepository = infrastructure.NewSqliteQuoteRepository(db)

	return &QuoteService{
		Create:   NewQuoteCreate(repo),
		Read:     NewQuoteRead(repo),
		ReadByID: NewQuoteReadByID(repo),
		Update:   NewQuoteUpdate(repo),
		Delete:   NewQuoteDelete(repo),
	}
}
