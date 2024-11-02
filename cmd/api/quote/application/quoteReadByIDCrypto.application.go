package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type QuoteReadByIDCrypto struct {
	repo domain.QuoteRepository
}

func NewQuoteReadByIDCrypto(repo domain.QuoteRepository) *QuoteReadByIDCrypto {
	return &QuoteReadByIDCrypto{repo}
}

func (this *QuoteReadByIDCrypto) Execute(id int) (*domain.Quote, error) {
	quote, _ := this.repo.ReadByIDCrypto(id)
	if quote == nil {
		return nil, utils.NewEntityNotFound(id, "Quote")
	}

	return quote, nil
}
