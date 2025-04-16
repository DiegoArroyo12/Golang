package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// username:password@tcp(localhost:3306)/database
const url = "username:password@tcp(localhost:3306)/database"

var db *sql.DB

func Connect() {
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexión Exitosa")
	db = conection
}

func Close() {
	db.Close()
}