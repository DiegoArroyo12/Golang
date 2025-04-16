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
	//user := models.CreateUser("arroyo", "a123", "arroyo@gmail.com")
	//fmt.Println(user)
	users := models.ListUser()
	fmt.Println(users)

	//db.TruncateTable("users")
	db.Close()
}