package application

import (
	domainCrypto "github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/domain"
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type QuoteUpdateByIDCrypto struct {
	repo       domain.QuoteRepository
	cryptoRepo domainCrypto.CryptoRepository
}

func NewQuoteUpdateByIDCrypto(repo domain.QuoteRepository, cryptoRepo domainCrypto.CryptoRepository) *QuoteUpdateByIDCrypto {
	return &QuoteUpdateByIDCrypto{repo, cryptoRepo}
}

func (this *QuoteUpdateByIDCrypto) Execute(id int, quote *domain.Quote) error {
	if crypto, _ := this.cryptoRepo.ReadByID(id); crypto == nil {
		return utils.NewEntityNotFound(id, "Crypto")
	}

	if quote, _ := this.repo.ReadByIDCrypto(id); quote == nil {
		return utils.NewEntityNotFound(id, "Quote")
	}

	newQuote := domain.NewQuote(quote.ID, quote.IDCrypto, quote.Price, quote.Volume24h, quote.PercentChange1h, quote.PercentChange24h, quote.PercentChange7d)
	return this.repo.UpdateByIDCrypto(newQuote)
}
