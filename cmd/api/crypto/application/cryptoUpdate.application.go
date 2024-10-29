package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type CryptoUpdate struct {
	repo domain.CryptoRepository
}

func NewCryptoUpdate(repo domain.CryptoRepository) *CryptoUpdate {
	return &CryptoUpdate{repo}
}

func (this *CryptoUpdate) Execute(id int, crypto *domain.Crypto) error {
	if crypto, _ := this.repo.ReadByID(id); crypto == nil {
		return utils.NewEntityNotFound(id, "Crypto")
	}

	newCrypto := domain.NewCrypto(id, crypto.Name, crypto.Symbol, crypto.Slug, crypto.CirculatingSupply, crypto.CmcRank)
	return this.repo.Update(newCrypto)
}
