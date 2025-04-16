package main

import (
	"gowebsql/db"
	"gowebsql/models"
)

func main() {
	db.Connect()
	
	db.CreateTable(models.UserSchema)

	db.Close()
}