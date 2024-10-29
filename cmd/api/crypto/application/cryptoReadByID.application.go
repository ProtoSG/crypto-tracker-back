package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type CryptoReadByID struct {
	repo domain.CryptoRepository
}

func NewCryptoReadByID(repo domain.CryptoRepository) *CryptoReadByID {
	return &CryptoReadByID{repo}
}

func (this *CryptoReadByID) Execute(id int) (*domain.Crypto, error) {
	crypto, _ := this.repo.ReadByID(id)
	if crypto == nil {
		return nil, utils.NewEntityNotFound(id, "Crypto")
	}

	return crypto, nil
}
