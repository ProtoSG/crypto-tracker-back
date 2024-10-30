package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type HistorialPriceDelete struct {
	repo domain.HistorialPriceRepository
}

func NewHistorialPriceDelete(repo domain.HistorialPriceRepository) *HistorialPriceDelete {
	return &HistorialPriceDelete{repo}
}

func (this *HistorialPriceDelete) Execute(id int) error {
	if historialPrice, _ := this.repo.ReadByID(id); historialPrice == nil {
		return utils.NewEntityNotFound(id, "HistorialPrice")
	}

	return this.repo.Delete(id)
}
