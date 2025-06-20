package main

import (
	"log"
	"roel/db"
	"roel/handlers"
)

func main() {
	//clase 113
	db, err := db.Conexion()
	if err != nil {
		log.Fatal(err)

	}

	defer db.Close()
	handlers.ListaMateriales(db)
}
