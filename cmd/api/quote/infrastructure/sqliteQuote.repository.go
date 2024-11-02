package infrastructure

import (
	"database/sql"

	"github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/domain"
)

type SqliteQuoteRepository struct {
	db *sql.DB
}

func NewSqliteQuoteRepository(db *sql.DB) *SqliteQuoteRepository {
	return &SqliteQuoteRepository{db}
}

func (this *SqliteQuoteRepository) Create(quote *domain.Quote) error {
	query := `INSERT INTO quote (id_crypto, price, volume_24h, percent_change_1h, percent_change_24h, percent_change_7d) 
    VALUES (?, ?, ?, ?, ?, ?)
  `
	_, err := this.db.Exec(
		query,
		quote.IDCrypto,
		quote.Price,
		quote.Volume24h,
		quote.PercentChange1h,
		quote.PercentChange24h,
		quote.PercentChange7d,
	)
	if err != nil {
		return err
	}

	return nil
}

func (this *SqliteQuoteRepository) Read() ([]*domain.Quote, error) {
	query := `SELECT * FROM quote`

	rows, err := this.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []*domain.Quote
	for rows.Next() {
		quote := &domain.Quote{}
		if err := rows.Scan(
			&quote.ID,
			&quote.IDCrypto,
			&quote.Price,
			&quote.Volume24h,
			&quote.PercentChange1h,
			&quote.PercentChange24h,
			&quote.PercentChange7d,
		); err != nil {
			return nil, err
		}

		quotes = append(quotes, quote)
	}

	return quotes, nil
}

func (this *SqliteQuoteRepository) ReadByID(id int) (*domain.Quote, error) {
	query := `SELECT * FROM quote WHERE id_quote = ?`

	quote := &domain.Quote{}

	if err := this.db.QueryRow(
		query,
		id,
	).Scan(
		&quote.ID,
		&quote.IDCrypto,
		&quote.Price,
		&quote.Volume24h,
		&quote.PercentChange1h,
		&quote.PercentChange24h,
		&quote.PercentChange7d,
	); err != nil {
		return nil, err
	}

	return quote, nil
}

func (this *SqliteQuoteRepository) ReadByIDCrypto(id int) (*domain.Quote, error) {
	query := `SELECT * FROM quote WHERE id_crypto = ?`

	quote := &domain.Quote{}

	if err := this.db.QueryRow(
		query,
		id,
	).Scan(
		&quote.ID,
		&quote.IDCrypto,
		&quote.Price,
		&quote.Volume24h,
		&quote.PercentChange1h,
		&quote.PercentChange24h,
		&quote.PercentChange7d,
	); err != nil {
		return nil, err
	}

	return quote, nil
}

func (this *SqliteQuoteRepository) Update(quote *domain.Quote) error {
	query := `UPDATE quote SET id_crypto = ?, price = ?, volume_24h = ?, percent_change_1h = ?, percent_change_24h = ?, percent_change_7d = ? WHERE id_quote = ?`

	if _, err := this.db.Exec(
		query,
		quote.IDCrypto,
		quote.Price,
		quote.Volume24h,
		quote.PercentChange1h,
		quote.PercentChange24h,
		quote.PercentChange7d,
		quote.ID,
	); err != nil {
		return err
	}

	return nil
}

func (this *SqliteQuoteRepository) UpdateByIDCrypto(quote *domain.Quote) error {
	query := `UPDATE quote SET price = ?, volume_24h = ?, percent_change_1h = ?, percent_change_24h = ?, percent_change_7d = ? WHERE id_crypto = ?`

	if _, err := this.db.Exec(
		query,
		quote.Price,
		quote.Volume24h,
		quote.PercentChange1h,
		quote.PercentChange24h,
		quote.PercentChange7d,
		quote.IDCrypto,
	); err != nil {
		return err
	}

	return nil
}

func (this *SqliteQuoteRepository) Delete(id int) error {
	query := `DELETE FROM quote WHERE id_quote = ?`

	if _, err := this.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}
