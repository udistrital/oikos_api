CREATE SCHEMA oikos;

CREATE SEQUENCE oikos.tipo_espacio_fisico_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE oikos.tipo_espacio_fisico
(
    id integer NOT NULL DEFAULT nextval('oikos.tipo_espacio_fisico_id_seq'::regclass),
	nombre character varying(100) NOT NULL,     -- 
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tipo_espacio_fisico PRIMARY KEY (id)
);

COMMENT ON TABLE oikos.tipo_espacio_fisico IS 'Tabla que contiene los tipos de espacios fisicos existentes.';
COMMENT ON COLUMN oikos.tipo_espacio_fisico.id IS 'Identificador de cada tipo de espacio fisico que pertenece a la Universidad Distr';
COMMENT ON COLUMN oikos.tipo_espacio_fisico.nombre IS 'Nombre del tipo de espacio fisico perteneciente a la Universidad Distrital';
COMMENT ON COLUMN oikos.tipo_espacio_fisico.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de tipo_espacio_fisico.';
COMMENT ON COLUMN oikos.tipo_espacio_fisico.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN oikos.tipo_espacio_fisico.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN oikos.tipo_espacio_fisico.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN oikos.tipo_espacio_fisico.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


CREATE SEQUENCE oikos.espacio_fisico_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE oikos.espacio_fisico
(   
    id integer NOT NULL DEFAULT nextval('oikos.espacio_fisico_id_seq'::regclass),
	nombre character varying(100) NOT NULL,     -- 
	descripcion character varying(100),
    area numeric(5,2),
    capacidad integer,
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
    tipo_espacio_fisico_id integer NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
   	CONSTRAINT pk_espacio_fisico PRIMARY KEY (id),
    CONSTRAINT fk_tipo_espacio_fisico_espacio_fisico FOREIGN KEY (tipo_espacio_fisico_id) REFERENCES oikos.tipo_espacio_fisico(id)
  
);

COMMENT ON TABLE oikos.espacio_fisico  IS 'Tabla de Rompimiento que reune los atributos de un espacio fisico';
COMMENT ON COLUMN oikos.espacio_fisico.id    IS 'Identificador del espacio fisico especifico de la Universidad Distrital.';
COMMENT ON COLUMN oikos.espacio_fisico.nombre    IS 'Nombre perteneciente al espacio físico';
COMMENT ON COLUMN oikos.espacio_fisico.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de espacio_fisico.';
COMMENT ON COLUMN oikos.espacio_fisico.area IS 'Área del espacio fisico en metros cuadrados.';
COMMENT ON COLUMN oikos.espacio_fisico.capacidad IS 'Número de personas en espacio físico';
COMMENT ON COLUMN oikos.espacio_fisico.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN oikos.espacio_fisico.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN oikos.espacio_fisico.tipo_espacio_fisico_id    IS 'Llave foranea que contiene el identificador del tipo de espacio fisico de la entidad tipo_espacio_fisico';
COMMENT ON COLUMN oikos.espacio_fisico.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN oikos.espacio_fisico.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


CREATE SEQUENCE oikos.tipo_uso_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE oikos.tipo_uso
(     
    id integer NOT NULL DEFAULT nextval('oikos.tipo_uso_id_seq'::regclass),
	nombre character varying(100) NOT NULL,     -- 
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tipo_uso PRIMARY KEY (id)
);

COMMENT ON TABLE oikos.tipo_uso     IS 'Tabla quie contiene los diversos usos que puede tener un espacio fisico.';
COMMENT ON COLUMN oikos.tipo_uso.id     IS 'Identificador del tipo de uso de espacios fisicos.';
COMMENT ON COLUMN oikos.tipo_uso.nombre     IS 'Nombre del uso que se le va a dar al espacio fisico.';
COMMENT ON COLUMN oikos.tipo_uso.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de tipo_uso.';
COMMENT ON COLUMN oikos.tipo_uso.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN oikos.tipo_uso.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN oikos.tipo_uso.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN oikos.tipo_uso.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

CREATE SEQUENCE oikos.tipo_uso_espacio_fisico_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE oikos.tipo_uso_espacio_fisico
(
    id integer NOT NULL DEFAULT nextval('oikos.tipo_uso_espacio_fisico_id_seq'::regclass),
    tipo_uso_id integer NOT NULL,
    espacio_fisico_id integer NOT NULL,
    activo boolean NOT NULL,
    fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
    CONSTRAINT pk_tipo_uso_espacio_fisico PRIMARY KEY (id),
    CONSTRAINT fk_tipo_uso_tipo_uso_espacio_fisico FOREIGN KEY (tipo_uso_id) REFERENCES oikos.tipo_uso(id),
    CONSTRAINT fk_espacio_fisico_tipo_uso_espacio_fisico FOREIGN KEY (espacio_fisico_id) REFERENCES oikos.espacio_fisico(id)
   
);

COMMENT ON TABLE oikos.tipo_uso_espacio_fisico  IS 'Tabla de rompimiento entre tipo_espacio_fisico y espacio_fisico';
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.id    IS 'Identificador de la tabla ';
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.tipo_uso_id    IS 'Campo que contiene el identificador del espacio fisico';
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.espacio_fisico_id    IS 'Campo para el id del espacio fisico';
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


CREATE SEQUENCE oikos.espacio_fisico_padre_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;


CREATE TABLE oikos.espacio_fisico_padre
(
    id integer NOT NULL DEFAULT nextval('oikos.espacio_fisico_padre_id_seq'::regclass),
    padre_id integer NOT NULL,
    hijo_id integer NOT NULL,
    fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
    CONSTRAINT pk_espacio_fisico_padre PRIMARY KEY (id),
    CONSTRAINT fk_espacio_fisico_espacio_fisico_padre FOREIGN KEY (padre_id) REFERENCES oikos.espacio_fisico(id),
    CONSTRAINT fk_espacio_fisico_espacio_fisico_hijo FOREIGN KEY (hijo_id) REFERENCES oikos.espacio_fisico(id),
    CONSTRAINT uq_hijo_id_espacio_fisico_padre UNIQUE (hijo_id),    
    CONSTRAINT uq_hijo_id_padre_id_espacio_fisico_padre UNIQUE (padre_id, hijo_id)
 
);

COMMENT ON TABLE oikos.espacio_fisico_padre    IS 'Contiene las relaciones de los espacios fisicos.';
COMMENT ON COLUMN oikos.espacio_fisico_padre.id    IS 'Identificador de la tabla.';
COMMENT ON COLUMN oikos.espacio_fisico_padre.padre_id    IS 'Identificador del espacio fisico padre.';
COMMENT ON COLUMN oikos.espacio_fisico_padre.hijo_id    IS 'Identificador del espacio fisico hijo.';
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
COMMENT ON CONSTRAINT "uq_hijo_id_espacio_fisico_padre" ON oikos.espacio_fisico_padre    IS 'Restringe que el arbol no se vuelva un grafo.';
COMMENT ON CONSTRAINT "uq_hijo_id_padre_id_espacio_fisico_padre" ON oikos.espacio_fisico_padre    IS 'Restriccion para que solo pueda existir una unica relacion entre un padre y un hijo.';


CREATE SEQUENCE oikos.tipo_dependencia_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;


CREATE TABLE oikos.tipo_dependencia
(
    id integer NOT NULL DEFAULT nextval('oikos.tipo_dependencia_id_seq'::regclass),
	nombre character varying(100) NOT NULL,     -- 
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tipo_dependencia PRIMARY KEY (id)
);

COMMENT ON TABLE oikos.tipo_dependencia   IS 'Tabla que contiene los distintos tipos de dependencia que hay en la Universidad Distrital.';
COMMENT ON COLUMN oikos.tipo_dependencia.id    IS 'Identificador de la tabla.';
COMMENT ON COLUMN oikos.tipo_dependencia.nombre    IS 'Campo que contiene el tipo de dependencia.';
COMMENT ON COLUMN oikos.tipo_dependencia.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de tipo_espacio_fisico.';
COMMENT ON COLUMN oikos.tipo_dependencia.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN oikos.tipo_dependencia.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN oikos.tipo_dependencia.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN oikos.tipo_dependencia.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


CREATE SEQUENCE oikos.dependencia_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE oikos.dependencia
(
    id integer NOT NULL DEFAULT nextval('oikos.dependencia_id_seq'::regclass),
    nombre character varying(100) NOT NULL,
    telefono_dependencia character varying(500) NOT NULL,
    correo_electronico character varying(100),
    activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
    CONSTRAINT pk_dependencia PRIMARY KEY (id)
  
);

COMMENT ON TABLE oikos.dependencia  IS 'Tabla que contiene las dependencias de la Universidad Distrital.';
COMMENT ON COLUMN oikos.dependencia.id    IS 'Identificador de la dependencia.';
COMMENT ON COLUMN oikos.dependencia.nombre    IS 'Nombre de la dependencia perteneciente a la Universidad Distrital.';
COMMENT ON COLUMN oikos.dependencia.telefono_dependencia     IS 'Indica el numero de telefono de la dependencia.';
COMMENT ON COLUMN oikos.dependencia.correo_electronico     IS 'Correo electrónico asociado a la dependencia.';
COMMENT ON COLUMN oikos.dependencia.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN oikos.dependencia.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN oikos.dependencia.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


CREATE SEQUENCE oikos.dependencia_tipo_dependencia_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE oikos.dependencia_tipo_dependencia
(
    id integer NOT NULL,
    tipo_dependencia_id integer NOT NULL,
    dependencia_id integer NOT NULL,
    activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
    CONSTRAINT pk_dependencia_tipo_dependencia PRIMARY KEY (id),
    CONSTRAINT fk_tipo_dependencia_dependencia_tipo_dependencia FOREIGN KEY (tipo_dependencia_id) REFERENCES oikos.tipo_dependencia(id),
    CONSTRAINT fk_dependencia_dependencia_tipo_dependencia FOREIGN KEY (dependencia_id) REFERENCES oikos.dependencia(id)
);

COMMENT ON TABLE oikos.dependencia_tipo_dependencia    IS 'Tabla de rompimiento entre tipo_dependencia y dependencia.';
COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.id    IS 'Identificador de la tabla.';
COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.tipo_dependencia_id    IS 'Campo que contiene el identificador del tipo dependencia.';
COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.dependencia_id    IS 'Campo para el id de la dependencia.';
COMMENT ON COLUMN oikos.dependencia.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN oikos.dependencia.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN oikos.dependencia.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

CREATE SEQUENCE oikos.dependencia_padre_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE oikos.dependencia_padre
(
    id integer NOT NULL DEFAULT nextval('oikos.dependencia_padre_id_seq'::regclass),
    padre_id integer NOT NULL,
    hija_id integer NOT NULL,
    activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
    CONSTRAINT pk_dependencia_padre PRIMARY KEY (id),
    CONSTRAINT uq_hija_id_dependencia_padre UNIQUE (hija_id),    
    CONSTRAINT uq_hija_id_padre_id_dependencia_padre UNIQUE (padre_id, hija_id),
    CONSTRAINT fk_dependencia_dependencia_padre FOREIGN KEY (padre_id) REFERENCES oikos.dependencia(id),
    CONSTRAINT fk_dependencia_dependencia_hija FOREIGN KEY (hija_id) REFERENCES oikos.dependencia(id)
);

COMMENT ON TABLE oikos.dependencia_padre   IS 'Tabla que contiene las dependencias de la Universidad Distrital';
COMMENT ON COLUMN oikos.dependencia_padre.id     IS 'Identificador de la dependencia.';
COMMENT ON COLUMN oikos.dependencia_padre.padre_id     IS 'Id dependencia padre';
COMMENT ON COLUMN oikos.dependencia_padre.hija_id     IS 'Identificador dependencia hija (si la hay)';
COMMENT ON COLUMN oikos.dependencia_padre.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN oikos.dependencia_padre.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN oikos.dependencia_padre.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

CREATE SEQUENCE oikos.asignacion_espacio_fisico_dependencia_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE oikos.asignacion_espacio_fisico_dependencia
(
    id integer NOT NULL DEFAULT nextval('oikos.dependencia_padre_id_seq'::regclass),
    espacio_fisico_id integer NOT NULL,
    dependencia_id integer NOT NULL,
    activo boolean NOT NULL,
    fecha_inicio TIMESTAMP NOT NULL,
    fecha_fin TIMESTAMP,
    documento_soporte integer,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
    CONSTRAINT pk_asignacion_espacio_fisico_dependencia PRIMARY KEY (id),
    CONSTRAINT fk_espacio_fisico_asignacion_espacio_fisico_dependencia FOREIGN KEY (espacio_fisico_id) REFERENCES oikos.espacio_fisico(id),
    CONSTRAINT fk_dependencia_asignacion_espacio_fisico_dependencia FOREIGN KEY (dependencia_id) REFERENCES oikos.dependencia(id)
);

COMMENT ON TABLE oikos.asignacion_espacio_fisico_dependencia IS 'Tabla de rompimiento que reune los atributos necesarios para la asignacion de un espacio fisico.';
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.id  IS 'Identificador de la asignacion del espacio fisico';
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.espacio_fisico_id     IS 'Campo que contiene el id, de la llave foranea';
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.dependencia_id     IS 'Identificador que contiene el id de dependencia';
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.fecha_inicio IS 'Fecha de inicio de la asignacion.';
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.fecha_fin IS 'Fecha en la que finaliza la asignacion.';
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.documento_soporte IS 'Documento que soporta la asignacion del espacio fisico.';
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


