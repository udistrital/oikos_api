package models

//Estructura para construir el arbol de dependencia
type Tree struct {
	Id       int
	Nombre   string
	Opciones *[]Tree
}

type DependenciaPadreHijo struct {
	Id       int    `orm:"column(id)"`
	Nombre   string `orm:"column(nombre)"`
	Padre    int    `orm:"column(padre_id)"`
	Hija     int    `orm:"column(hija_id)"`
	Opciones []DependenciaPadreHijo
}

type EspacioFisicoPadreHijo struct {
	Id       int    `orm:"column(id)"`
	Nombre   string `orm:"column(nombre)"`
	Padre    int    `orm:"column(padre_id)"`
	Hijo     int    `orm:"column(hijo_id)"`
	Opciones []EspacioFisicoPadreHijo
}
