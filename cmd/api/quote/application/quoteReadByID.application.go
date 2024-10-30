package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type QuoteReadByID struct {
	repo domain.QuoteRepository
}

func NewQuoteReadByID(repo domain.QuoteRepository) *QuoteReadByID {
	return &QuoteReadByID{repo}
}

func (this *QuoteReadByID) Execute(id int) (*domain.Quote, error) {
	quote, _ := this.repo.ReadByID(id)
	if quote == nil {
		return nil, utils.NewEntityNotFound(id, "Quote")
	}

	return quote, nil
}
