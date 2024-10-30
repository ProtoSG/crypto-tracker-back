package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type QuoteCreate struct {
	repo domain.QuoteRepository
}

func NewQuoteCreate(repo domain.QuoteRepository) *QuoteCreate {
	return &QuoteCreate{repo}
}

func (this *QuoteCreate) Execute(id int, quote *domain.Quote) error {
	if quote, _ := this.repo.ReadByID(id); quote == nil {
		return utils.NewEntityNotFound(id, "Quote")
	}

	newQuote := domain.NewQuote(id, quote.IDCrypto, quote.Price, quote.Volume24h, quote.PercentChange1h, quote.PercentChange24h, quote.PercentChange7d)
	return this.repo.Update(newQuote)
}
