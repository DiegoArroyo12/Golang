package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dns := "root:rootmac24@tcp(localhost:3306)/db_contacts"

	// Abrir una conexión a la base de datos
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Conexión Exitosa a la Base de Datos")
}