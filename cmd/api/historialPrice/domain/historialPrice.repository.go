package domain

type HistorialPriceRepository interface {
	Create(*HistorialPrice) error
	Read() ([]*HistorialPrice, error)
	ReadByID(id int) (*HistorialPrice, error)
	ReadByIDCrypto(id int) (*HistorialPrice, error)
	Update(historialPrice *HistorialPrice) error
	UpdateByIDCrypto(historialPrice *HistorialPrice) error
	Delete(id int) error
}
