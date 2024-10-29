package utils

import "fmt"

type EntityNotFound struct {
	ID   int
	Name string
}

func NewEntityNotFound(id int, name string) *EntityNotFound {
	return &EntityNotFound{
		ID:   id,
		Name: name,
	}
}

func (this *EntityNotFound) Error() string {
	return fmt.Sprintf("%s con ID %d no encontrado", this.Name, this.ID)
}
