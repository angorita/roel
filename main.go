package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dns := "oscar:emi@tcp(localhost:3306)/loft"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Conexion exitosa")
}
