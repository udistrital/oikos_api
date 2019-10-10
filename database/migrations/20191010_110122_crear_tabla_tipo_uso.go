package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoUso_20191010_110122 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoUso_20191010_110122{}
	m.Created = "20191010_110122"

	migration.Register("CrearTablaTipoUso_20191010_110122", m)
}

// Run the migrations
func (m *CrearTablaTipoUso_20191010_110122) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS oikos.tipo_uso ( id serial NOT NULL, nombre character varying(100) NOT NULL, -- descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_tipo_uso PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE oikos.tipo_uso OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE oikos.tipo_uso IS 'Tabla quie contiene los diversos usos que puede tener un espacio fisico.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_uso.id IS 'Identificador del tipo de uso de espacios fisicos.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_uso.nombre IS 'Nombre del uso que se le va a dar al espacio fisico.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_uso.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de tipo_uso.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_uso.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_uso.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_uso.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_uso.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CrearTablaTipoUso_20191010_110122) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS oikos.tipo_uso")
}
