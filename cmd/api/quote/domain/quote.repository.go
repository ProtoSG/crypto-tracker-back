package domain

type QuoteRepository interface {
	Create(*Quote) error
	Read() ([]*Quote, error)
	ReadByID(id int) (*Quote, error)
	Update(quote *Quote) error
	Delete(id int) error
}
