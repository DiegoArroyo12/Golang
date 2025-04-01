package main

import (
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
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

	// Crear una instancia de Contact
	newContact := models.Contact{
		Id: 4,
		Name: "Modificado",
		Email: "modificado@example.com",
		Phone: "987654321",
	}

	handlers.ListContacts(db)

	handlers.GetContactByID(db, 2)

	// Crear Nuevo Registro
	/* handlers.CreateContact(db, newContact)
	handlers.ListContacts(db) */

	handlers.UpdateContact(db, newContact)
	handlers.ListContacts(db)
}