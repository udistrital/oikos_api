CREATE TABLE IF NOT EXISTS oikos.centro_costos (
    id SERIAL NOT NULL,
    dependencia_id INTEGER,
    sede_id INTEGER,
    codigo VARCHAR(12) NOT NULL,
    nombre TEXT NOT NULL,
    activo BOOLEAN NOT NULL,
    fecha_creacion TIMESTAMP NOT NULL,
    fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT uq_codigo_centro_costos UNIQUE (codigo),
    CONSTRAINT pk_centro_costos PRIMARY KEY (id)
);

COMMENT ON TABLE oikos.centro_costos IS 'Tabla para relacionar los centros de costos';
COMMENT ON COLUMN oikos.centro_costos.dependencia_id IS 'Dependencia asociada al centro de costos';
COMMENT ON COLUMN oikos.centro_costos.sede_id IS 'Sede asociada al centro de costos';
COMMENT ON COLUMN oikos.centro_costos.codigo IS 'Código único del centro de costos';
COMMENT ON COLUMN oikos.centro_costos.nombre IS 'Nombre del centro de costos';

ALTER TABLE oikos.centro_costos
    ADD CONSTRAINT fk_centro_costos_dependencia_id FOREIGN KEY (dependencia_id)
    REFERENCES oikos.dependencia (id) MATCH FULL
    ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE oikos.centro_costos
    ADD CONSTRAINT fk_centro_costos_sede_id FOREIGN KEY (sede_id)
    REFERENCES oikos.espacio_fisico (id) MATCH FULL
    ON DELETE RESTRICT ON UPDATE CASCADE;
