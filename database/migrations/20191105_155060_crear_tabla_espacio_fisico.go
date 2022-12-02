package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaEspacioFisico_20191105_155060 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaEspacioFisico_20191105_155060{}
	m.Created = "20191105_155059"

	migration.Register("CrearTablaEspacioFisico_20191105_155060", m)
}

// Run the migrations
func (m *CrearTablaEspacioFisico_20191105_155060) Up() {
	m.SQL("CREATE TABLE IF NOT EXISTS oikos.espacio_fisico ( id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, tipo_espacio_fisico_id integer NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_espacio_fisico PRIMARY KEY (id), CONSTRAINT fk_tipo_espacio_fisico_espacio_fisico FOREIGN KEY (tipo_espacio_fisico_id) REFERENCES oikos.tipo_espacio_fisico(id) );")
	m.SQL("COMMENT ON TABLE oikos.espacio_fisico IS 'Tabla de Rompimiento que reune los atributos de un espacio fisico';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico.id IS 'Identificador del espacio fisico especifico de la Universidad Distrital.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico.nombre IS 'Nombre perteneciente al espacio físico';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de espacio_fisico.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico.tipo_espacio_fisico_id IS 'Llave foranea que contiene el identificador del tipo de espacio fisico de la entidad tipo_espacio_fisico';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")

}

// Reverse the migrations
func (m *CrearTablaEspacioFisico_20191105_155060) Down() {
	m.SQL("DROP TABLE IF EXISTS oikos.espacio_fisico")
}
