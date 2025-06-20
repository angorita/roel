package handlers

import (
	"database/sql"
	"fmt"
	"log"

	// "roel/models"
	"github.com/angorita/roel/models"
)

func ListaMateriales(db *sql.DB) {
	//consulta sql
	query := `select * from producto limit 2`
	//ejecutar consulta
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	//iterar
	fmt.Println("\nLista de Materiales")
	Ventana()
	for rows.Next() {
		//instancia de materiales
		materiales := models.Materiales{}
		err := rows.Scan(&materiales.Descripcion, &materiales.Precio, &materiales.Cantidad,
			&materiales.Fecha, &materiales.Dolar, &materiales.Id, &materiales.Bhabilitado)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(`Descripcion: %s, Precio: %f,Cantidad: %d,Fecha: %s,Dolar: %f,Id: %d, Bhabilitado %t\n`, materiales.Descripcion, materiales.Precio, materiales.Cantidad, materiales.Fecha, materiales.Dolar, materiales.Id, materiales.Bhabilitado)

		Ventana()
	}
}
func Ventana() {
	fmt.Println("\n------------------------------------------------------------------------------------------------------------------------------------------------------------")

}
