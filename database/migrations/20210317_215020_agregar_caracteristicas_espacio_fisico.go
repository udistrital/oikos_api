package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarCaracteristicasEspacioFisico_20210317_215020 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarCaracteristicasEspacioFisico_20210317_215020{}
	m.Created = "20210317_215020"

	migration.Register("AgregarCaracteristicasEspacioFisico_20210317_215020", m)
}

// Run the migrations
func (m *AgregarCaracteristicasEspacioFisico_20210317_215020) Up() {
	// to make schema update
	file, err := ioutil.ReadFile("../scripts/20210317_215020_agregar_caracteristicas_espacio_fisico_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}

}

// Reverse the migrations
func (m *AgregarCaracteristicasEspacioFisico_20210317_215020) Down() {
	// to make schema update
	file, err := ioutil.ReadFile("../scripts/20210317_215020_agregar_caracteristicas_espacio_fisico_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
	}

}
