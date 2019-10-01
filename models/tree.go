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