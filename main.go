package main

import (
	"log"
	"roel/db"
)

func main() {
	db, err := db.Conexion()
	if err != nil {
		log.Fatal(err)
		
	}
	defer db.Close()
}
