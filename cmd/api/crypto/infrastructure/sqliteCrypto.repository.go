package infrastructure

import (
	"database/sql"

	"github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/domain"
)

type SqliteCryptoRepository struct {
	db *sql.DB
}

func NewSqliteCryptoRepository(db *sql.DB) *SqliteCryptoRepository {
	return &SqliteCryptoRepository{db}
}

func (this *SqliteCryptoRepository) Create(crypto *domain.Crypto) error {
	query := `INSERT INTO crypto (id_crypto, name, symbol, slug, circulating_supply, cmc_rank) 
    VALUES (?, ?, ?, ? ,?, ?)
  `
	_, err := this.db.Exec(
		query,
		crypto.ID,
		crypto.Name,
		crypto.Symbol,
		crypto.Slug,
		crypto.CirculatingSupply,
		crypto.CmcRank,
	)
	if err != nil {
		return err
	}

	return nil
}

func (this *SqliteCryptoRepository) Read() ([]*domain.CryptoResponse, error) {
	query := `
    SELECT
      c.id_crypto,
      c.name,
      c.symbol,
      c.slug,
      c.circulating_supply,
      c.cmc_rank,
      q.id_quote,
      q.id_crypto,
      q.price,
      q.volume_24h,
      q.percent_change_1h,
      q.percent_change_24h,
      q.percent_change_7d
    FROM crypto c 
    INNER JOIN quote q ON c.id_crypto = q.id_crypto;
  `

	rows, err := this.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cryptos []*domain.CryptoResponse
	for rows.Next() {
		crypto := &domain.CryptoResponse{}
		if err := rows.Scan(
			&crypto.ID,
			&crypto.Name,
			&crypto.Symbol,
			&crypto.Slug,
			&crypto.CirculatingSupply,
			&crypto.CmcRank,
			&crypto.Quote.USD.ID,
			&crypto.Quote.USD.IDCrypto,
			&crypto.Quote.USD.Price,
			&crypto.Quote.USD.Volume24h,
			&crypto.Quote.USD.PercentChange1h,
			&crypto.Quote.USD.PercentChange24h,
			&crypto.Quote.USD.PercentChange7d,
		); err != nil {
			return nil, err
		}

		cryptos = append(cryptos, crypto)
	}

	return cryptos, nil
}

func (this *SqliteCryptoRepository) ReadByID(id int) (*domain.Crypto, error) {
	query := `SELECT * FROM crypto WHERE id_crypto = ?`

	crypto := &domain.Crypto{}

	if err := this.db.QueryRow(
		query,
		id,
	).Scan(
		&crypto.ID,
		&crypto.Name,
		&crypto.Symbol,
		&crypto.Slug,
		&crypto.CirculatingSupply,
		&crypto.CmcRank,
	); err != nil {
		return nil, err
	}

	return crypto, nil
}

func (this *SqliteCryptoRepository) ReadByName(name string) (*domain.Crypto, error) {
	query := `SELECT * FROM crypto WHERE name = ?`

	crypto := &domain.Crypto{}

	if err := this.db.QueryRow(
		query,
		name,
	).Scan(
		&crypto.ID,
		&crypto.Name,
		&crypto.Symbol,
		&crypto.Slug,
		&crypto.CirculatingSupply,
		&crypto.CmcRank,
	); err != nil {
		return nil, err
	}

	return crypto, nil
}

func (this *SqliteCryptoRepository) Update(crypto *domain.Crypto) error {
	query := `UPDATE crypto SET name = ?, symbol = ?, slug = ?, circulating_supply = ?, cmc_rank = ? WHERE id_crypto = ?`

	if _, err := this.db.Exec(
		query,
		crypto.Name,
		crypto.Symbol,
		crypto.Slug,
		crypto.CirculatingSupply,
		crypto.CmcRank,
		crypto.ID,
	); err != nil {
		return err
	}

	return nil
}

func (this *SqliteCryptoRepository) UpdateByName(crypto *domain.Crypto) error {
	query := `UPDATE crypto SET symbol = ?, slug = ?, circulating_supply = ?, cmc_rank = ? WHERE name = ?`

	if _, err := this.db.Exec(
		query,
		crypto.Symbol,
		crypto.Slug,
		crypto.CirculatingSupply,
		crypto.CmcRank,
		crypto.Name,
	); err != nil {
		return err
	}

	return nil
}

func (this *SqliteCryptoRepository) Delete(id int) error {
	query := `DELETE FROM crypto WHERE id_crypto = ?`

	if _, err := this.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}
