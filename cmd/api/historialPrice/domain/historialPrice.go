package domain

type HistorialPrice struct {
	ID       int     `json:"id"`
	IDCrypto int     `json:"id_crypto"`
	Price    float64 `json:"price"`
	Date     string  `json:"date"`
}

func NewHistorialPrice(id int, idCrypto int, price float64, date string) *HistorialPrice {
	return &HistorialPrice{
		ID:       id,
		IDCrypto: idCrypto,
		Price:    price,
		Date:     date,
	}
}
