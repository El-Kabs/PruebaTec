package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func main() {

	db, err = sql.Open("postgres", "postgresql://root@localhost:26257/recetas?sslmode=disable")

	if err != nil {
		log.Fatal("Error conectandose a la base de datos: ", err)
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS recetas (id INT PRIMARY KEY, nombre STRING NOT NULL, descripcion STRING NOT NULL)"); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	http.HandleFunc("/recetas", getRecetas)
	http.HandleFunc("/createReceta", createRecetas)
	http.HandleFunc("/updateReceta", updateRecetas)
	http.HandleFunc("/deleteReceta", deleteRecetas)

	err = http.ListenAndServe(":9001", nil)
	if err != nil {
		panic(err)
	}
}
