package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type HistorialPriceUpdate struct {
	repo domain.HistorialPriceRepository
}

func NewHistorialPriceUpdate(repo domain.HistorialPriceRepository) *HistorialPriceUpdate {
	return &HistorialPriceUpdate{repo}
}

func (this *HistorialPriceUpdate) Execute(id int, historialPrice *domain.HistorialPrice) error {
	if historialPrice, _ := this.repo.ReadByID(id); historialPrice == nil {
		return utils.NewEntityNotFound(id, "HistorialPrice")
	}

	newHistorialPrice := domain.NewHistorialPrice(id, historialPrice.IDCrypto, historialPrice.Price, historialPrice.Date)
	return this.repo.Update(newHistorialPrice)
}
