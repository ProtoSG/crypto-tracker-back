package domain

type CryptoRepository interface {
	Create(*Crypto) error
	Read() ([]*Crypto, error)
	ReadByID(id int) (*Crypto, error)
	ReadByName(name string) (*Crypto, error)
	Update(crypto *Crypto) error
	UpdateByName(crypto *Crypto) error
	Delete(id int) error
}
