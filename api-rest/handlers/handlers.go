package handlers

import (
	"apirest/db"
	"apirest/models"
	"strconv"

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	db.Connect()
	users := models.ListUser()
	db.Close()
	
	output, _ := json.Marshal(users)
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	// Obtener id
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	user := models.GetUser(userId)
	db.Close()
	
	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
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