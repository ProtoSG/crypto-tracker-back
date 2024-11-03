package infrastructure

import (
	"database/sql"

	"github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/domain"
)

type SqliteHistorialPriceRepository struct {
	db *sql.DB
}

func NewSqliteHistorialPriceRepository(db *sql.DB) *SqliteHistorialPriceRepository {
	return &SqliteHistorialPriceRepository{db}
}

func (this *SqliteHistorialPriceRepository) Create(historialPrice *domain.HistorialPrice) error {
	query := `INSERT INTO historial_price (id_crypto, price, date) 
    VALUES (?, ?, ?)
  `
	_, err := this.db.Exec(
		query,
		historialPrice.IDCrypto,
		historialPrice.Price,
		historialPrice.Date,
	)
	if err != nil {
		return err
	}

	return nil
}

func (this *SqliteHistorialPriceRepository) Read() ([]*domain.HistorialPrice, error) {
	query := `SELECT * FROM historial_price`

	rows, err := this.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var historialPrices []*domain.HistorialPrice
	for rows.Next() {
		historialPrice := &domain.HistorialPrice{}
		if err := rows.Scan(
			&historialPrice.ID,
			&historialPrice.IDCrypto,
			&historialPrice.Price,
			&historialPrice.Date,
		); err != nil {
			return nil, err
		}

		historialPrices = append(historialPrices, historialPrice)
	}

	return historialPrices, nil
}

func (this *SqliteHistorialPriceRepository) ReadByID(id int) (*domain.HistorialPrice, error) {
	query := `SELECT * FROM historial_price WHERE id_historial_price = ?`

	historialPrice := &domain.HistorialPrice{}

	if err := this.db.QueryRow(
		query,
		id,
	).Scan(
		&historialPrice.ID,
		&historialPrice.IDCrypto,
		&historialPrice.Price,
		&historialPrice.Date,
	); err != nil {
		return nil, err
	}

	return historialPrice, nil
}

func (this *SqliteHistorialPriceRepository) ReadByIDCrypto(id int) ([]*domain.HistorialPrice, error) {
	query := `SELECT * FROM historial_price WHERE id_crypto = ?`

	rows, err := this.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var historialPrices []*domain.HistorialPrice
	for rows.Next() {
		historialPrice := &domain.HistorialPrice{}
		if err := rows.Scan(
			&historialPrice.ID,
			&historialPrice.IDCrypto,
			&historialPrice.Price,
			&historialPrice.Date,
		); err != nil {
			return nil, err
		}

		historialPrices = append(historialPrices, historialPrice)
	}

	return historialPrices, nil
}

func (this *SqliteHistorialPriceRepository) Update(historialPrice *domain.HistorialPrice) error {
	query := `UPDATE historial_price SET id_crypto = ?, price = ?, date = ? WHERE id_historial_price = ?`

	if _, err := this.db.Exec(
		query,
		historialPrice.IDCrypto,
		historialPrice.Price,
		historialPrice.Date,
		historialPrice.ID,
	); err != nil {
		return err
	}

	return nil
}

func (this *SqliteHistorialPriceRepository) UpdateByIDCrypto(historialPrice *domain.HistorialPrice) error {
	query := `UPDATE historial_price SET  price = ?, date = ? WHERE id_crypto = ?`

	if _, err := this.db.Exec(
		query,
		historialPrice.Price,
		historialPrice.Date,
		historialPrice.ID,
		historialPrice.IDCrypto,
	); err != nil {
		return err
	}

	return nil
}

func (this *SqliteHistorialPriceRepository) Delete(id int) error {
	query := `DELETE FROM historialPrice WHERE id_historial_price = ?`

	if _, err := this.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}
