package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ProtoSG/crypto-tracker-back/internal/routes"
	"github.com/ProtoSG/crypto-tracker-back/internal/services"
	"github.com/ProtoSG/crypto-tracker-back/internal/utils"
	"github.com/gorilla/handlers"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewApi(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (this *APIServer) Run() error {
	serviceContainer := services.NewServiceContainer(this.db)

	router := routes.NewRoutes(serviceContainer)
	router.HandleFunc("/", this.checkHandler)

	log.Printf("Listening on http://localhost%s", this.addr)

	svr := &http.Server{
		Addr:    this.addr,
		Handler: handlers.CORS()(router),
	}

	if err := svr.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (this *APIServer) checkHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"message": "Hello World",
	}

	if err := utils.WriteJSON(w, http.StatusOK, data); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
	}
}
