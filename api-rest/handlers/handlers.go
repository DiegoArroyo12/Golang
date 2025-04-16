package handlers

import (
	"apirest/db"
	"apirest/models"

	//"encoding/json"
	//"encoding/xml"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2" // go get gopkg.in/yaml.v2
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	//rw.Header().Set("Content-Type", "application/json")
	//rw.Header().Set("Content-Type", "text/xml")

	db.Connect()
	users := models.ListUser()
	db.Close()
	
	output, _ := yaml.Marshal(users)
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Obtiene un usuario")
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Crea un usuario")
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Actualiza un usuario")
}

func DeletUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Elimina un usuario")
}