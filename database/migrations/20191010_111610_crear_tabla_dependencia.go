package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaDependencia_20191010_111610 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaDependencia_20191010_111610{}
	m.Created = "20191010_111610"

	migration.Register("CrearTablaDependencia_20191010_111610", m)
}

// Run the migrations
func (m *CrearTablaDependencia_20191010_111610) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS oikos.dependencia ( id serial NOT NULL, nombre character varying(100) NOT NULL, telefono_dependencia character varying(500) NOT NULL, correo_electronico character varying(100), activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_dependencia PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE oikos.dependencia OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE oikos.dependencia  IS 'Tabla que contiene las dependencias de la Universidad Distrital.';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia.id IS 'Identificador de la dependencia.';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia.nombre IS 'Nombre de la dependencia perteneciente a la Universidad Distrital.';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia.telefono_dependencia IS 'Indica el numero de telefono de la dependencia.';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia.correo_electronico IS 'Correo electrónico asociado a la dependencia.';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CrearTablaDependencia_20191010_111610) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS oikos.dependencia")
}
