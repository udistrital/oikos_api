package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaCampo_20191105_155230 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaCampo_20191105_155230{}
	m.Created = "20191105_155230"

	migration.Register("CrearTablaCampo_20191105_155230", m)
}

// Run the migrations
func (m *CrearTablaCampo_20191105_155230) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS oikos.campo ( id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_campo PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE oikos.campo OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE oikos.campo IS 'Tabla de los campos o atributos que se necesiten en un espacio fisico.';")
	m.SQL("COMMENT ON COLUMN oikos.campo.id IS 'Identificador de la tabla.';")
	m.SQL("COMMENT ON COLUMN oikos.campo.nombre IS 'Nombre del nuevo campo que se requiere para espacios fisicos.';")
	m.SQL("COMMENT ON COLUMN oikos.campo.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion del nuevo campo.';")
	m.SQL("COMMENT ON COLUMN oikos.campo.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN oikos.campo.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN oikos.campo.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN oikos.campo.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CrearTablaCampo_20191105_155230) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS oikos.campo")
}
