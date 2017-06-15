package utils

import (
	"fmt"

	"github.com/asdine/storm"
)

func Query(fn func(*storm.DB)) {
	db, err := storm.Open("wallet.db")
	if err != nil {
		fmt.Println("error al abrir la base de datos")
	} /* else {
		fmt.Println("Conexion exitosa")
	}*/
	fn(db)
	db.Close()
}
