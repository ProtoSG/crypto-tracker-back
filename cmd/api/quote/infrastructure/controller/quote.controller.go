package controller

import (
	"net/http"

	"github.com/ProtoSG/crypto-tracker-back/cmd/api/quote/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/services"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type QuoteController struct {
	serviceContainer *services.ServiceContainer
}

func NewQuoteController(serviceContainer *services.ServiceContainer) *QuoteController {
	return &QuoteController{serviceContainer}
}

func (this *QuoteController) Create(w http.ResponseWriter, r *http.Request) {
	quote := domain.Quote{}
	if err := utils.ReadJSON(r, quote); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := this.serviceContainer.Quote.Create.Execute(quote.ID, &quote); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := map[string]string{
		"message": "Quote creado",
	}

	if err := utils.WriteJSON(w, http.StatusCreated, data); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (this *QuoteController) Read(w http.ResponseWriter, r *http.Request) {
	quotes, err := this.serviceContainer.Quote.Read.Execute()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, quotes); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (this *QuoteController) ReadByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetID(w, r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	quote, err := this.serviceContainer.Quote.ReadByID.Execute(id)
	if err != nil {
		if _, ok := err.(*utils.EntityNotFound); ok {
			utils.WriteError(w, http.StatusNotFound, err.Error())
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, quote); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (this *QuoteController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetID(w, r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	quote := domain.Quote{}
	if err := utils.ReadJSON(r, quote); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := this.serviceContainer.Quote.Update.Execute(id, &quote); err != nil {
		if _, ok := err.(*utils.EntityNotFound); ok {
			utils.WriteError(w, http.StatusNotFound, err.Error())
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	data := map[string]string{
		"message": "Quote actualizado",
	}
	if err := utils.WriteJSON(w, http.StatusOK, data); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (this *QuoteController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetID(w, r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := this.serviceContainer.Quote.Delete.Execute(id); err != nil {
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
