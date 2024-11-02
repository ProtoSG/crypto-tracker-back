package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/domain"
)

type QuoteCreate struct {
	repo domain.QuoteRepository
}

func NewQuoteCreate(repo domain.QuoteRepository) *QuoteCreate {
	return &QuoteCreate{repo}
}

func (this *QuoteCreate) Execute(id int, idCrypto int, price, volume24h, percentChange1h, percentChange24h, percentChange7d float64) error {
	newQuote := domain.NewQuote(id, idCrypto, price, volume24h, percentChange1h, percentChange24h, percentChange7d)
	return this.repo.Create(newQuote)
}
