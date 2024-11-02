package controller

import (
	"net/http"

	"github.com/ProtoSG/crypto-tracker-back/cmd/api/historialPrice/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/services"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type HistorialPriceController struct {
	serviceContainer *services.ServiceContainer
}

func NewHistorialPriceController(serviceContainer *services.ServiceContainer) *HistorialPriceController {
	return &HistorialPriceController{serviceContainer}
}

func (this *HistorialPriceController) Create(w http.ResponseWriter, r *http.Request) {
	historialPrice := &domain.HistorialPrice{}
	if err := utils.ReadJSON(r, historialPrice); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := this.serviceContainer.HistorialPrice.Create.Execute(
		historialPrice.ID,
		historialPrice.IDCrypto,
		historialPrice.Price,
		historialPrice.Date,
	); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := map[string]string{
		"message": "HistorialPrice creado",
	}

	if err := utils.WriteJSON(w, http.StatusCreated, data); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (this *HistorialPriceController) Read(w http.ResponseWriter, r *http.Request) {
	historialPrices, err := this.serviceContainer.HistorialPrice.Read.Execute()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, historialPrices); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (this *HistorialPriceController) ReadByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetID(w, r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	historialPrice, err := this.serviceContainer.HistorialPrice.ReadByID.Execute(id)
	if err != nil {
		if _, ok := err.(*utils.EntityNotFound); ok {
			utils.WriteError(w, http.StatusNotFound, err.Error())
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, historialPrice); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (this *HistorialPriceController) ReadByIDCrypto(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetID(w, r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	historialPrice, err := this.serviceContainer.HistorialPrice.ReadByIDCrypto.Execute(id)
	if err != nil {
		if _, ok := err.(*utils.EntityNotFound); ok {
			utils.WriteError(w, http.StatusNotFound, err.Error())
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, historialPrice); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (this *HistorialPriceController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetID(w, r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	historialPrice := &domain.HistorialPrice{}
	if err := utils.ReadJSON(r, historialPrice); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := this.serviceContainer.HistorialPrice.Update.Execute(id, historialPrice); err != nil {
		if _, ok := err.(*utils.EntityNotFound); ok {
			utils.WriteError(w, http.StatusNotFound, err.Error())
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	data := map[string]string{
		"message": "HistorialPrice actualizado",
	}
	if err := utils.WriteJSON(w, http.StatusOK, data); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (this *HistorialPriceController) UpdateByIDCrypto(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetID(w, r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	historialPrice := &domain.HistorialPrice{}
	if err := utils.ReadJSON(r, historialPrice); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := this.serviceContainer.HistorialPrice.UpdateByIDCrypto.Execute(id, historialPrice); err != nil {
		if _, ok := err.(*utils.EntityNotFound); ok {
			utils.WriteError(w, http.StatusNotFound, err.Error())
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	data := map[string]string{
		"message": "HistorialPrice actualizado",
	}
	if err := utils.WriteJSON(w, http.StatusOK, data); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (this *HistorialPriceController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetID(w, r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := this.serviceContainer.HistorialPrice.Delete.Execute(id); err != nil {
		if _, ok := err.(*utils.EntityNotFound); ok {
			utils.WriteError(w, http.StatusNotFound, err.Error())
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := utils.WriteJSON(nil, http.StatusNoContent, nil); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
