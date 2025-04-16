package main

import "gowebsql/db"

func main() {
	db.Connect()
	db.Close()
}