package application

import (
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/domain"
)

type HistorialPriceCreate struct {
	repo domain.HistorialPriceRepository
}

func NewHistorialPriceCreate(repo domain.HistorialPriceRepository) *HistorialPriceCreate {
	return &HistorialPriceCreate{repo}
}

func (this *HistorialPriceCreate) Execute(id, idCrypto int, price float64, date string) error {
	newHistorialPrice := domain.NewHistorialPrice(id, idCrypto, price, date)
	return this.repo.Create(newHistorialPrice)
}
