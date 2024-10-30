package application

import (
	"database/sql"

	"github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/domain"
	"github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/infrastructure"
)

type HistorialPriceService struct {
	Create   HistorialPriceCreate
	Read     HistorialPriceRead
	ReadByID HistorialPriceReadByID
	Update   HistorialPriceUpdate
	Delete   HistorialPriceDelete
}

func NewHistorialPriceService(db *sql.DB) *HistorialPriceService {
	var repo domain.HistorialPriceRepository = infrastructure.NewSqliteHistorialPriceRepository(db)

	return &HistorialPriceService{
		Create:   *NewHistorialPriceCreate(repo),
		Read:     *NewHistorialPriceRead(repo),
		ReadByID: *NewHistorialPriceReadByID(repo),
		Update:   *NewHistorialPriceUpdate(repo),
		Delete:   *NewHistorialPriceDelete(repo),
	}
}
