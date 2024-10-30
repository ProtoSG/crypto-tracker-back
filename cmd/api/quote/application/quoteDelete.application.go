package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type QuoteDelete struct {
	repo domain.QuoteRepository
}

func NewQuoteDelete(repo domain.QuoteRepository) *QuoteDelete {
	return &QuoteDelete{repo}
}

func (this *QuoteDelete) Execute(id int) error {
	if _, err := this.repo.ReadByID(id); err != nil {
		return utils.NewEntityNotFound(id, "Quote")
	}

	return this.repo.Delete(id)
}
