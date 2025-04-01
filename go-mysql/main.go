package main

import (
	"go-mysql/database"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Establecer conexión a la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}