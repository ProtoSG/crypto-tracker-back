package application

import (
	domainCrypto "github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/domain"
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type HistorialPriceUpdateByIDCrypto struct {
	repo       domain.HistorialPriceRepository
	cryptoRepo domainCrypto.CryptoRepository
}

func NewHistorialPriceUpdateByIDCrypto(repo domain.HistorialPriceRepository, cryptoRepo domainCrypto.CryptoRepository) *HistorialPriceUpdateByIDCrypto {
	return &HistorialPriceUpdateByIDCrypto{repo, cryptoRepo}
}

func (this *HistorialPriceUpdateByIDCrypto) Execute(id int, historialPrice *domain.HistorialPrice) error {
	if crypto, _ := this.cryptoRepo.ReadByID(id); crypto == nil {
		return utils.NewEntityNotFound(id, "Crypto")
	}

	if historialPrice, _ := this.repo.ReadByIDCrypto(id); historialPrice == nil {
		return utils.NewEntityNotFound(id, "HistorialPrice")
	}

	newHistorialPrice := domain.NewHistorialPrice(id, historialPrice.IDCrypto, historialPrice.Price, historialPrice.Date)
	return this.repo.UpdateByIDCrypto(newHistorialPrice)
}
