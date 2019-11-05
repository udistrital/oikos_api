package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreatTablaEspacioFisicoCampo_20191105_111617 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreatTablaEspacioFisicoCampo_20191105_111617{}
	m.Created = "20191105_111617"

	migration.Register("CreatTablaEspacioFisicoCampo_20191105_111617", m)
}

// Run the migrations
func (m *CreatTablaEspacioFisicoCampo_20191105_111617) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE oikos.espacio_fisico_campo ( id serial NOT NULL, valor character varying(50) NOT NULL, espacio_fisico_id integer NOT NULL, campo_id integer NOT NULL, activo boolean NOT NULL, fecha_inicio TIMESTAMP NOT NULL, fecha_fin TIMESTAMP, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_espacio_fisico_campo PRIMARY KEY (id), CONSTRAINT fk_espacio_fisico_espacio_fisico_campo FOREIGN KEY (espacio_fisico_id) REFERENCES oikos.espacio_fisico(id), CONSTRAINT fk_campo_espacio_fisico_campo FOREIGN KEY (campo_id) REFERENCES oikos.campo(id), CONSTRAINT UQ_CAMPO UNIQUE (campo_id, espacio_fisico_id) );")
	m.SQL("ALTER TABLE oikos.espacio_fisico_campo OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE oikos.espacio_fisico_campo IS 'Tabla de rompimiento entre campo y espacio_fisico.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico_campo.id IS 'Identificador de la asignacion del espacio fisico.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico_campo.valor IS 'Valor del nuevo campo para el espacio fisico.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico_campo.espacio_fisico_id IS 'Identificador de la tabla espacio fisico.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico_campo.campo_id IS 'Identificador de la tabla campo';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico_campo.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico_campo.fecha_inicio IS 'Fecha de inicio de la asignacion.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico_campo.fecha_fin IS 'Fecha en la que finaliza la asignacion.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico_campo.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN oikos.espacio_fisico_campo.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CreatTablaEspacioFisicoCampo_20191105_111617) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS oikos.espacio_fisico_campo")
}
