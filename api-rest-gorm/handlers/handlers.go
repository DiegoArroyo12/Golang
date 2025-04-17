package handlers

import (
	"encoding/json"
	"gorm/db"
	"gorm/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	users := getUserById(r)

	db.Database.Find(&users)
	sendData(rw, users, http.StatusOK)
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	sendData(rw, user, http.StatusOK)
}

func getUserById(r *http.Request) (models.User) {
	// Obtener ID
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	user := models.User{}
	db.Database.First(&user, userId)

	return user
} 

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	// Obtener Registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if  err := decoder.Decode(&user); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Save(&user)
		sendData(rw, user, http.StatusCreated)
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	// Obtener Registro
	var userId int64

	user_ant := getUserById(r)
	userId = user_ant.Id

	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if  err := decoder.Decode(&user); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		user.Id = userId
		db.Database.Save(&user)
		sendData(rw, user, http.StatusCreated)
	}
}

/* 
func DeletUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		user.Delete()
		models.SendData(rw, user)
	}
}
*/