package application

import (
	"fmt"

	"github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type CryptoReadByName struct {
	repo domain.CryptoRepository
}

func NewCryptoReadByName(repo domain.CryptoRepository) *CryptoReadByName {
	return &CryptoReadByName{repo}
}

func (this *CryptoReadByName) Execute(name string) (*domain.Crypto, error) {
	crypto, _ := this.repo.ReadByName(name)
	if crypto == nil {
		return nil, utils.NewEntityNotFound(0, fmt.Sprintf("Crypto %s", crypto.Name))
	}

	return crypto, nil
}
