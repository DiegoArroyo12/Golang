package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Realiza la conexión a DB
var dsn = "root:rootmac24@tcp(localhost:3306)/goweb_db?charset=utf8mb4&parseTime=True&loc=America%2FMexico_City"
var Database = func () (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la conexión", err)
		panic(err)
	} else {
		fmt.Println("Conexión Exitosa")
		return db
	}
} ()

/* // username:password@tcp(localhost:3306)/database
const url = "root:rootmac24@tcp(127.0.0.1:3306)/goweb_db"

// Guarda la conexión
var db *sql.DB


func Connect() {
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexión Exitosa")
	db = conection
}

// Cierra la conexión
func Close() {
	db.Close()
}

// Verificar la conexión
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

// Verificar si una tabla existe o no
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows.Next()
}

// Crea una tabla
func CreateTable(schema string, name string) {
	if !ExistsTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Reiniciar el registro de una tabla
func TruncateTable(tablename string) {
	sql := fmt.Sprintf("TRUNCATE %s", tablename)
	Exec(sql)
}

// Polimorfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}

	return result, err
}

// Polimorfismo de Exec
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}

	return rows,err
} */