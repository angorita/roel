package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"roel/models"
)

func ListaMateriales(db *sql.DB) {
	//consulta sql
	query := `select * from producto`
	//ejecutar consulta
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	//iterar
	fmt.Println("\nLista de Materiales")
	fmt.Println("*******************************************************")
	for rows.Next(){
		//instancia de materiales
		materiales:=models.Materiales{}
		rows.Scan(&materiales.Descripcion,&materiales.Precio,&materiales.Cantidad,
		&materiales.Fecha,&materiales.Dolar,&materiales.Id,&materiales.Bhabilitado)
		if err !=nil{
			log.Fatal(err)
		}
		fmt.Print(materiales)
	}
}
