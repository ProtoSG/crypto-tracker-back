package utils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetID(w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return 0, fmt.Errorf("Formato del ID inválido: %w", err)
	}

	return id, nil
}
