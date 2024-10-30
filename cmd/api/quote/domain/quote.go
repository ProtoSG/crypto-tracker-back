package domain

type Quote struct {
	ID               int     `json:"id"`
	IDCrypto         int     `json:"id_crypto"`
	Price            float64 `json:"price"`
	Volume24h        float64 `json:"volume_24h"`
	PercentChange1h  float64 `json:"percent_change_1h"`
	PercentChange24h float64 `json:"percent_change_24h"`
	PercentChange7d  float64 `json:"percent_change_7d"`
}

func NewQuote(id int, idCrypto int, price float64, volume24h float64, percentChange1h float64, percentChange24h float64, percentChange7d float64) *Quote {
	return &Quote{
		ID:               id,
		IDCrypto:         idCrypto,
		Price:            price,
		Volume24h:        volume24h,
		PercentChange1h:  percentChange1h,
		PercentChange24h: percentChange24h,
		PercentChange7d:  percentChange7d,
	}
}
