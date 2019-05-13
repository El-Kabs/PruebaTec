package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func getRecetas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != "GET" {
		http.Error(w, "Metodo no valido.", http.StatusMethodNotAllowed)
		return
	}

	if len(r.URL.Query()) == 0 {
		registros, err := db.Query("SELECT * FROM recetas")

		if err != nil {
			http.Error(w, err.Error(), 500)
			fmt.Fprint(w, "Internal Server Error.", http.StatusInternalServerError)
			return
		}

		var recetas []Receta
		var receta Receta

		for registros.Next() {

			err = registros.Scan(&receta.Nombre, &receta.Descripcion)

			recetas = append(recetas, receta)

		}

		jsonRetorno, _ := json.Marshal(recetas)

		if string(jsonRetorno) == "null" {
			fmt.Fprint(w, "No hay recetas en la base de datos.")
		} else {
			fmt.Fprint(w, string(jsonRetorno))
		}

		defer registros.Close()

	} else {
		recetaBuscar := r.URL.Query().Get("nombre")
		registrosB, err := db.Query("SELECT * FROM recetas WHERE lower(nombre) = lower($1)", recetaBuscar)

		if err != nil {
			http.Error(w, err.Error(), 500)
			fmt.Fprint(w, "Internal Server Error.", http.StatusInternalServerError)
			return
		}

		var recetas []Receta
		var receta Receta

		for registrosB.Next() {

			err = registrosB.Scan(&receta.Nombre, &receta.Descripcion)

			recetas = append(recetas, receta)

		}

		jsonRetorno, _ := json.Marshal(recetas)

		if string(jsonRetorno) == "null" {
			fmt.Fprint(w, "No hay recetas con el nombre dado en la base de datos.")
		} else {
			fmt.Fprint(w, string(jsonRetorno))
		}

		defer registrosB.Close()

	}
}

func createRecetas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	setupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Metodo no valido.", http.StatusMethodNotAllowed)
		return
	}

	cuerpo, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := Receta{}

	if err := json.Unmarshal([]byte(cuerpo), &data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if _, err := db.Exec("INSERT INTO recetas (nombre, descripcion) VALUES ($1, $2)", data.Nombre, data.Descripcion); err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Fprint(w, "Internal Server Error.", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Fila guardada correctamente.")
}

func updateRecetas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	setupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "PUT" {
		http.Error(w, "Metodo no valido.", http.StatusMethodNotAllowed)
		return
	}

	cuerpo, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := RecetaActualizar{}

	if err := json.Unmarshal([]byte(cuerpo), &data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if _, err := db.Exec("UPDATE recetas SET nombre=$1, descripcion=$2 WHERE nombre=$3 ", data.Nombre, data.Descripcion, data.NombreActualizar); err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Fprint(w, "Internal Server Error.", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Fila actualizada correctamente.")
}

func deleteRecetas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	setupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "DELETE" {
		http.Error(w, "Metodo no valido.", http.StatusMethodNotAllowed)
		return
	}

	cuerpo, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	data := RecetaBorrar{}

	if err := json.Unmarshal([]byte(cuerpo), &data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if _, err := db.Exec("DELETE FROM recetas WHERE nombre=$1", data.Nombre); err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Fprint(w, "Internal Server Error.", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Fila eliminada correctamente.")
}
