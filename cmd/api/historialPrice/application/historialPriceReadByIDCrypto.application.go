package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type HistorialPriceReadByIDCrypto struct {
	repo domain.HistorialPriceRepository
}

func NewHistorialPriceReadByIDCrypto(repo domain.HistorialPriceRepository) *HistorialPriceReadByIDCrypto {
	return &HistorialPriceReadByIDCrypto{repo}
}

func (this *HistorialPriceReadByIDCrypto) Execute(id int) (*domain.HistorialPrice, error) {
	historialPrice, _ := this.repo.ReadByIDCrypto(id)
	if historialPrice == nil {
		return nil, utils.NewEntityNotFound(id, "HistorialPrice")
	}

	return historialPrice, nil
}
