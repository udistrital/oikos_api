package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaDependenciaPadre_20191105_155202 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaDependenciaPadre_20191105_155202{}
	m.Created = "20191105_155202"

	migration.Register("CrearTablaDependenciaPadre_20191105_155202", m)
}

// Run the migrations
func (m *CrearTablaDependenciaPadre_20191105_155202) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS oikos.dependencia_padre ( id serial NOT NULL, padre_id integer NOT NULL, hija_id integer NOT NULL, activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_dependencia_padre PRIMARY KEY (id), CONSTRAINT uq_hija_id_dependencia_padre UNIQUE (hija_id), CONSTRAINT uq_hija_id_padre_id_dependencia_padre UNIQUE (padre_id, hija_id), CONSTRAINT fk_dependencia_dependencia_padre FOREIGN KEY (padre_id) REFERENCES oikos.dependencia(id), CONSTRAINT fk_dependencia_dependencia_hija FOREIGN KEY (hija_id) REFERENCES oikos.dependencia(id) );")
	m.SQL("ALTER TABLE oikos.dependencia_padre OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE oikos.dependencia_padre IS 'Tabla que contiene las dependencias de la Universidad Distrital';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia_padre.id IS 'Identificador de la dependencia.';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia_padre.padre_id IS 'Id dependencia padre';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia_padre.hija_id IS 'Identificador dependencia hija (si la hay)';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia_padre.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia_padre.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN oikos.dependencia_padre.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CrearTablaDependenciaPadre_20191105_155202) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS oikos.dependencia_padre")
}
