package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaAsignacionEspacioFisicoDependencia_20191010_112624 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaAsignacionEspacioFisicoDependencia_20191010_112624{}
	m.Created = "20191010_112624"

	migration.Register("CrearTablaAsignacionEspacioFisicoDependencia_20191010_112624", m)
}

// Run the migrations
func (m *CrearTablaAsignacionEspacioFisicoDependencia_20191010_112624) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS oikos.asignacion_espacio_fisico_dependencia ( id serial NOT NULL, espacio_fisico_id integer NOT NULL, dependencia_id integer NOT NULL, activo boolean NOT NULL, fecha_inicio TIMESTAMP NOT NULL, fecha_fin TIMESTAMP, documento_soporte integer, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_asignacion_espacio_fisico_dependencia PRIMARY KEY (id), CONSTRAINT fk_espacio_fisico_asignacion_espacio_fisico_dependencia FOREIGN KEY (espacio_fisico_id) REFERENCES oikos.espacio_fisico(id), CONSTRAINT fk_dependencia_asignacion_espacio_fisico_dependencia FOREIGN KEY (dependencia_id) REFERENCES oikos.dependencia(id) );")
	m.SQL("ALTER TABLE oikos.asignacion_espacio_fisico_dependencia OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE oikos.asignacion_espacio_fisico_dependencia IS 'Tabla de rompimiento que reune los atributos necesarios para la asignacion de un espacio fisico.';")
	m.SQL("COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.id IS 'Identificador de la asignacion del espacio fisico';")
	m.SQL("COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.espacio_fisico_id IS 'Campo que contiene el id, de la llave foranea';")
	m.SQL("COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.dependencia_id IS 'Identificador que contiene el id de dependencia';")
	m.SQL("COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.fecha_inicio IS 'Fecha de inicio de la asignacion.';")
	m.SQL("COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.fecha_fin IS 'Fecha en la que finaliza la asignacion.';")
	m.SQL("COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.documento_soporte IS 'Documento que soporta la asignacion del espacio fisico.';")
	m.SQL("COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CrearTablaAsignacionEspacioFisicoDependencia_20191010_112624) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS oikos.asignacion_espacio_fisico_dependencia")
}
