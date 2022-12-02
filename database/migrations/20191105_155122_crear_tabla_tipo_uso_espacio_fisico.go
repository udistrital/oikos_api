package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoUsoEspacioFisico_20191105_155122 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoUsoEspacioFisico_20191105_155122{}
	m.Created = "20191105_155122"

	migration.Register("CrearTablaTipoUsoEspacioFisico_20191105_155122", m)
}

// Run the migrations
func (m *CrearTablaTipoUsoEspacioFisico_20191105_155122) Up() {
	m.SQL("CREATE TABLE IF NOT EXISTS oikos.tipo_uso_espacio_fisico ( id serial NOT NULL, tipo_uso_id integer NOT NULL, espacio_fisico_id integer NOT NULL, activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_tipo_uso_espacio_fisico PRIMARY KEY (id), CONSTRAINT fk_tipo_uso_tipo_uso_espacio_fisico FOREIGN KEY (tipo_uso_id) REFERENCES oikos.tipo_uso(id), CONSTRAINT fk_espacio_fisico_tipo_uso_espacio_fisico FOREIGN KEY (espacio_fisico_id) REFERENCES oikos.espacio_fisico(id) );")
	m.SQL("COMMENT ON TABLE oikos.tipo_uso_espacio_fisico  IS 'Tabla de rompimiento entre tipo_espacio_fisico y espacio_fisico';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.id    IS 'Identificador de la tabla ';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.tipo_uso_id    IS 'Campo que contiene el identificador del espacio fisico';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.espacio_fisico_id    IS 'Campo para el id del espacio fisico';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")

}

// Reverse the migrations
func (m *CrearTablaTipoUsoEspacioFisico_20191105_155122) Down() {
	m.SQL("DROP TABLE IF EXISTS oikos.tipo_uso_espacio_fisico")
}
