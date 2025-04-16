package main

import (
	"fmt"
	"gowebsql/db"
	"gowebsql/models"
)

func main() {
	db.Connect()
	
	fmt.Println(db.ExistsTable("users"))
	//db.CreateTable(models.UserSchema, "users")
	user := models.CreateUser("diego", "d123", "diego@gmail.com")
	fmt.Println(user)

	//db.TruncateTable("users")
	db.Close()
}