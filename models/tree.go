package models
//Estructura para construir el arbol de dependencia
type Tree struct {
	Id       int
	Nombre   string
	Opciones *[]Tree
}

type DependenciaPadreHijo struct {
	Id       int
	Nombre   string
	Padre    int
	Hija     int
	Opciones []DependenciaPadreHijo
}