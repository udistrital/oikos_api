package models
//Estructura para construir el arbol de dependencia
type Tree struct {
	Id       int
	Nombre   string
	Opciones *[]Tree
}

type TreePadre struct {
	Id       int
	Nombre   string
	Padre    int
	Opciones *[]TreePadre
}

type TreeDos struct {
	Id       int
	Nombre   string
	Padre    int
	Hijo     int
	Opciones *[]TreeDos
	
}

type Key struct {
    X, Y int
}