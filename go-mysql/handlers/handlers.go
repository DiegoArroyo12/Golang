package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

// ListContacts lista todos los contactos desde la base de datos
func ListContacts(db *sql.DB) {
	// Consulta SQL para seleccionar todos los contactos
	query := "SELECT * FROM contact"

	// Ejecutar la consulta
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// Iterar sobre los resultados y mostrarlos
	fmt.Println("\nLISTA DE CONTACTOS")
	fmt.Println("---------------------------------------------------------------------------------------------------")
	for rows.Next() {
		// Instancia de modelo contact
		contact := models.Contact{}

		// Control de nulos
		var valueEmail sql.NullString

		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		// Control de nulos
		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "Sin Correo Electrónico"
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Teléfono: %s\n",
				contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("---------------------------------------------------------------------------------------------------")
	}
}

// GetContactByID obtiene un contacto de la base de datos mediante su ID
func GetContactByID(db *sql.DB, contactID int) {
	// Consulta SQL para seleccionar un contacto por su ID
	query := "SELECT * FROM contact WHERE id = ?"

	row := db.QueryRow(query, contactID)

	// Instancia de modelo contact
	contact := models.Contact{}
	var valueEmail sql.NullString // Manejar valo null

	// Escanear el resultado en el modelo contact
	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No se encontró ningún contacto con el ID %d", contactID)
		}
	}

	// Control de nulos
	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "Sin Correo Electrónico"
	}

	fmt.Println("\nDatos del Contacto")
	fmt.Println("---------------------------------------------------------------------------------------------------")
	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Teléfono: %s\n",
				contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("---------------------------------------------------------------------------------------------------")
}

// CreateContact registra un nuevo contacto en la base de datos
func CreateContact(db *sql.DB, contact models.Contact) {
	// Sentencia SQL para insertar un nuevo contacto
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Nuevo contacto registrado con éxito")
}