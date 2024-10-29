package controller

import (
	"net/http"

	"github.com/ProtoSG/crypto-tracker-back/cmd/api/crypto/domain"
	"github.com/ProtoSG/crypto-tracker-back/internal/services"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
)

type CryptoController struct {
	serviceContainer *services.ServiceContainer
}

func NewCryptoController(serviceContainer *services.ServiceContainer) *CryptoController {
	return &CryptoController{serviceContainer}
}

func (this *CryptoController) Create(w http.ResponseWriter, r *http.Request) {
	crypto := domain.Crypto{}
	if err := utils.ReadJSON(r, crypto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := this.serviceContainer.Crypto.Create.Execute(
		crypto.ID,
		crypto.Name,
		crypto.Symbol,
		crypto.Slug,
		crypto.CirculatingSupply,
		crypto.CmcRank,
	); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := map[string]string{
		"message": "Crypto creado",
	}

	if err := utils.WriteJSON(w, http.StatusCreated, data); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (this *CryptoController) Read(w http.ResponseWriter, r *http.Request) {
	cryptos, err := this.serviceContainer.Crypto.Read.Execute()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, cryptos); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (this *CryptoController) ReadByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetID(w, r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	crypto, err := this.serviceContainer.Crypto.ReadByID.Execute(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, crypto); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (this *CryptoController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetID(w, r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	crypto := domain.Crypto{}
	if err := utils.ReadJSON(r, crypto); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := this.serviceContainer.Crypto.Update.Execute(id, &crypto); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := map[string]string{
		"message": "Crypto actualizado",
	}
	if err := utils.WriteJSON(w, http.StatusOK, data); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (this *CryptoController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetID(w, r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := this.serviceContainer.Crypto.Delete.Execute(id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := utils.WriteJSON(nil, http.StatusNoContent, nil); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
