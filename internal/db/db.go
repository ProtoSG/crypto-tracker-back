package db

import (
	"database/sql"
	"log"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type DB struct {
	addr string
}

func NewDB(addr string) *DB {
	return &DB{addr}
}

func (this *DB) initDB() (*sql.DB, error) {
	db, err := sql.Open("libsql", this.addr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (this *DB) SetupDB() *sql.DB {
	db, err := this.initDB()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error al hacer ping a la base de datos: %v", err)
	}

	log.Println("Base de datos conectada exitosamente")
	return db
}
