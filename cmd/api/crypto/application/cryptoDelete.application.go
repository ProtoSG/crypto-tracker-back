package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type CryptoDelete struct {
	repo domain.CryptoRepository
}

func NewCryptoDelete(repo domain.CryptoRepository) *CryptoDelete {
	return &CryptoDelete{repo}
}

func (this *CryptoDelete) Execute(id int) error {
	if crypto, _ := this.repo.ReadByID(id); crypto == nil {
		return utils.NewEntityNotFound(id, "Crypto")
	}

	return this.repo.Delete(id)
}
