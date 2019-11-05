package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaDependenciaTipoDependencia_20191105_155154 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaDependenciaTipoDependencia_20191105_155154{}
	m.Created = "20191105_155154"

	migration.Register("CrearTablaDependenciaTipoDependencia_20191105_155154", m)
}

// Run the migrations
func (m *CrearTablaDependenciaTipoDependencia_20191105_155154) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS oikos.dependencia_tipo_dependencia ( id serial NOT NULL, tipo_dependencia_id integer NOT NULL, dependencia_id integer NOT NULL, activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_dependencia_tipo_dependencia PRIMARY KEY (id), CONSTRAINT fk_tipo_dependencia_dependencia_tipo_dependencia FOREIGN KEY (tipo_dependencia_id) REFERENCES oikos.tipo_dependencia(id), CONSTRAINT fk_dependencia_dependencia_tipo_dependencia FOREIGN KEY (dependencia_id) REFERENCES oikos.dependencia(id) );")
	m.SQL("ALTER TABLE oikos.dependencia_tipo_dependencia OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE oikos.dependencia_tipo_dependencia IS 'Tabla de rompimiento entre tipo_dependencia y dependencia.';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.id IS 'Identificador de la tabla.';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.tipo_dependencia_id IS 'Campo que contiene el identificador del tipo dependencia.';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.dependencia_id IS 'Campo para el id de la dependencia.';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CrearTablaDependenciaTipoDependencia_20191105_155154) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS oikos.dependencia_tipo_dependencia")
}
