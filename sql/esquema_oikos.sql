-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.1
-- PostgreSQL version: 10.0
-- Project Site: pgmodeler.io
-- Model Author: ---


-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- -- object: bd_oas | type: DATABASE --
-- -- DROP DATABASE IF EXISTS bd_oas;
-- CREATE DATABASE bd_oas
-- 	ENCODING = 'UTF8'
-- 	LC_COLLATE = 'en_US.UTF-8'
-- 	LC_CTYPE = 'en_US.UTF-8'
-- 	TABLESPACE = pg_default
-- 	OWNER = postgres;
-- -- ddl-end --
-- 

-- object: oikos | type: SCHEMA --
-- DROP SCHEMA IF EXISTS oikos CASCADE;
CREATE SCHEMA oikos;
-- ddl-end --
-- ddl-end --

SET search_path TO pg_catalog,public,oikos;
-- ddl-end --

-- object: oikos.tipo_espacio_fisico_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS oikos.tipo_espacio_fisico_id_seq CASCADE;
CREATE SEQUENCE oikos.tipo_espacio_fisico_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
-- ddl-end --

-- object: oikos.tipo_espacio_fisico | type: TABLE --
-- DROP TABLE IF EXISTS oikos.tipo_espacio_fisico CASCADE;
CREATE TABLE oikos.tipo_espacio_fisico(
	id integer NOT NULL DEFAULT nextval('oikos.tipo_espacio_fisico_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_espacio_fisico PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE oikos.tipo_espacio_fisico IS 'Tabla que contiene los tipos de espacios fisicos existentes.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_espacio_fisico.id IS 'Identificador de cada tipo de espacio fisico que pertenece a la Universidad Distr';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_espacio_fisico.nombre IS 'Nombre del tipo de espacio fisico perteneciente a la Universidad Distrital';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_espacio_fisico.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de tipo_espacio_fisico.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_espacio_fisico.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_espacio_fisico.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_espacio_fisico.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_espacio_fisico.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
-- ddl-end --
-- ddl-end --

-- object: oikos.espacio_fisico_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS oikos.espacio_fisico_id_seq CASCADE;
CREATE SEQUENCE oikos.espacio_fisico_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
-- ddl-end --

-- object: oikos.espacio_fisico | type: TABLE --
-- DROP TABLE IF EXISTS oikos.espacio_fisico CASCADE;
CREATE TABLE oikos.espacio_fisico(
	id integer NOT NULL DEFAULT nextval('oikos.espacio_fisico_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	tipo_espacio_fisico_id integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	tipo_edificacion_id integer,
	tipo_terreno_id integer,
	CONSTRAINT pk_espacio_fisico PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE oikos.espacio_fisico IS 'Tabla de Rompimiento que reune los atributos de un espacio fisico';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico.id IS 'Identificador del espacio fisico especifico de la Universidad Distrital.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico.nombre IS 'Nombre perteneciente al espacio físico';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de espacio_fisico.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico.tipo_espacio_fisico_id IS 'Llave foranea que contiene el identificador del tipo de espacio fisico de la entidad tipo_espacio_fisico';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico.tipo_edificacion_id IS 'Hace referencia al id al tipo de edificación en el api parametros_crud';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico.tipo_terreno_id IS 'Hace referencia al id al tipo de terreno en el api parametros_crud';
-- ddl-end --
-- ddl-end --

-- object: oikos.tipo_uso_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS oikos.tipo_uso_id_seq CASCADE;
CREATE SEQUENCE oikos.tipo_uso_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
-- ddl-end --

-- object: oikos.tipo_uso | type: TABLE --
-- DROP TABLE IF EXISTS oikos.tipo_uso CASCADE;
CREATE TABLE oikos.tipo_uso(
	id integer NOT NULL DEFAULT nextval('oikos.tipo_uso_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_uso PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE oikos.tipo_uso IS 'Tabla quie contiene los diversos usos que puede tener un espacio fisico.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_uso.id IS 'Identificador del tipo de uso de espacios fisicos.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_uso.nombre IS 'Nombre del uso que se le va a dar al espacio fisico.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_uso.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de tipo_uso.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_uso.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_uso.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_uso.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_uso.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
-- ddl-end --
-- ddl-end --

-- object: oikos.tipo_uso_espacio_fisico_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS oikos.tipo_uso_espacio_fisico_id_seq CASCADE;
CREATE SEQUENCE oikos.tipo_uso_espacio_fisico_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
-- ddl-end --

-- object: oikos.tipo_uso_espacio_fisico | type: TABLE --
-- DROP TABLE IF EXISTS oikos.tipo_uso_espacio_fisico CASCADE;
CREATE TABLE oikos.tipo_uso_espacio_fisico(
	id integer NOT NULL DEFAULT nextval('oikos.tipo_uso_espacio_fisico_id_seq'::regclass),
	tipo_uso_id integer NOT NULL,
	espacio_fisico_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_uso_espacio_fisico PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE oikos.tipo_uso_espacio_fisico IS 'Tabla de rompimiento entre tipo_espacio_fisico y espacio_fisico';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.id IS 'Identificador de la tabla ';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.tipo_uso_id IS 'Campo que contiene el identificador del espacio fisico';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.espacio_fisico_id IS 'Campo para el id del espacio fisico';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
-- ddl-end --
-- ddl-end --

-- object: oikos.espacio_fisico_padre_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS oikos.espacio_fisico_padre_id_seq CASCADE;
CREATE SEQUENCE oikos.espacio_fisico_padre_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
-- ddl-end --

-- object: oikos.espacio_fisico_padre | type: TABLE --
-- DROP TABLE IF EXISTS oikos.espacio_fisico_padre CASCADE;
CREATE TABLE oikos.espacio_fisico_padre(
	id integer NOT NULL DEFAULT nextval('oikos.espacio_fisico_padre_id_seq'::regclass),
	padre_id integer NOT NULL,
	hijo_id integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_espacio_fisico_padre PRIMARY KEY (id),
	CONSTRAINT uq_hijo_id_espacio_fisico_padre UNIQUE (hijo_id),
	CONSTRAINT uq_hijo_id_padre_id_espacio_fisico_padre UNIQUE (padre_id,hijo_id)

);
-- ddl-end --
COMMENT ON TABLE oikos.espacio_fisico_padre IS 'Contiene las relaciones de los espacios fisicos.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico_padre.id IS 'Identificador de la tabla.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico_padre.padre_id IS 'Identificador del espacio fisico padre.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico_padre.hijo_id IS 'Identificador del espacio fisico hijo.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico_padre.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico_padre.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
-- ddl-end --
COMMENT ON CONSTRAINT uq_hijo_id_espacio_fisico_padre ON oikos.espacio_fisico_padre  IS 'Restringe que el arbol no se vuelva un grafo.';
-- ddl-end --
COMMENT ON CONSTRAINT uq_hijo_id_padre_id_espacio_fisico_padre ON oikos.espacio_fisico_padre  IS 'Restriccion para que solo pueda existir una unica relacion entre un padre y un hijo.';
-- ddl-end --
-- ddl-end --

-- object: oikos.tipo_dependencia_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS oikos.tipo_dependencia_id_seq CASCADE;
CREATE SEQUENCE oikos.tipo_dependencia_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
-- ddl-end --

-- object: oikos.tipo_dependencia | type: TABLE --
-- DROP TABLE IF EXISTS oikos.tipo_dependencia CASCADE;
CREATE TABLE oikos.tipo_dependencia(
	id integer NOT NULL DEFAULT nextval('oikos.tipo_dependencia_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_dependencia PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE oikos.tipo_dependencia IS 'Tabla que contiene los distintos tipos de dependencia que hay en la Universidad Distrital.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_dependencia.id IS 'Identificador de la tabla.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_dependencia.nombre IS 'Campo que contiene el tipo de dependencia.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_dependencia.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de tipo_espacio_fisico.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_dependencia.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_dependencia.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_dependencia.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN oikos.tipo_dependencia.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
-- ddl-end --
-- ddl-end --

-- object: oikos.dependencia_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS oikos.dependencia_id_seq CASCADE;
CREATE SEQUENCE oikos.dependencia_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
-- ddl-end --

-- object: oikos.dependencia | type: TABLE --
-- DROP TABLE IF EXISTS oikos.dependencia CASCADE;
CREATE TABLE oikos.dependencia(
	id integer NOT NULL DEFAULT nextval('oikos.dependencia_id_seq'::regclass),
	nombre character varying(130) NOT NULL,
	telefono_dependencia character varying(500) NOT NULL,
	correo_electronico character varying(100),
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_dependencia PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE oikos.dependencia IS 'Tabla que contiene las dependencias de la Universidad Distrital.';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia.id IS 'Identificador de la dependencia.';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia.nombre IS 'Nombre de la dependencia perteneciente a la Universidad Distrital.';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia.telefono_dependencia IS 'Indica el numero de telefono de la dependencia.';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia.correo_electronico IS 'Correo electrónico asociado a la dependencia.';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
-- ddl-end --
-- ddl-end --

-- object: oikos.dependencia_tipo_dependencia_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS oikos.dependencia_tipo_dependencia_id_seq CASCADE;
CREATE SEQUENCE oikos.dependencia_tipo_dependencia_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
-- ddl-end --

-- object: oikos.dependencia_tipo_dependencia | type: TABLE --
-- DROP TABLE IF EXISTS oikos.dependencia_tipo_dependencia CASCADE;
CREATE TABLE oikos.dependencia_tipo_dependencia(
	id integer NOT NULL DEFAULT nextval('oikos.dependencia_tipo_dependencia_id_seq'::regclass),
	tipo_dependencia_id integer NOT NULL,
	dependencia_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_dependencia_tipo_dependencia PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE oikos.dependencia_tipo_dependencia IS 'Tabla de rompimiento entre tipo_dependencia y dependencia.';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.id IS 'Identificador de la tabla.';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.tipo_dependencia_id IS 'Campo que contiene el identificador del tipo dependencia.';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.dependencia_id IS 'Campo para el id de la dependencia.';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
-- ddl-end --
-- ddl-end --

-- object: oikos.dependencia_padre_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS oikos.dependencia_padre_id_seq CASCADE;
CREATE SEQUENCE oikos.dependencia_padre_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
-- ddl-end --

-- object: oikos.dependencia_padre | type: TABLE --
-- DROP TABLE IF EXISTS oikos.dependencia_padre CASCADE;
CREATE TABLE oikos.dependencia_padre(
	id integer NOT NULL DEFAULT nextval('oikos.dependencia_padre_id_seq'::regclass),
	padre_id integer NOT NULL,
	hija_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_dependencia_padre PRIMARY KEY (id),
	CONSTRAINT uq_hija_id_dependencia_padre UNIQUE (hija_id),
	CONSTRAINT uq_hija_id_padre_id_dependencia_padre UNIQUE (padre_id,hija_id)

);
-- ddl-end --
COMMENT ON TABLE oikos.dependencia_padre IS 'Tabla que contiene las dependencias de la Universidad Distrital';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia_padre.id IS 'Identificador de la dependencia.';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia_padre.padre_id IS 'Id dependencia padre';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia_padre.hija_id IS 'Identificador dependencia hija (si la hay)';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia_padre.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia_padre.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN oikos.dependencia_padre.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
-- ddl-end --
-- ddl-end --

-- object: oikos.asignacion_espacio_fisico_dependencia_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS oikos.asignacion_espacio_fisico_dependencia_id_seq CASCADE;
CREATE SEQUENCE oikos.asignacion_espacio_fisico_dependencia_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
-- ddl-end --

-- object: oikos.asignacion_espacio_fisico_dependencia | type: TABLE --
-- DROP TABLE IF EXISTS oikos.asignacion_espacio_fisico_dependencia CASCADE;
CREATE TABLE oikos.asignacion_espacio_fisico_dependencia(
	id integer NOT NULL DEFAULT nextval('oikos.asignacion_espacio_fisico_dependencia_id_seq'::regclass),
	espacio_fisico_id integer NOT NULL,
	dependencia_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_inicio timestamp NOT NULL,
	fecha_fin timestamp,
	documento_soporte integer,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_asignacion_espacio_fisico_dependencia PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE oikos.asignacion_espacio_fisico_dependencia IS 'Tabla de rompimiento que reune los atributos necesarios para la asignacion de un espacio fisico.';
-- ddl-end --
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.id IS 'Identificador de la asignacion del espacio fisico';
-- ddl-end --
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.espacio_fisico_id IS 'Campo que contiene el id, de la llave foranea';
-- ddl-end --
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.dependencia_id IS 'Identificador que contiene el id de dependencia';
-- ddl-end --
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
-- ddl-end --
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.fecha_inicio IS 'Fecha de inicio de la asignacion.';
-- ddl-end --
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.fecha_fin IS 'Fecha en la que finaliza la asignacion.';
-- ddl-end --
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.documento_soporte IS 'Documento que soporta la asignacion del espacio fisico.';
-- ddl-end --
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
-- ddl-end --
-- ddl-end --

-- object: oikos.campo_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS oikos.campo_id_seq CASCADE;
CREATE SEQUENCE oikos.campo_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
-- ddl-end --

-- object: oikos.campo | type: TABLE --
-- DROP TABLE IF EXISTS oikos.campo CASCADE;
CREATE TABLE oikos.campo(
	id integer NOT NULL DEFAULT nextval('oikos.campo_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_campo PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE oikos.campo IS 'Tabla de los campos o atributos que se necesiten en un espacio fisico.';
-- ddl-end --
COMMENT ON COLUMN oikos.campo.id IS 'Identificador de la tabla.';
-- ddl-end --
COMMENT ON COLUMN oikos.campo.nombre IS 'Nombre del nuevo campo que se requiere para espacios fisicos.';
-- ddl-end --
COMMENT ON COLUMN oikos.campo.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion del nuevo campo.';
-- ddl-end --
COMMENT ON COLUMN oikos.campo.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
-- ddl-end --
COMMENT ON COLUMN oikos.campo.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
-- ddl-end --
COMMENT ON COLUMN oikos.campo.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN oikos.campo.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
-- ddl-end --
-- ddl-end --

-- object: oikos.espacio_fisico_campo_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS oikos.espacio_fisico_campo_id_seq CASCADE;
CREATE SEQUENCE oikos.espacio_fisico_campo_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;
-- ddl-end --
-- ddl-end --

-- object: oikos.espacio_fisico_campo | type: TABLE --
-- DROP TABLE IF EXISTS oikos.espacio_fisico_campo CASCADE;
CREATE TABLE oikos.espacio_fisico_campo(
	id integer NOT NULL DEFAULT nextval('oikos.espacio_fisico_campo_id_seq'::regclass),
	valor character varying(50) NOT NULL,
	espacio_fisico_id integer NOT NULL,
	campo_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_inicio timestamp NOT NULL,
	fecha_fin timestamp,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_espacio_fisico_campo PRIMARY KEY (id),
	CONSTRAINT uq_campo UNIQUE (campo_id,espacio_fisico_id)

);
-- ddl-end --
COMMENT ON TABLE oikos.espacio_fisico_campo IS 'Tabla de rompimiento entre campo y espacio_fisico.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico_campo.id IS 'Identificador de la asignacion del espacio fisico.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico_campo.valor IS 'Valor del nuevo campo para el espacio fisico.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico_campo.espacio_fisico_id IS 'Identificador de la tabla espacio fisico.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico_campo.campo_id IS 'Identificador de la tabla campo';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico_campo.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico_campo.fecha_inicio IS 'Fecha de inicio de la asignacion.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico_campo.fecha_fin IS 'Fecha en la que finaliza la asignacion.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico_campo.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
-- ddl-end --
COMMENT ON COLUMN oikos.espacio_fisico_campo.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
-- ddl-end --
-- ddl-end --

-- object: fk_tipo_espacio_fisico_espacio_fisico | type: CONSTRAINT --
-- ALTER TABLE oikos.espacio_fisico DROP CONSTRAINT IF EXISTS fk_tipo_espacio_fisico_espacio_fisico CASCADE;
ALTER TABLE oikos.espacio_fisico ADD CONSTRAINT fk_tipo_espacio_fisico_espacio_fisico FOREIGN KEY (tipo_espacio_fisico_id)
REFERENCES oikos.tipo_espacio_fisico (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_tipo_uso_tipo_uso_espacio_fisico | type: CONSTRAINT --
-- ALTER TABLE oikos.tipo_uso_espacio_fisico DROP CONSTRAINT IF EXISTS fk_tipo_uso_tipo_uso_espacio_fisico CASCADE;
ALTER TABLE oikos.tipo_uso_espacio_fisico ADD CONSTRAINT fk_tipo_uso_tipo_uso_espacio_fisico FOREIGN KEY (tipo_uso_id)
REFERENCES oikos.tipo_uso (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_espacio_fisico_tipo_uso_espacio_fisico | type: CONSTRAINT --
-- ALTER TABLE oikos.tipo_uso_espacio_fisico DROP CONSTRAINT IF EXISTS fk_espacio_fisico_tipo_uso_espacio_fisico CASCADE;
ALTER TABLE oikos.tipo_uso_espacio_fisico ADD CONSTRAINT fk_espacio_fisico_tipo_uso_espacio_fisico FOREIGN KEY (espacio_fisico_id)
REFERENCES oikos.espacio_fisico (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_espacio_fisico_espacio_fisico_padre | type: CONSTRAINT --
-- ALTER TABLE oikos.espacio_fisico_padre DROP CONSTRAINT IF EXISTS fk_espacio_fisico_espacio_fisico_padre CASCADE;
ALTER TABLE oikos.espacio_fisico_padre ADD CONSTRAINT fk_espacio_fisico_espacio_fisico_padre FOREIGN KEY (padre_id)
REFERENCES oikos.espacio_fisico (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_espacio_fisico_espacio_fisico_hijo | type: CONSTRAINT --
-- ALTER TABLE oikos.espacio_fisico_padre DROP CONSTRAINT IF EXISTS fk_espacio_fisico_espacio_fisico_hijo CASCADE;
ALTER TABLE oikos.espacio_fisico_padre ADD CONSTRAINT fk_espacio_fisico_espacio_fisico_hijo FOREIGN KEY (hijo_id)
REFERENCES oikos.espacio_fisico (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_tipo_dependencia_dependencia_tipo_dependencia | type: CONSTRAINT --
-- ALTER TABLE oikos.dependencia_tipo_dependencia DROP CONSTRAINT IF EXISTS fk_tipo_dependencia_dependencia_tipo_dependencia CASCADE;
ALTER TABLE oikos.dependencia_tipo_dependencia ADD CONSTRAINT fk_tipo_dependencia_dependencia_tipo_dependencia FOREIGN KEY (tipo_dependencia_id)
REFERENCES oikos.tipo_dependencia (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_dependencia_dependencia_tipo_dependencia | type: CONSTRAINT --
-- ALTER TABLE oikos.dependencia_tipo_dependencia DROP CONSTRAINT IF EXISTS fk_dependencia_dependencia_tipo_dependencia CASCADE;
ALTER TABLE oikos.dependencia_tipo_dependencia ADD CONSTRAINT fk_dependencia_dependencia_tipo_dependencia FOREIGN KEY (dependencia_id)
REFERENCES oikos.dependencia (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_dependencia_dependencia_padre | type: CONSTRAINT --
-- ALTER TABLE oikos.dependencia_padre DROP CONSTRAINT IF EXISTS fk_dependencia_dependencia_padre CASCADE;
ALTER TABLE oikos.dependencia_padre ADD CONSTRAINT fk_dependencia_dependencia_padre FOREIGN KEY (padre_id)
REFERENCES oikos.dependencia (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_dependencia_dependencia_hija | type: CONSTRAINT --
-- ALTER TABLE oikos.dependencia_padre DROP CONSTRAINT IF EXISTS fk_dependencia_dependencia_hija CASCADE;
ALTER TABLE oikos.dependencia_padre ADD CONSTRAINT fk_dependencia_dependencia_hija FOREIGN KEY (hija_id)
REFERENCES oikos.dependencia (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_espacio_fisico_asignacion_espacio_fisico_dependencia | type: CONSTRAINT --
-- ALTER TABLE oikos.asignacion_espacio_fisico_dependencia DROP CONSTRAINT IF EXISTS fk_espacio_fisico_asignacion_espacio_fisico_dependencia CASCADE;
ALTER TABLE oikos.asignacion_espacio_fisico_dependencia ADD CONSTRAINT fk_espacio_fisico_asignacion_espacio_fisico_dependencia FOREIGN KEY (espacio_fisico_id)
REFERENCES oikos.espacio_fisico (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_dependencia_asignacion_espacio_fisico_dependencia | type: CONSTRAINT --
-- ALTER TABLE oikos.asignacion_espacio_fisico_dependencia DROP CONSTRAINT IF EXISTS fk_dependencia_asignacion_espacio_fisico_dependencia CASCADE;
ALTER TABLE oikos.asignacion_espacio_fisico_dependencia ADD CONSTRAINT fk_dependencia_asignacion_espacio_fisico_dependencia FOREIGN KEY (dependencia_id)
REFERENCES oikos.dependencia (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_espacio_fisico_espacio_fisico_campo | type: CONSTRAINT --
-- ALTER TABLE oikos.espacio_fisico_campo DROP CONSTRAINT IF EXISTS fk_espacio_fisico_espacio_fisico_campo CASCADE;
ALTER TABLE oikos.espacio_fisico_campo ADD CONSTRAINT fk_espacio_fisico_espacio_fisico_campo FOREIGN KEY (espacio_fisico_id)
REFERENCES oikos.espacio_fisico (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_campo_espacio_fisico_campo | type: CONSTRAINT --
-- ALTER TABLE oikos.espacio_fisico_campo DROP CONSTRAINT IF EXISTS fk_campo_espacio_fisico_campo CASCADE;
ALTER TABLE oikos.espacio_fisico_campo ADD CONSTRAINT fk_campo_espacio_fisico_campo FOREIGN KEY (campo_id)
REFERENCES oikos.campo (id) MATCH SIMPLE
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --
