package application

import "github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/domain"

type CryptoRead struct {
	repo domain.CryptoRepository
}

func NewCryptoRead(repo domain.CryptoRepository) *CryptoRead {
	return &CryptoRead{repo}
}

func (this *CryptoRead) Execute() ([]*domain.CryptoResponse, error) {
	return this.repo.Read()
}
