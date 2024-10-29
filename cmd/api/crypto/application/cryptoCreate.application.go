package application

import "github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/domain"

type CryptoCreate struct {
	repo domain.CryptoRepository
}

func NewCryptoCreate(repo domain.CryptoRepository) *CryptoCreate {
	return &CryptoCreate{repo}
}

func (this *CryptoCreate) Execute(id int, name, symbol, slug string, circulatingSupply float32, cmcRank int) error {
	crypto := domain.NewCrypto(id, name, symbol, slug, circulatingSupply, cmcRank)

	return this.repo.Create(crypto)
}
