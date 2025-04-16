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

	user := models.GetUser(2)
	fmt.Println(user)

	//db.TruncateTable("users")
	db.Close()
}