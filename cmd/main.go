package main

import (
	"log"

	"github.com/ProtoSG/crypto-tracker-back/cmd/api"
	"github.com/ProtoSG/crypto-tracker-back/internal/db"
)

func main() {
	startDB := *db.NewDB("crypto.db")
	db := startDB.SetupDB()

	app := api.NewApi(":8080", db)
	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
