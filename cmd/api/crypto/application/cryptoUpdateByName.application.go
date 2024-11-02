package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type CryptoUpdateByName struct {
	repo domain.CryptoRepository
}

func NewCryptoUpdateByName(repo domain.CryptoRepository) *CryptoUpdateByName {
	return &CryptoUpdateByName{repo}
}

func (this *CryptoUpdateByName) Execute(name string, crypto *domain.Crypto) error {
	if crypto, _ := this.repo.ReadByName(name); crypto == nil {
		return utils.NewEntityNotFound(0, "Crypto")
	}

	newCrypto := domain.NewCrypto(crypto.ID, crypto.Name, crypto.Symbol, crypto.Slug, crypto.CirculatingSupply, crypto.CmcRank)
	return this.repo.Update(newCrypto)
}
