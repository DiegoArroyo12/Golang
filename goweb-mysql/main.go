package main

import (
	"fmt"
	"gowebsql/db"
	"gowebsql/models"
)

func main() {
	db.Connect()
	
	fmt.Println(db.ExistsTable("users"))
	db.CreateTable(models.UserSchema, "users")

	db.Close()
}