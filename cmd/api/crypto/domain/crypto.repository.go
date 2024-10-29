package domain

type CryptoRepository interface {
	Create(*Crypto) error
	Read() ([]*Crypto, error)
	ReadByID(id int) (*Crypto, error)
	Update(crypto *Crypto) error
	Delete(id int) error
}
