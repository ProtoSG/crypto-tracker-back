package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type HistorialPriceCreate struct {
	repo domain.HistorialPriceRepository
}

func NewHistorialPriceCreate(repo domain.HistorialPriceRepository) *HistorialPriceCreate {
	return &HistorialPriceCreate{repo}
}

func (this *HistorialPriceCreate) Execute(id int, historialPrice *domain.HistorialPrice) error {
	if historialPrice, _ := this.repo.ReadByID(id); historialPrice == nil {
		return utils.NewEntityNotFound(id, "HistorialPrice")
	}

	newHistorialPrice := domain.NewHistorialPrice(id, historialPrice.IDCrypto, historialPrice.Price, historialPrice.Date)
	return this.repo.Create(newHistorialPrice)
}
