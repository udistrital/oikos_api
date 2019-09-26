CREATE TABLE oikos.campo
(
    id integer NOT NULL,
    nombre character varying COLLATE NOT NULL,
    descripcion character varying(50),
    CONSTRAINT opcion_dato_pk_id PRIMARY KEY (id)
)

COMMENT ON TABLE oikos.campo IS 'Tabla de los campos o atributos que se necesiten en un espacio fisico';
COMMENT ON COLUMN oikos.campo.nombre IS 'Nombre del nuevo campo'; 
COMMENT ON COLUMN oikos.campo.descripcion IS 'Descripcion del campo que se acabo de crear';


CREATE TABLE oikos.asignacion_espacio_fisico_dependencia
(
    id integer NOT NULL ,
    estado character varying NOT NULL,
    fecha_inicio date NOT NULL,
    fecha_fin date,
    documento_soporte character varying NOT NULL,
    espacio_fisico_id integer NOT NULL,
    dependencia_id integer NOT NULL,
    CONSTRAINT id_asignacion PRIMARY KEY (id),
    CONSTRAINT "FK_asignacion_espacio_fisico_dependencia_dependencia" FOREIGN KEY (dependencia_id)
        REFERENCES oikos.dependencia (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT "FK_asignacion_espacio_fisico_dependencia_espacio_fisico" FOREIGN KEY (espacio_fisico_id)
        REFERENCES oikos.espacio_fisico (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)

COMMENT ON TABLE oikos.asignacion_espacio_fisico_dependencia IS 'Tabla de rompimiento que reune los atributos necesarios para la asignacion de un espacio fisico.';
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.id  IS 'Identificador de la asignacion del espacio fisico';
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.espacio_fisico_id     IS 'Campo que contiene el id, de la llave foranea';
COMMENT ON COLUMN oikos.asignacion_espacio_fisico_dependencia.dependencia_id     IS 'Identificador que contiene el id de dependencia';
COMMENT ON CONSTRAINT id_asignacion ON oikos.asignacion_espacio_fisico_dependencia     IS 'Indica el codigo de asignacion del espacio fisico';


CREATE TABLE oikos.dependencia
(
    id integer NOT NULL DEFAULT nextval('oikos.dependencia_id_seq'::regclass),
    nombre character varying COLLATE pg_catalog."default" NOT NULL,
    telefono_dependencia character varying(500) NOT NULL,
    correo_electronico character varying COLLATE ,
    CONSTRAINT "PK_dependencia" PRIMARY KEY (id),
    CONSTRAINT "UQ_nombre_dependencia" UNIQUE (nombre)

)

COMMENT ON TABLE oikos.dependencia  IS 'Tabla que contiene las dependencias de la Universidad Distrital';
COMMENT ON COLUMN oikos.dependencia.id    IS 'Identificador de la dependencia.';
COMMENT ON COLUMN oikos.dependencia.nombre    IS 'Nombre de la dependencia perteneciente a la Universidad Distrital';
COMMENT ON COLUMN oikos.dependencia.telefono_dependencia     IS 'Indica el numero de telefono de la dependencia.';
COMMENT ON COLUMN oikos.dependencia.correo_electronico     IS 'Correo electrónico asociado a la dependencia';

CREATE TABLE oikos.dependencia_padre
(
    id integer NOT NULL DEFAULT nextval('oikos.dependencia_padre_id_seq'::regclass),
    padre integer NOT NULL,
    hija integer NOT NULL,
    CONSTRAINT "PK_dependencia_padre" PRIMARY KEY (id),
    CONSTRAINT "UQ_HIJA" UNIQUE (hija)
,
    CONSTRAINT "UQ_PADRE_HIJO" UNIQUE (padre, hija)
,
    CONSTRAINT "FK_DEPENDENCIA_CON_DEPENDENCIA_PADRE" FOREIGN KEY (padre)
        REFERENCES oikos.dependencia (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT "FK_DEPENDENCIA_DEPENDENCIA_PADRE_HIJA" FOREIGN KEY (hija)
        REFERENCES oikos.dependencia (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)

COMMENT ON TABLE oikos.dependencia_padre   IS 'Tabla que contiene las dependencias de la Universidad Distrital';
COMMENT ON COLUMN oikos.dependencia_padre.id     IS 'Identificador de la dependencia.';
COMMENT ON COLUMN oikos.dependencia_padre.padre     IS 'Id dependencia padre';
COMMENT ON COLUMN oikos.dependencia_padre.hija     IS 'Identificador dependencia hija (si la hay)';
COMMENT ON CONSTRAINT "UQ_HIJA" ON oikos.dependencia_padre     IS 'Restringe que el arbol, no se vuelva un grafo';
COMMENT ON CONSTRAINT "UQ_PADRE_HIJO" ON oikos.dependencia_padre     IS 'Restringe que una vez hecha la relacion, no se pueda repetir';

CREATE TABLE oikos.dependencia_tipo_dependencia
(
    id integer NOT NULL,
    tipo_dependencia_id integer NOT NULL,
    dependencia_id integer NOT NULL,
    CONSTRAINT "PK_DEPENDENCIA_TIPO_DEPENDENCIA" PRIMARY KEY (id),
    CONSTRAINT "FK_DEPENDENCIA" FOREIGN KEY (dependencia_id)
        REFERENCES oikos.dependencia (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT "FK_TIPO_DEPENDENCIA" FOREIGN KEY (tipo_dependencia_id)
        REFERENCES oikos.tipo_dependencia (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)


COMMENT ON TABLE oikos.dependencia_tipo_dependencia    IS 'Tabla de rompimiento entre tipo_dependencia y dependencia';
COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.id    IS 'Identificador de la tabla ';
COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.tipo_dependencia_id    IS 'Campo que contiene el identificador del tipo dependencia';
COMMENT ON COLUMN oikos.dependencia_tipo_dependencia.dependencia_id    IS 'Campo para el id de la dependencia';
COMMENT ON CONSTRAINT "PK_DEPENDENCIA_TIPO_DEPENDENCIA" ON oikos.dependencia_tipo_dependencia    IS 'Llave primaria de la tabla';


CREATE TABLE oikos.espacio_fisico
(
    id integer NOT NULL ,
    estado character varying ,
    tipo_espacio integer NOT NULL,
    nombre character varying COLLATE,
    codigo character varying COLLATE,
    CONSTRAINT "PK_espacio_fisico" PRIMARY KEY (id),
    CONSTRAINT "FK_TIPO_ESPACIO_FISICO" FOREIGN KEY (tipo_espacio)
        REFERENCES oikos.tipo_espacio_fisico (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT "CHECK_ESPACIO_FISICO_ESTADO" CHECK (estado::text = ANY (ARRAY['Activo'::character varying::text, 'Inactivo'::character varying::text]))
)

COMMENT ON TABLE oikos.espacio_fisico  IS 'Tabla de Rompimiento que reune los atributos de un espacio fisico';
COMMENT ON COLUMN oikos.espacio_fisico.id    IS 'Identificador del espacio fisico especifico de la Universidad Distrital.';
COMMENT ON COLUMN oikos.espacio_fisico.estado    IS 'Indica el estado del espacio fisico';
COMMENT ON COLUMN oikos.espacio_fisico.tipo_espacio    IS 'Llave foranea que contiene el identificador del tipo de espacio fisico de la entidad tipo_espacio_fisico';
COMMENT ON COLUMN oikos.espacio_fisico.nombre    IS 'Nombre perteneciente al espacio físico';
COMMENT ON COLUMN oikos.espacio_fisico.codigo    IS 'Código pertinente al espacio físico';
COMMENT ON CONSTRAINT "CHECK_ESPACIO_FISICO_ESTADO" ON oikos.espacio_fisico    IS 'Check que solo admite los estados de Activo O Inactivo de un espacio fisico.';


CREATE TABLE oikos.espacio_fisico_campo
(
    id integer NOT NULL,
    valor character varying(50) NOT NULL,
    espacio_fisico integer NOT NULL,
    campo integer NOT NULL,
    CONSTRAINT "PK_ESPACIO_FISICO_CAMPO_ID" PRIMARY KEY (id),
    CONSTRAINT "UQ_CAMPO" UNIQUE (campo, espacio_fisico)
,
    CONSTRAINT "FK_CAMPO" FOREIGN KEY (campo)
        REFERENCES oikos.campo (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT "FK_ESPACIO_FISICO" FOREIGN KEY (espacio_fisico)
        REFERENCES oikos.espacio_fisico (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)

COMMENT ON TABLE oikos.espacio_fisico_campo    IS 'Tabla de rompimiento entre campo y espacio fisico';
COMMENT ON COLUMN oikos.espacio_fisico_campo.valor    IS 'Contiene la informacion del campo';
COMMENT ON COLUMN oikos.espacio_fisico_campo.espacio_fisico    IS 'Campo de referencia de la entidad espacio fisico';
COMMENT ON COLUMN oikos.espacio_fisico_campo.campo    IS 'Referencia del campo';
COMMENT ON CONSTRAINT "PK_ESPACIO_FISICO_CAMPO_ID" ON oikos.espacio_fisico_campo    IS 'Llave primaria de la tabla de rompimiento';
COMMENT ON CONSTRAINT "UQ_CAMPO" ON oikos.espacio_fisico_campo    IS 'Restricción que hace que a un espacio físico solo se le pueda asociar un campo ya asociado al espacio fisico';

CREATE TABLE oikos.espacio_fisico_padre
(
    id integer NOT NULL ,
    padre integer NOT NULL,
    hijo integer NOT NULL,
    CONSTRAINT "PK_ESPACIO_FISICO_PADRE" PRIMARY KEY (id),
    CONSTRAINT "UK_HIJO" UNIQUE (hijo)
,
    CONSTRAINT "UK_PADRE_HIJO" UNIQUE (padre, hijo)
,
    CONSTRAINT "FK_ESPACIO_FISICO_PADRE_CON_ESPACIO_FISICO" FOREIGN KEY (padre)
        REFERENCES oikos.espacio_fisico (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT "FK_ESPACIO_FISICO_PADRE_HIJO_CON_ESPACIO_FISICO" FOREIGN KEY (hijo)
        REFERENCES oikos.espacio_fisico (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)

COMMENT ON TABLE oikos.espacio_fisico_padre    IS 'Contiene las relaciones de los espacios fisicos ';
COMMENT ON COLUMN oikos.espacio_fisico_padre.id    IS 'Identificador de la tabla ';
COMMENT ON COLUMN oikos.espacio_fisico_padre.padre    IS 'Identificador del espacio fisico padre';
COMMENT ON COLUMN oikos.espacio_fisico_padre.hijo    IS 'Identificador del espacio fisico hijo';
COMMENT ON CONSTRAINT "PK_ESPACIO_FISICO_PADRE" ON oikos.espacio_fisico_padre    IS 'Llave primaria de espacio fisico padre';
COMMENT ON CONSTRAINT "UK_HIJO" ON oikos.espacio_fisico_padre    IS 'Restringe que el arbol no se vuelva un grafo';
COMMENT ON CONSTRAINT "UK_PADRE_HIJO" ON oikos.espacio_fisico_padre    IS 'Restriccion para que solo pueda existir una unica relacion entre un padre y un hijo';


CREATE TABLE oikos.tipo_dependencia
(
    id integer NOT NULL,
    nombre character varying NOT NULL,
    CONSTRAINT "PK_TIPO_DEPENDENCIA" PRIMARY KEY (id)
)

COMMENT ON TABLE oikos.tipo_dependencia   IS 'Tabla que contiene los distintos tipos de dependencia que hay en la Universidad Distrital.';
COMMENT ON COLUMN oikos.tipo_dependencia.id    IS 'Identificador de la tabla ';
COMMENT ON COLUMN oikos.tipo_dependencia.nombre    IS 'Campo que contiene el tipo de dependencia';
COMMENT ON CONSTRAINT "PK_TIPO_DEPENDENCIA" ON oikos.tipo_dependencia    IS 'Llave primaria de la tabla';

CREATE TABLE oikos.tipo_espacio_fisico
(
    id integer NOT NULL ,
    nombre character varying NOT NULL,
    CONSTRAINT "PK_tipo_espacio_fisico" PRIMARY KEY (id),
    CONSTRAINT "UQ_nombre_tipo_espacio_fisico" UNIQUE (nombre)

)

COMMENT ON TABLE oikos.tipo_espacio_fisico    IS 'Tabla que contiene los tipos de espacios fisicos existentes.';
COMMENT ON COLUMN oikos.tipo_espacio_fisico.id    IS 'Identificador de cada tipo de espacio fisico que pertenece a la Universidad Distr';
COMMENT ON COLUMN oikos.tipo_espacio_fisico.nombre    IS 'Nombre del tipo de espacio fisico perteneciente a la Universidad Distrital';


CREATE TABLE oikos.tipo_uso
(
    id integer NOT NULL,
    nombre character varying NOT NULL,
    CONSTRAINT "PK_TIPO_USO" PRIMARY KEY (id),
    CONSTRAINT "UQ_NOMBRE_TIPO_USO" UNIQUE (nombre)
,
    CONSTRAINT "CHECK_TIPO_USO_ESPACIO_FISICO" CHECK (nombre::text = ANY (ARRAY['Administrativo'::character varying::text, 'Deportivo'::character varying::text, 'Académico'::character varying::text]))
)

COMMENT ON TABLE oikos.tipo_uso     IS 'Tabla quie contiene los diversos usos que puede tener un espacio fisicoi';
COMMENT ON COLUMN oikos.tipo_uso.id     IS 'Identificador del tipo de uso de espacios fisicos';
COMMENT ON COLUMN oikos.tipo_uso.nombre     IS 'Nombre del uso que se le va a dar al espacio fisico';
COMMENT ON CONSTRAINT "PK_TIPO_USO" ON oikos.tipo_uso     IS 'Llave primaria del tipo de uso de espacio fisico';
COMMENT ON CONSTRAINT "UQ_NOMBRE_TIPO_USO" ON oikos.tipo_uso    IS 'Restriccion de nombre para el uso de espacio fisico';
COMMENT ON CONSTRAINT "CHECK_TIPO_USO_ESPACIO_FISICO" ON oikos.tipo_uso     IS 'Check que restringe los valores que se pueden seleccionar';

CREATE TABLE oikos.tipo_uso_espacio_fisico
(
    id integer NOT NULL,
    tipo_uso_id integer NOT NULL,
    espacio_fisico_id integer NOT NULL,
    CONSTRAINT "PK_TIPO_USO_ESPACIO_FISICO" PRIMARY KEY (id),
    CONSTRAINT "FK_ESPACIO_FISICO" FOREIGN KEY (espacio_fisico_id)
        REFERENCES oikos.espacio_fisico (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT "FK_TIPO_USO" FOREIGN KEY (tipo_uso_id)
        REFERENCES oikos.tipo_uso (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)

COMMENT ON TABLE oikos.tipo_uso_espacio_fisico    IS 'Tabla de rompimiento entre tipo_espacio_fisico y espacio_fisico';
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.id    IS 'Identificador de la tabla ';
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.tipo_uso_id    IS 'Campo que contiene el identificador del espacio fisico';
COMMENT ON COLUMN oikos.tipo_uso_espacio_fisico.espacio_fisico_id    IS 'Campo para el id del espacio fisico';
COMMENT ON CONSTRAINT "PK_TIPO_USO_ESPACIO_FISICO" ON oikos.tipo_uso_espacio_fisico    IS 'Llave primaria de la tabla';