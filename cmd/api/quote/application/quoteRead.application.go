package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type QuoteRead struct {
	repo domain.QuoteRepository
}

func NewQuoteRead(repo domain.QuoteRepository) *QuoteRead {
	return &QuoteRead{repo}
}

func (this *QuoteRead) Execute() ([]*domain.Quote, error) {
	quotes, err := this.repo.Read()
	if err != nil {
		return nil, utils.NewEntityNotFound(0, "Quote")
	}

	return quotes, nil
}
