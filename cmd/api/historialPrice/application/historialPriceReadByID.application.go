package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type HistorialPriceReadByID struct {
	repo domain.HistorialPriceRepository
}

func NewHistorialPriceReadByID(repo domain.HistorialPriceRepository) *HistorialPriceReadByID {
	return &HistorialPriceReadByID{repo}
}

func (this *HistorialPriceReadByID) Execute(id int) (*domain.HistorialPrice, error) {
	historialPrice, _ := this.repo.ReadByID(id)
	if historialPrice == nil {
		return nil, utils.NewEntityNotFound(id, "HistorialPrice")
	}

	return historialPrice, nil
}
