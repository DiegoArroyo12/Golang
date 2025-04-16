package main

import (
	"fmt"
	"gowebsql/db"
)

func main() {
	db.Connect()
	
	fmt.Println(db.ExistsTable("users"))
	//db.CreateTable(models.UserSchema, "users")
	db.TruncateTable("users")

	db.Close()
}