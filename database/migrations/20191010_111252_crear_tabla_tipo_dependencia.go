package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoDependencia_20191010_111252 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoDependencia_20191010_111252{}
	m.Created = "20191010_111252"

	migration.Register("CrearTablaTipoDependencia_20191010_111252", m)
}

// Run the migrations
func (m *CrearTablaTipoDependencia_20191010_111252) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS oikos.tipo_dependencia ( id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_tipo_dependencia PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE oikos.tipo_dependencia OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE oikos.tipo_dependencia IS 'Tabla que contiene los distintos tipos de dependencia que hay en la Universidad Distrital.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_dependencia.id IS 'Identificador de la tabla.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_dependencia.nombre IS 'Campo que contiene el tipo de dependencia.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_dependencia.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de tipo_espacio_fisico.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_dependencia.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_dependencia.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_dependencia.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_dependencia.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CrearTablaTipoDependencia_20191010_111252) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS oikos.tipo_dependencia")
}
