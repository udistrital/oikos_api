package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoEspacioFisico_20191010_105308 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoEspacioFisico_20191010_105308{}
	m.Created = "20191010_105308"

	migration.Register("CrearTablaTipoEspacioFisico_20191010_105308", m)
}

// Run the migrations
func (m *CrearTablaTipoEspacioFisico_20191010_105308) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS oikos.tipo_espacio_fisico ( id serial NOT NULL, nombre character varying(100) NOT NULL, -- descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_tipo_espacio_fisico PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE oikos.tipo_espacio_fisico OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE oikos.tipo_espacio_fisico IS 'Tabla que contiene los tipos de espacios fisicos existentes.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_espacio_fisico.id IS 'Identificador de cada tipo de espacio fisico que pertenece a la Universidad Distr';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_espacio_fisico.nombre IS 'Nombre del tipo de espacio fisico perteneciente a la Universidad Distrital';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_espacio_fisico.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de tipo_espacio_fisico.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_espacio_fisico.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_espacio_fisico.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_espacio_fisico.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_espacio_fisico.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CrearTablaTipoEspacioFisico_20191010_105308) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS oikos.tipo_espacio_fisico")
}
