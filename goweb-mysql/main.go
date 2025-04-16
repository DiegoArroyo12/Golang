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
	/* users := models.ListUser()
	fmt.Println(users)

	user := models.GetUser(2)
	fmt.Println(user)
	user.Username = "Arroyo"
	user.Password = "A123"
	user.Email = "a12@gmail.com"
	user.Save()
	fmt.Println(models.ListUser()) */

	user := models.GetUser(2)
	fmt.Println(user)

	user.Delete()
	fmt.Println(models.ListUser())

	db.TruncateTable("users")
	fmt.Println(models.ListUser())
	
	db.Close()
}