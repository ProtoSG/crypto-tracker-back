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
	query := `INSERT INTO crypto (name, symbol, slug, circulating_supply, cmc_rank) 
    VALUES (?, ?, ? ,?, ?)
  `
	_, err := this.db.Exec(
		query,
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

func (this *SqliteCryptoRepository) Read() ([]*domain.Crypto, error) {
	query := `SELECT * FROM crypto`

	rows, err := this.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cryptos []*domain.Crypto
	for rows.Next() {
		crypto := &domain.Crypto{}
		if err := rows.Scan(
			&crypto.ID,
			&crypto.Name,
			&crypto.Symbol,
			&crypto.Slug,
			&crypto.CirculatingSupply,
			&crypto.CmcRank,
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

func (this *SqliteCryptoRepository) Delete(id int) error {
	query := `DELETE FROM crypt WHERE id_crypto = ?`

	if _, err := this.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}
