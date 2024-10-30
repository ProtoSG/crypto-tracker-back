package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type HistorialPriceRead struct {
	repo domain.HistorialPriceRepository
}

func NewHistorialPriceRead(repo domain.HistorialPriceRepository) *HistorialPriceRead {
	return &HistorialPriceRead{repo}
}

func (this *HistorialPriceRead) Execute() ([]*domain.HistorialPrice, error) {
	historialPrices, err := this.repo.Read()
	if err != nil {
		return nil, utils.NewEntityNotFound(0, "HistorialPrice")
	}

	return historialPrices, nil
}
