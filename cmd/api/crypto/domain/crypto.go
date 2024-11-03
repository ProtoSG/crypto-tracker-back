package domain

import (
	domainQuote "github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/domain"
)

type Crypto struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	Symbol            string  `json:"symbol"`
	Slug              string  `json:"slug"`
	CirculatingSupply float32 `json:"circulating_supply"`
	CmcRank           int     `json:"cmc_rank"`
}

func NewCrypto(id int, name, symbol, slug string, circulatingSupply float32, cmcRank int) *Crypto {
	return &Crypto{
		ID:                id,
		Name:              name,
		Symbol:            symbol,
		Slug:              slug,
		CirculatingSupply: circulatingSupply,
		CmcRank:           cmcRank,
	}
}

type CryptoResponse struct {
	Crypto
	Quote struct {
		USD domainQuote.Quote `json:"USD"`
	} `json:"quote"`
}
