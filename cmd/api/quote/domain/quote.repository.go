package domain

type QuoteRepository interface {
	Create(quote *Quote) error
	Read() ([]*Quote, error)
	ReadByID(id int) (*Quote, error)
	ReadByIDCrypto(id int) (*Quote, error)
	Update(quote *Quote) error
	UpdateByIDCrypto(quote *Quote) error
	Delete(id int) error
}
