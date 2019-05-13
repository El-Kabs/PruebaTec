package main

type Receta struct {
	Nombre      string `json:"Nombre"`
	Descripcion string `json:"Descripcion"`
}

type RecetaActualizar struct {
	Nombre           string `json:"Nombre"`
	Descripcion      string `json:"Descripcion"`
	NombreActualizar string `json:"NombreActualizar"`
}

type RecetaBorrar struct {
	Nombre string `json:"Nombre"`
}
